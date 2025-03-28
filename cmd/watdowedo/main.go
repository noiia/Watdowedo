package main

import (
	"flag"
	"path/filepath"

	"watdowedo/internal/common/logger"
	"watdowedo/internal/common/routing"
	"watdowedo/internal/database"
)

func main() {
	ENV_PATH := filepath.Join(".", ".env")
	DB_SCRIPT_PATH := filepath.Join(".", "internal", "database", "scripts", "database.sql")

	verbose := flag.Bool("v", false, "verbose output in terminal")

	flag.Parse()

	err := logger.InitLogger(*verbose)
	if err != nil {
		panic(err)
	}

	_, dbPool, err := database.ConnectWithEnvFile(ENV_PATH)
	if err != nil {
		logger.GlobalLogger.Error(err.Error())
	}
	defer dbPool.Close()

	if err = database.DeployDbFromFile(dbPool, DB_SCRIPT_PATH, true); err != nil {
		logger.GlobalLogger.Error(err.Error())
	}

	exists, err := database.TableExists(dbPool, "public", "trips")
	if err != nil {
		logger.GlobalLogger.Error(err.Error())
	}

	if !exists {
		server := routing.Routing()
		defer server.Close()

		logger.GlobalLogger.Info("starting server at http://localhost" + server.Addr + "/home")

		if err := server.ListenAndServe(); err != nil {
			logger.GlobalLogger.Error("Server internal error : " + err.Error())
		}
	}
}
