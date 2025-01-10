package database

import (
	"watdowedo/internal/common/logger"

	"github.com/jackc/pgx/v5/pgxpool"
)

func DeployDbFromFile(dbPool *pgxpool.Pool, scriptPath string, test bool) error {
	result, err := ExecScriptFile(dbPool, scriptPath)
	if err != nil {
		return err
	}
	if !test {
		logger.GlobalLogger.Info("Database deployment successfull\n" + result.String())
	}
	return nil
}
