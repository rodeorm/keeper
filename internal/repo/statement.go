package repo

func (s *postgresStorage) prepareStatements() error {

	stmtRegUser, err := s.DB.Preparex(`INSERT INTO cmn.Users (login, name, email, phone, password) SELECT $1, $2, $3, $4, $5 RETURNING id;`)
	if err != nil {
		return err
	}
	stmtUpdateUser, err := s.DB.Preparex(`UPDATE cmn.Users SET name = $2, email = $3, phone = $4, password = $5, verified = $6 WHERE ID = $1;`)
	if err != nil {
		return err
	}
	stmtAuthUser, err := s.DB.Preparex(`SELECT Login, Password FROM cmn.Users WHERE Login = $1;`)
	if err != nil {
		return err
	}
	stmtVerifyUser, err := s.DB.Preparex(`SELECT e.id FROM cmn.emails AS e INNER JOIN cmn.Users AS u ON u.id = e.UserID WHERE u.Login = $1 AND e.sendeddate + ($2 * INTERVAL '1 hour') > NOW() 
	AND e.Used = false AND e.OTP = $3; `)
	if err != nil {
		return err
	}

	stmtAddByte, err := s.DB.Preparex(`INSERT INTO dbo.Data (userid, typeid, bytedata, createddate) SELECT $1, $2, $3, $4 RETURNING id;`)
	if err != nil {
		return err
	}
	stmtSelectByte, err := s.DB.Preparex(`SELECT id, bytedata, createddate FROM dbo.Data WHERE userid = $1 AND typeid = $2;`)
	if err != nil {
		return err
	}
	stmtUpdateByte, err := s.DB.Preparex(`UPDATE dbo.Data SET userid = $2, typeid = $3, bytedata = $4, createddate = $5 WHERE id = $1;`)
	if err != nil {
		return err
	}
	stmtDeleteByte, err := s.DB.Preparex(`DELETE FROM dbo.Data WHERE id = $1;`)
	if err != nil {
		return err
	}
	stmtStartSession, err := s.DB.Preparex(`INSERT INTO cmn.Sessions (UserID, LoginDate, LastActionDate) SELECT $1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP`)
	if err != nil {
		return err
	}
	stmtUpdateSession, err := s.DB.Preparex(`UPDATE cmn.Sessions SET LastActionDate = CURRENT_TIMESTAMP WHERE id = $1;`)
	if err != nil {
		return err
	}
	stmtEndSession, err := s.DB.Preparex(`UPDATE cmn.Sessions SET LogoutDate = CURRENT_TIMESTAMP WHERE id = $1;`)
	if err != nil {
		return err
	}
	stmtAddEmail, err := s.DB.Preparex(`INSERT INTO cmn.Emails (UserID, OTP, Email) SELECT $1, $2, $3 RETURNING id;`)
	if err != nil {
		return err
	}
	stmtUpdateEmail, err := s.DB.Preparex(`UPDATE cmn.Emails SET OTP = $2, Email = $3, SendedDate = $4, Used = $5, Queued = $6 WHERE id = $1;`)
	if err != nil {
		return err
	}
	stmpSelectEmailForSending, err := s.DB.Preparex(`SELECT id, userid, otp, email AS destination FROM cmn.Emails WHERE SendedDate IS NULL AND (Queued IS NULL OR Queued = false);`)
	if err != nil {
		return err
	}

	s.preparedStatements["RegUser"] = stmtRegUser
	s.preparedStatements["UpdateUser"] = stmtUpdateUser
	s.preparedStatements["AuthUser"] = stmtAuthUser
	s.preparedStatements["VerifyUser"] = stmtVerifyUser
	s.preparedStatements["AddByte"] = stmtAddByte
	s.preparedStatements["SelectByte"] = stmtSelectByte
	s.preparedStatements["UpdateByte"] = stmtUpdateByte
	s.preparedStatements["DeleteByte"] = stmtDeleteByte
	s.preparedStatements["StartSession"] = stmtStartSession
	s.preparedStatements["UpdateSession"] = stmtUpdateSession
	s.preparedStatements["EndSession"] = stmtEndSession
	s.preparedStatements["AddEmail"] = stmtAddEmail
	s.preparedStatements["UpdateEmail"] = stmtUpdateEmail
	s.preparedStatements["SelectEmailForSending"] = stmpSelectEmailForSending

	return nil
}
