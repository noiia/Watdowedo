package database_test

import (
	"path/filepath"
	"testing"
	"watdowedo/internal/common/errornow"
	"watdowedo/internal/database"
)

func TestDbDeploymentFromFile(t *testing.T) {
	t.Parallel()

	dbScriptFilePath := filepath.Join(".", "..", "..", "test", "database_test.sql")
	dbEnvFilePath := filepath.Join(".", "..", "..", "test", ".env_test")

	dbFields, dbPool, err := database.ConnectWithEnvFile(dbEnvFilePath)
	if err != nil {
		errornow.KillComment(t, err.Error())
	}

	defer dbPool.Close()

	if err := database.DeleteTable(dbPool, "type_activities"); err != nil {
		errornow.KillComment(t, err.Error())
	}

	if err = database.DeployDbFromFile(dbPool, dbScriptFilePath, true); err != nil {
		errornow.KillComment(t, err.Error())
	}

	exists, err := database.TableExists(dbPool, "public", "type_activities")
	if err != nil {
		errornow.KillComment(t, err.Error())
	}

	if !exists {
		errornow.KillComment(t, "Deployment failed, exist table test failed")
	}

	if err := database.DeleteDB(dbPool, dbFields); err != nil {
		errornow.KillComment(t, err.Error())
	}

	t.Log("Deployment test : successfull")
}
