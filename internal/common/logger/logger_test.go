package logger_test

import (
	"os"
	"testing"

	"gobackup/internal/common/editstring"
	"gobackup/internal/common/errornow"
	"gobackup/internal/common/logger"
)

func TestLogger(t *testing.T) {
	t.Parallel()

	writeLogers, err := logger.Logger(false)
	if err != nil {
		errornow.KillComment(t, err)
	}

	writeLogers.Error("verbose logger test")

	body, err := os.ReadFile("logs/logs.log")
	if err != nil {
		errornow.KillComment(t, err)
	}

	if editstring.Clean(string(body)) != "" {
		t.Log("non verbose logger test : success")
	}
}
