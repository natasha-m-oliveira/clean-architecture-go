package connection

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/natasha-m-oliveira/clean-architecture-go/infrastructure/config"
)

func NewSqlcConnection(ctx context.Context, wg *sync.WaitGroup) (*pgx.Conn, error) {
	config := config.Config

	connString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", config.DBUser, config.DBPassword, config.DBName, config.DBHost, config.DBPort)

	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		return nil, err
	}

	wg.Add(1)

	go func() {
		<-ctx.Done()

		if err := conn.Close(ctx); err != nil {
			panic(err)
		}

		wg.Done()
	}()

	return conn, nil
}
