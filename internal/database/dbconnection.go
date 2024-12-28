package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	var err error
	if err = godotenv.Load("../../.env"); err != nil {
		return nil, err
	}

	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("DATABASE_NAME")
	dbUser := os.Getenv("DATABASE_USER")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbMaxRetry := os.Getenv("DATABASE_MAX_RETRY")
	var dbMaxRetryInt int
	dbRetryDelay := os.Getenv("DATABASE_RETRY_DELAY")
	var dbRetryDelayInt int

	if dbMaxRetry != "" {
		dbMaxRetryInt, err = strconv.Atoi(dbMaxRetry)
		if err != nil {
			return nil, err
		}
	}

	if dbRetryDelay != "" {
		dbRetryDelayInt, err = strconv.Atoi(dbRetryDelay)
		if err != nil {
			return nil, err
		}
	}

	psqlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	for i := 0; i < dbMaxRetryInt; i++ {
		db, err := sql.Open("postgres", psqlConn)
		if err == nil {
			if err := db.Ping(); err == nil {
				return db, nil
			}
		}

		time.Sleep(time.Duration(dbRetryDelayInt) * time.Millisecond)
	}

	return nil, err
}
