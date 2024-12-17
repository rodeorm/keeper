package repo

import (
	"context"
	"fmt"
	"log"

	"github.com/rodeorm/keeper/internal/core"
	"github.com/rodeorm/keeper/internal/crypt"
	"github.com/rodeorm/keeper/internal/logger"
	"go.uber.org/zap"
)

// RegUser создает пользователя в БД
func (s *postgresStorage) RegUser(ctx context.Context, u *core.User) error {
	passwordHash, err := crypt.HashPassword(u.Password)
	if err != nil {
		return err
	}

	err = s.preparedStatements["RegUser"].GetContext(ctx, &u.ID, u.Login, u.Email, u.Name, u.Phone, passwordHash)
	if err != nil {
		return err
	}

	// Добавление аутентификационных сообщений не мешает регистрации, чтобы можно было подтвердить адрес и позднее. Время жизни токена ограничено.
	// Поэтому в транзакцию не завернуто, и ошибка не возвращается
	m, err := core.NewAuthMessage(u)
	if err != nil {
		logger.Log.Error("Ошибка при попытке получить новое авторизационное сообщение при регистрации пользователя",
			zap.String(u.Login, fmt.Sprintf("Для email %s: %v", u.Email, err)),
		)
	}

	err = s.AddMessage(ctx, m)
	if err != nil {
		logger.Log.Error("Ошибка при попытке добавить в БД новое авторизационное сообщение при регистрации пользователя",
			zap.String(u.Login, fmt.Sprintf("Для email %s: %v", u.Email, err)),
		)
	}
	return nil
}

// Аутентифицирует пользователя на основании данных в БД и возвращает все его данные
func (s *postgresStorage) AuthUser(ctx context.Context, u *core.User) bool {
	pass := u.Password
	err := s.preparedStatements["AuthUser"].GetContext(ctx, u, u.Login)

	log.Println(u)

	if err != nil {
		logger.Log.Error("Ошибка при аутентификации пользователя",
			zap.String(u.Login, err.Error()),
		)
		return false
	}
	authenticated := crypt.CheckPasswordByHash(pass, u.Password)
	u.Password = pass // Возвращаем обратно пароль, а не хэш. Он понадобится для шифрования данных пользователя

	if authenticated {
		logger.Log.Info("Пользователь прошел аутентификацию успешно",
			zap.String(u.Login, fmt.Sprintf("С идентификатором: %d", u.ID)),
		)
	}

	return authenticated
}

func (s *postgresStorage) VerifyUserOTP(ctx context.Context, otpLiveTime int, u *core.User) bool {
	// SELECT e.id FROM cmn.emails AS e WHERE e.UserID = $1 AND e.sendeddate + ($2 * INTERVAL '1 hour') > NOW()
	// AND e.Used = false AND e.OTP = $3;
	em := &core.Message{}
	err := s.preparedStatements["VerifyUser"].GetContext(ctx, em, u.ID, otpLiveTime, u.OTP)
	if err != nil {
		logger.Log.Error("не удалось верифицировать OTP",
			zap.String(u.Login, fmt.Sprintf("одноразовый пароль: %s", u.OTP)),
		)
		return false
	}
	if em.ID == 0 {
		return false
	}
	return true
}

func (s *postgresStorage) UpdateUser(context.Context, *core.User) error {

	return nil
}
func (s *postgresStorage) DeleteUser(context.Context, *core.User) error {
	return nil
}
