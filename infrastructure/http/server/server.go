package server

import (
	"context"
	"fmt"
	"strconv"
	"sync"

	appConfig "github.com/natasha-m-oliveira/clean-architecture-go/infrastructure/config"
	"github.com/natasha-m-oliveira/clean-architecture-go/infrastructure/database/sqlc/connection"

	"github.com/natasha-m-oliveira/clean-architecture-go/infrastructure/http/router"
)

type config struct {
	webServer router.Server
}

func NewConfig() *config {
	return &config{}
}

func (config *config) WithWebServer(ctx context.Context, wg *sync.WaitGroup) *config {
	intPort, err := strconv.ParseInt(appConfig.Config.HttpServerPort, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("error parsing port to int: %v", err))
	}

	// prismaClient, err := client.NewPrismaClient(ctx, wg)
	// if err != nil {
	// 	panic(fmt.Sprintf("Database connection failed: %v", err.Error()))
	// }

	sqlcConnection, err := connection.NewSqlcConnection(ctx, wg)
	if err != nil {
		panic(fmt.Sprintf("Database connection failed: %v", err.Error()))
	}

	server := router.
		NewGinServer().
		WithPort(intPort).
		WithControllers(sqlcConnection, ctx)

	fmt.Println("Router server has been successfully configured.")

	config.webServer = server

	return config
}

func Init(ctx context.Context, wg *sync.WaitGroup) *config {
	return NewConfig().WithWebServer(ctx, wg)
}

func (config *config) Start(ctx context.Context, wg *sync.WaitGroup) {
	config.webServer.Listen(ctx, wg)
}
