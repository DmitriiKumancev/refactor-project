package main

import (
	"log"
	"net/http"

	"github.com/DmitriiKumancev/refactor-project/api"
	"github.com/DmitriiKumancev/refactor-project/storage"
	"github.com/DmitriiKumancev/refactor-project/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	// Инициализация логгера
	if err := logger.InitLogger(); err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	storage.InitStore()

	r := api.NewRouter()
	http.Handle("/", r)

	addr := ":3333"
	logger.GetLogger().Info("Server is running", zap.String("address", addr))
	if err := http.ListenAndServe(addr, nil); err != nil {
		logger.GetLogger().Fatal("Failed to start server", zap.Error(err))
	}
}
