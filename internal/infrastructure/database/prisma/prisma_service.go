package prisma

import (
	"context"
	"fmt"
	"sync"

	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infrastructure/config"
	"github.com/natasha-m-oliveira/clean-architecture-go/prisma/db"
)

func Init(ctx context.Context, wg *sync.WaitGroup) (*db.PrismaClient, error) {
	config := config.Config
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
