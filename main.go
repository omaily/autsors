package main

import (
	"context"
	"fmt"
	"log/slog"

	"june+/config"
)

func main() {
	fmt.Println("hello my friend")

	conf := config.MustLoad()

	serv, err := NewServer(context.Background(), conf)
	if err != nil {
		slog.Error("could not initialize server: %w", err)
		return
	}

	serv.Start()
}
