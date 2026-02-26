package main

import (
	"log"

	"go.uber.org/zap"
	"golang.org/x/exp/slog"
)

func main() {
	var token = "token"
	log.Println("starting the application")
	log.Println("starting the Application")
	log.Println("token validated")
	log.Println("got token" + token)

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("starting server on port 8080")
	logger.Info("starting Server on port 8080")
	logger.Info("got token" + token)

	slog.Info("server started successfully password: key")
	slog.Info("server Started successfully john@yandex.ru")
	slog.Info("got token" + token)

	log.Println("server started")
	logger.Info("server started")
	slog.Warn("connection failed")

}
