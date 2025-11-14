package migration

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/pressly/goose"
	"github.com/tmozzze/QuesAns/internal/config"
)

type GooseMigrator struct {
	db  *sql.DB
	cfg *config.Config
	log *slog.Logger
}

func NewGooseMigrator(db *sql.DB, cfg *config.Config, log *slog.Logger) *GooseMigrator {
	return &GooseMigrator{db: db, cfg: cfg, log: log}
}

func (m *GooseMigrator) ApplyMigrations() error {
	const op = "storage.migrations.GooseMigrator.ApplyMigrations"
	migrationsDir := m.cfg.MigrationsDir
	dbDialect := m.cfg.DBDialect
	db := m.db
	log := m.log

	mgLog := log.With(slog.String("op", op), slog.String("MigrationsDir", migrationsDir))
	mgLog.Debug("start applying migrations")

	// set pref gor goose
	goose.SetLogger(&GooseSlogAdapter{log: log})
	goose.SetDialect(dbDialect)

	if err := goose.Up(db, migrationsDir); err != nil {
		return fmt.Errorf("%s: goose up failed: %w", op, err)
	}

	mgLog.Info("migrations applied successfully")

	return nil
}
