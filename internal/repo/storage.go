package repo

import (
	"sync"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/rodeorm/keeper/internal/logger"
)

// Реализация хранилища в СУБД Postgres
type postgresStorage struct {
	ConnectionString   string                // 16 байт. Строка подключения из конфиг.файла
	DB                 *sqlx.DB              // 8 байт (только указатель). Драйвер подключения к СУБД
	preparedStatements map[string]*sqlx.Stmt //8 байт (только указатель)
}

// PostgresStorage возвращает хранилище данных в Postgres (создает, если его не было ранее)
func PostgresStorage(connectionString string) (*postgresStorage, error) {
	var (
		dbErr error
		db    *sqlx.DB
		ps    *postgresStorage
		once  sync.Once
	)
	once.Do(
		func() {
			db, dbErr = sqlx.Open("pgx", connectionString)
			if dbErr != nil {
				return
			}
			ps = &postgresStorage{DB: db, ConnectionString: connectionString, preparedStatements: map[string]*sqlx.Stmt{}}
			/*
				ctx := context.TODO()
					if dbErr = ps.createTables(ctx); dbErr != nil {
						return
					}
			*/
			logger.Info("GetPostgresStorage", "Postgres storage", connectionString)
		})

	if dbErr != nil {
		logger.Error("GetPostgresStorage", "ошибка при инициализации подключения к БД", dbErr.Error())
		return nil, dbErr
	}

	return ps, nil
}
