package router

import (
	"context"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/config"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/database/in_memory/repositories"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/http/render"
)

func NewRouter(ctx context.Context, wg *sync.WaitGroup, env *config.Env) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)

	// prismaClient, err := prisma.NewClient(ctx, wg, prisma.PrismaConfig{
	// 	DBUser:     env.DBUser,
	// 	DBPassword: env.DBPassword,
	// 	DBHost:     env.DBHost,
	// 	DBPort:     env.DBPort,
	// 	DBName:     env.DBName,
	// })

	// if err != nil {
	// 	panic(err)
	// }

	// productsRepository := repositories.NewPrismaProductsRepository(ctx, prismaClient)
	// cartsRepository := repositories.NewPrismaCartsRepository(ctx, prismaClient)

	productsRepository := repositories.NewInMemoryProductsRepository()
	cartsRepository := repositories.NewInMemoryCartsRepository()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		render.NewSuccess("OK", http.StatusOK).Send(w)
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/products", NewProductRoutes(productsRepository))
			r.Route("/carts", NewCartRoutes(cartsRepository, productsRepository))
		})
	})

	return r
}
