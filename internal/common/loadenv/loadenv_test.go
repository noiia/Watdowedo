package loadenv_test

import (
	"path/filepath"
	"reflect"
	"testing"
	"watdowedo/internal/common/errornow"
	"watdowedo/internal/common/loadenv"
)

func TestLoadEntireEnvFile(t *testing.T) {
	t.Parallel()

	envFilePath := filepath.FromSlash("/usr/src/Watdowedo/test/.env_test")

	expectedFields := []string{"watdowedo.db", "5432", "watdowedo_test", "postgres", "postgres", "10", "1000", "user@test.com", "root", "true"}

	fields, err := loadenv.LoadEntireEnvFile(envFilePath)
	if err != nil {
		errornow.KillComment(t, "Load entire env file test : failed\n"+err.Error())
	}

	val := reflect.ValueOf(fields)

	for i := 0; i < val.NumField(); i++ {
		fieldValue := val.Field(i).Interface()
		if expectedFields[i] == "10" {
			if 10 != fieldValue {
				errornow.KillComment(t)
			}
		} else if expectedFields[i] == "1000" {
			if 1000 != fieldValue {
				errornow.KillComment(t)
			}
		} else if expectedFields[i] != fieldValue {
			errornow.KillComment(t)
		}

	}

	t.Log("Load entire env file test : success")
}
