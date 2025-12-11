package loadenv

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEntireEnvFile(envFilePath string) (DotEnvFields, error) {
	db, err := LoadDbFields(envFilePath)
	if err != nil {
		return DotEnvFields{}, err
	}

	envVar := DotEnvFields{
		db.Host,
		db.Port,
		db.Name,
		db.User,
		db.Password,
		db.MaxRetry,
		db.RetryDelay,
	}

	return envVar, nil
}

func LoadDbFields(envFilePath string) (DbFields, error) {
	if err := godotenv.Load(envFilePath); err != nil {
		return DbFields{}, err
	}

	dbEnvVar, err := FormatDbFieldsToStruct(
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_MAX_RETRY"),
		os.Getenv("DATABASE_RETRY_DELAY"),
	)

	if err != nil {
		return DbFields{}, err
	}
	return dbEnvVar, nil
}
