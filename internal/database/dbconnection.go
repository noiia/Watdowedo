package database

import (
	"context"
	"fmt"
	"watdowedo/internal/common/loadenv"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

// Create a database connexion pool to a dedicated database by giving a .env file path with server variables.
//
// Parameters :
//   - envFilePath : string
//
// Returns :
//   - dbFields : Loaded env structure from loadenv.DbFields
//   - dbPool : database pool
//   - error : potential errors from dpPool creation
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

// Create a database pool to dbFields loaded env variable.
//
// Parameters :
//   - dbFields (Loaded env structure from loadenv.DbFields) : struct DbFields
//
// Returns :
//   - dbPool : database pool
//   - error : potential errors from dpPool creation
func CreateDbPool(dbFields loadenv.DbFields) (*pgxpool.Pool, error) {
	databaseUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbFields.User, dbFields.Password, dbFields.Host, dbFields.Port, dbFields.Name)

	dbpool, err := pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		return nil, err
	}

	return dbpool, nil
}
