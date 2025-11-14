package main

import (
	"log/slog"
	"os"

	"github.com/tmozzze/QuesAns/internal/config"
	"github.com/tmozzze/QuesAns/internal/storage"
	"github.com/tmozzze/QuesAns/internal/storage/migration"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// Load config
	cfg := config.MustLoad()

	// Setup logger
	log := setupLogger(cfg.Env)

	log.Info("starting QuesAns", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	// Init Postgres
	storage, err := storage.New(cfg, log)
	if err != nil {
		log.Error("failed to init storage", slog.Any("err", err))
		os.Exit(1)
	}

	// Database access
	sqlDB, err := storage.Postgres.DB.DB()
	if err != nil {
		log.Error("failed to connect Postgres", slog.Any("err", err))
	}
	defer sqlDB.Close()

	// Migrations
	if err := migration.ApplyMigrations(sqlDB, cfg, log); err != nil {
		log.Error("failed to apply migrations", slog.Any("err", err))
		os.Exit(1)
	}

	// TODO: init router: net/http

	// TODO: run server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal: // Text Debug
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev: // JSON Debug
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd: // JSON Info
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
