package migration

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/pressly/goose"
	"github.com/tmozzze/QuesAns/internal/config"
)

func ApplyMigrations(db *sql.DB, cfg *config.Config, log *slog.Logger) error {
	const op = "storage.migrations.ApplyMigrations"
	migrationsDir := cfg.MigrationsDir
	dbDialect := cfg.DBDialect

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
