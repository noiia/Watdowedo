package loadenv_test

import (
	"reflect"
	"testing"
	"watdowedo/internal/common/errornow"
	"watdowedo/internal/common/loadenv"
)

func TestLoadEntireEnvFile(t *testing.T) {
	t.Parallel()

	const envFilePath string = "./test/.env_test"

	expectedFields := []string{"hosttest", "1111", "dbTest123", "userTest!", "testPassw0rdT3st", "20", "4987"}

	fields, err := loadenv.LoadEntireEnvFile(envFilePath)
	if err != nil {
		errornow.KillComment(t, "Load entire env file test : failed\n"+err.Error())
	}

	val := reflect.ValueOf(fields)

	for i := 0; i < val.NumField(); i++ {
		fieldValue := val.Field(i).Interface()

		if expectedFields[i] == "20" {
			if 20 != fieldValue {
				errornow.KillComment(t)
			}
		} else if expectedFields[i] == "4987" {
			if 4987 != fieldValue {
				errornow.KillComment(t)
			}
		} else if expectedFields[i] != fieldValue {
			errornow.KillComment(t)
		}

	}

	t.Log("Load entire env file test : success")
}
