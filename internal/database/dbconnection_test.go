package database_test

import (
	"testing"
	"watdowedo/internal/common/errornow"
	"watdowedo/internal/database"
)

func TestDbConnection(t *testing.T) {
	t.Parallel()

	if _, err := database.Connect(); err != nil {
		errornow.KillComment(t, err)
	}

	t.Log("Test database connexion test : success")
}
