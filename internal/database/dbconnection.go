package database

import (
	"context"
	"fmt"
	"watdowedo/internal/common/loadenv"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

func ConnectWithEnvFile(envFilePath string) (loadenv.DbFields, *pgxpool.Pool, error) {
	dbFields, err := loadenv.LoadDbFields(envFilePath)
	if err != nil {
		return loadenv.DbFields{}, nil, err
	}

	dbpool, err := CreateDbPool(dbFields)
	if err != nil {
		return loadenv.DbFields{}, nil, err
	}

	return dbFields, dbpool, nil
}

// Connect the server to a database
//
// db : field build from addDB function
// sslMode : disable or able
func CreateDbPool(dbFields loadenv.DbFields) (*pgxpool.Pool, error) {
	databaseUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbFields.User, dbFields.Password, dbFields.Host, dbFields.Port, dbFields.Name)

	dbpool, err := pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		return nil, err
	}

	return dbpool, nil
}

// func Connect(dbFields loadenv.DbFields, sslMode string) (*sql.DB, error) {
// 	var err error
// 	var db *sql.DB

// 	psqlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s connect_timeout=10", dbFields.Host, dbFields.Port, dbFields.User, dbFields.Password, dbFields.Name, sslMode)

// 	for i := 0; i < dbFields.MaxRetry; i++ {
// 		db, err = sql.Open("postgres", psqlConn)
// 		if err == nil {
// 			if err = db.Ping(); err == nil {
// 				fmt.Println("ping success")
// 				if dbFields.Host != "watdowedo.db.test" {
// 					logger.GlobalLogger.Info("db ping successful")
// 				}

// 				return db, nil
// 			}

// 			if dbFields.Host != "watdowedo.db.test" {
// 				logger.GlobalLogger.Error("db ping error : " + err.Error())
// 			} else {
// 				fmt.Println("ping error : " + err.Error())
// 			}
// 		}
// 		time.Sleep(time.Duration(dbFields.RetryDelay) * time.Millisecond)
// 	}

// 	if db != nil && err == nil {
// 		return nil, fmt.Errorf("connection failed : db content not nil but database unreachable")
// 	}

// 	return nil, err
// }
