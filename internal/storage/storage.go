package storage

import (
	"fmt"
	"log/slog"

	"github.com/tmozzze/QuesAns/internal/config"
	repoPg "github.com/tmozzze/QuesAns/internal/repository/postgres"
	stPg "github.com/tmozzze/QuesAns/internal/storage/postgres"
)

type Storage struct {
	Postgres *stPg.Storage      // storage/postgres
	Repos    *repoPg.Repository // repository/postgres
}

func New(cfg *config.Config, log *slog.Logger) (*Storage, error) {
	const op = "storage.New"
	st := &Storage{}

	stLog := log.With(slog.String("op", op), slog.String("StoragePath", cfg.StoragePath))
	stLog.Debug("start init PostgreSQL and Repos")

	// Postgres init
	pg, err := stPg.New(cfg.PostgresCfg, log)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to init PostgreSQL: %w", op, err)
	}

	repos := repoPg.NewRepository(pg.DB)

	// fill Storage struct
	st.Postgres = pg
	st.Repos = repos

	stLog.Debug("init PostgreSQL and Repos finished")

	return st, nil
}
