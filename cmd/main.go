package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/cristian0193/golang-service-template/http/server"

	"github.com/cristian0193/golang-service-template/builder"
)

func main() {
	logger := builder.NewLogger()
	logger.Info("Starting golang-service-template ... ")
	defer builder.Sync(logger)

	config, err := builder.LoadConfiguration()
	if err != nil {
		logger.Fatalf("LoadConfiguration: %v", err)
	}

	_, err = builder.NewDatabase(logger)
	if err != nil {
		logger.Fatalf("LoadConfiguration: %v", err)
	}

	// log := builder.LogLevel(*config)

	serve := server.NewServer(config.ServerPort)
	if err := serve.Start(); err != nil {
		logger.Fatalf("server.Start: %v", err)
	}

	sigQuit := make(chan os.Signal, 1)
	signal.Notify(sigQuit, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGTERM)
	sig := <-sigQuit
	logger.Infof("Shutting down server with signal [%s] ...", sig.String())

	if err := serve.Stop(); err != nil {
		logger.Fatalf("server.Stop: %v", err)
	}

	logger.Info("Stop Service...")
}
