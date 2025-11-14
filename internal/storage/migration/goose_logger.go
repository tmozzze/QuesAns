package migration

import (
	"fmt"
	"log/slog"
)

type GooseSlogAdapter struct {
	log *slog.Logger
}

// Print — message
func (g *GooseSlogAdapter) Print(v ...interface{}) {
	g.log.Info(fmt.Sprint(v...))
}

// Printf — format message
func (g *GooseSlogAdapter) Printf(format string, v ...interface{}) {
	g.log.Info(fmt.Sprintf(format, v...))
}

// Println — Info
func (g *GooseSlogAdapter) Println(v ...interface{}) {
	g.log.Info(fmt.Sprint(v...))
}

// Fatal — error and panic
func (g *GooseSlogAdapter) Fatal(v ...interface{}) {
	g.log.Error(fmt.Sprint(v...))
	panic(fmt.Sprint(v...))
}

// Fatalf — format error and panic
func (g *GooseSlogAdapter) Fatalf(format string, v ...interface{}) {
	g.log.Error(fmt.Sprintf(format, v...))
	panic(fmt.Sprintf(format, v...))
}
