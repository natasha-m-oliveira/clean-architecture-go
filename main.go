package main

import (
	"context"
	"log"
	"os/signal"
	"sync"
	"syscall"

	"github.com/natasha-m-oliveira/clean-architecture-go/infrastructure/config"
	"github.com/natasha-m-oliveira/clean-architecture-go/infrastructure/http/server"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	var wg sync.WaitGroup
	if err := config.Load(); err != nil {
		log.Fatalf("Failed to load config from environment variables: %v", err.Error())
	}

	server.Init(ctx, &wg).Start(ctx, &wg)

	wg.Wait()
}
