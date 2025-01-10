package database_test

import (
	"testing"
	"watdowedo/internal/common/errornow"
	"watdowedo/internal/database"
)

func TestDbDeploymentFromFile(t *testing.T) {
	t.Parallel()

	const dbScriptFilePath string = "./../../test/database_test.sql"
	const dbEnvFilePath string = "./../../test/.env_test"

	dbFields, dbPool, err := database.ConnectWithEnvFile(dbEnvFilePath)
	if err != nil {
		errornow.KillComment(t, err.Error())
	}

	defer dbPool.Close()

	if err := database.DeleteTable(dbPool, "type_activities"); err != nil {
		errornow.KillComment(t, err.Error())
	}

	if err = database.DeployDbFromFile(dbPool, dbScriptFilePath, true); err != nil {
		errornow.KillComment(t, "2 : "+err.Error())
	}

	exists, err := database.TableExists(dbPool, "public", "type_activities")
	if err != nil {
		errornow.KillComment(t, "3 : "+err.Error())
	}

	if !exists {
		errornow.KillComment(t, "Deployment failed, exist table test failed")
	}

	if err := database.DeleteDB(dbPool, dbFields); err != nil {
		errornow.KillComment(t, "5 : "+err.Error())
	}

	t.Log("Deployment test : successfull")
}
