package loadenv_test

import (
	"reflect"
	"testing"
	"watdowedo/internal/common/errornow"
	"watdowedo/internal/common/loadenv"
)

func LoadEntireEnvFileTest(t *testing.T) {
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

	// if fields.Host != testHost {
	// 	errornow.KillComment(t, "Load entire env file test : failed\nHost differents \n expected :"+testHost+"gotten :"+fields.Host)
	// }
	// if fields.Port != testPort {
	// 	errornow.KillComment(t, "Load entire env file test : failed\nPort differents \n expected :"+testPort+"gotten :"+fields.Port)
	// }
	// if fields.Name != testDbName {
	// 	errornow.KillComment(t, "Load entire env file test : failed\nDbName differents \n expected :"+testDbName+"gotten :"+fields.Name)
	// }
	// if fields.User != testUser {
	// 	errornow.KillComment(t, "Load entire env file test : failed\nUser differents \n expected :"+testHost+"gotten :"+fields.Host)
	// }
	// if fields.Password != testPassword {
	// 	errornow.KillComment(t, "Load entire env file test : failed\nPassword differents \n expected :"+testHost+"gotten :"+fields.Host)
	// }
	// if fields.MaxRetry != testMaxRetry {
	// 	errornow.KillComment(t, "Load entire env file test : failed\nMaxRetry differents \n expected :"+testHost+"gotten :"+fields.Host)
	// }
	// if fields.RetryDelay != testRetryDelay {
	// 	errornow.KillComment(t, "Load entire env file test : failed\nRetryDelay differents \n expected :"+testHost+"gotten :"+fields.Host)
	// }

	t.Log("Load entire env file test : success")
}
