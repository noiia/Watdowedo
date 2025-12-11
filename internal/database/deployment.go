package database

import (
	"watdowedo/internal/common/logger"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Deploy a database from a sql script file.
//
// Parameters :
//   - dbPool : database pool
//   - scriptPath : string
//   - Testmode : True if test mode activated
//
// Returns :
//
//   - error : potential errors from dpPool creation
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
