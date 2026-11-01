package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "morse\t", log.LstdFlags|log.Lshortfile)

	if err := server.New(logger).Run(); err != nil {
		logger.Fatalf("сервер остановлен: %v", err)
	}
}
