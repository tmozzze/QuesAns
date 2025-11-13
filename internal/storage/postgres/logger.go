package postgres

import (
	"log/slog"
	"time"

	"gorm.io/gorm/logger"
)

func NewGormLogger(log *slog.Logger) logger.Interface {
	lvl := gormLogLevelFromSlog(log)
	// GORM-logger
	gormLogger := logger.New(
		slog.NewLogLogger(log.Handler(), slog.LevelInfo),
		logger.Config{
			SlowThreshold: time.Second, // logging query longer than 1 sec
			LogLevel:      lvl,
			Colorful:      false,
		},
	)
	return gormLogger
}

// gormLogLevelFromSlog — mapping slog level to log level
func gormLogLevelFromSlog(log *slog.Logger) logger.LogLevel {
	var lvl logger.LogLevel

	switch {
	case log.Handler().Enabled(nil, slog.LevelDebug):
		lvl = logger.Info // SQL queries
	default:
		lvl = logger.Error // Error
	}

	return lvl
}
