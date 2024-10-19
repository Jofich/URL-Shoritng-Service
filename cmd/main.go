package main

import (
	"log/slog"
	"os"

	"github.com/Jofich/URL-Shoritng-Service/internal/config"
	"github.com/Jofich/URL-Shoritng-Service/internal/lib/logger/sl"
	"github.com/Jofich/URL-Shoritng-Service/internal/server"
	storage "github.com/Jofich/URL-Shoritng-Service/internal/storage/postgres"
)

func main() {
	//TODO: parse config
	cfg := config.MustLoad()
	//TODO: Init bd

	database, err := storage.New(cfg.StorageCfg)
	if err != nil {
		slog.Error("Failed to connect to the storage", sl.Error(err))
		os.Exit(1)
	}
	server.Start(&cfg.HTTPServer, database)
}
