package prisma

import (
	"context"
	"fmt"
	"sync"

	"github.com/natasha-m-oliveira/clean-architecture-go/prisma/db"
)

type PrismaConfig struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

func NewClient(ctx context.Context, wg *sync.WaitGroup, config PrismaConfig) (*db.PrismaClient, error) {
	url := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)

	client := db.NewClient(db.WithDatasourceURL(url))
	if err := client.Prisma.Connect(); err != nil {
		return nil, err
	}

	wg.Add(1)

	go func() {
		<-ctx.Done()

		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}

		wg.Done()
	}()

	return client, nil
}
