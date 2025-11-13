package main

import (
	"fmt"

	"github.com/tmozzze/QuesAns/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// TODO: init logger: logrus

	// TODO: init storage: postgres

	// TODO: init router: net/http

	// TODO: run server
}
