package postgres

import (
	"fmt"
	"log/slog"
	"time"

	"gorm.io/gorm/logger"
)

type slogWriter struct {
	log *slog.Logger
}

func (w *slogWriter) Printf(format string, args ...interface{}) {
	w.log.Info(fmt.Sprintf(format, args...))
}

func NewGormLogger(log *slog.Logger) logger.Interface {
	// GORM-logger with Writer
	gormLogger := logger.New(
		&slogWriter{log: log},
		logger.Config{
			SlowThreshold: time.Second, // logging query longer than 1 sec
			LogLevel:      logger.Info,
			Colorful:      false,
		},
	)
	return gormLogger
}
