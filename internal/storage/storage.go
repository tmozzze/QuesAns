package storage

import (
	"fmt"
	"log/slog"

	"github.com/tmozzze/QuesAns/internal/config"
	"github.com/tmozzze/QuesAns/internal/storage/postgres"
)

type Storage struct {
	Postgres *postgres.Storage
}

func New(cfg *config.Config, log *slog.Logger) (*Storage, error) {
	const op = "storage.New"
	st := &Storage{}

	// Postgres init
	pg, err := postgres.New(cfg.PostgresCfg, log)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to init PostgreSQL: %w", op, err)
	}
	st.Postgres = pg

	return st, nil
}
