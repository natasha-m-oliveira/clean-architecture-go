package main

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"

	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/config"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/http/router"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/http/server"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	var wg sync.WaitGroup
	env, err := config.Load()
	if err != nil {
		panic(fmt.Sprintf("Failed to load config from environment variables: %v", err.Error()))
	}

	router := router.NewRouter(ctx, &wg, env)

	server := server.NewServer(router).WithPort(env.HttpServerPort)

	server.Start(ctx, &wg)

	<-ctx.Done()
	server.Shutdown(ctx)

	wg.Wait()
}
