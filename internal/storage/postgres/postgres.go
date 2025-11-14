package postgres

import (
	"fmt"
	"log/slog"

	"github.com/tmozzze/QuesAns/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	DB *gorm.DB
}

func New(cfg config.PostgresCfg, log *slog.Logger) (*Storage, error) {
	const op = "storage.postgres.New"

	pgLog := log.With(
		slog.String("op", op),
		slog.String("host", cfg.Host),
		slog.String("db", cfg.DBName),
		slog.String("user", cfg.User),
	)
	pgLog.Debug("connecting to PostgreSQL")

	// DSN
	dsn := cfg.DSN()

	// GORM-logger
	gormLogger := NewGormLogger(log)

	// DB connection
	db, err := gorm.Open(
		postgres.New(postgres.Config{DSN: dsn, PreferSimpleProtocol: true}),
		&gorm.Config{Logger: gormLogger},
	)
	if err != nil {
		return nil, fmt.Errorf("%s: open connection failed: %w", op, err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("%s: get sql.DB failed: %w", op, err)
	}

	// set env
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	// Ping DB
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("%s: ping failed: %w", op, err)
	}

	pgLog.Info("connected to PostgreSQL successfully")

	return &Storage{DB: db}, nil

}
