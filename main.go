package main

import (
	"log/slog"

	"github.com/trianglehasfoursides/totoro/server"
)

func main() {
	if err := server.Run(); err != nil {
		slog.Error(err.Error())
	}
}
