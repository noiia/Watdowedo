package database

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"watdowedo/internal/common/loadenv"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lib/pq"
)

func ExecScriptFile(dbPool *pgxpool.Pool, scriptPath string) (pgconn.CommandTag, error) {
	sqlQuery, err := os.ReadFile(scriptPath)
	if err != nil {
		return pgconn.CommandTag{}, err
	}

	return dbPool.Exec(context.Background(), string(sqlQuery))
}

func DeleteDB(dbPool *pgxpool.Pool, dbFields loadenv.DbFields) error {
	exec.Command("psql -U %s -h %s -p %s -c 'DROP DATABASE %s'", dbFields.User, dbFields.Host, dbFields.Port, dbFields.Name)
	return nil
}

func DeleteTable(dbPool *pgxpool.Pool, tableName string) error {
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	query := fmt.Sprintf("DROP TABLE IF EXISTS %s", pq.QuoteIdentifier(tableName))

	_, err := dbPool.Exec(context.Background(), query)
	if err != nil {
		return err
	}

	return nil
}

func DeleteRow(dbPool *pgxpool.Pool, rowName string) error {

	return nil
}

func TableExists(dbPool *pgxpool.Pool, schemaName, tableName string) (bool, error) {
	query := `
		SELECT EXISTS (
			SELECT 1
			FROM   information_schema.tables
			WHERE  table_schema = $1
			AND    table_name = $2
		)
	`

	var exists bool
	err := dbPool.QueryRow(context.Background(), query, schemaName, tableName).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
