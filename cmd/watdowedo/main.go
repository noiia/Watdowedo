package main

import (
	"flag"

	"watdowedo/internal/common/logger"
	"watdowedo/internal/common/routing"
)

func main() {
	verbose := flag.Bool("v", false, "verbose output in terminal")

	flag.Parse()

	err := logger.InitLogger(*verbose)
	if err != nil {
		panic(err)
	}

	server := routing.Routing()

	logger.GlobalLogger.Info("starting server at http://localhost" + server.Addr + "/home")

	if err := server.ListenAndServe(); err != nil {
		logger.GlobalLogger.Error("Server internal error : " + err.Error())
	}

}
