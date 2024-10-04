package router

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/adapter/controllers"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/adapter/repositories"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/usecases"
	"github.com/natasha-m-oliveira/clean-architecture-go/prisma/db"
)

type (
	Port int64

	Server interface {
		Listen(ctx context.Context, wg *sync.WaitGroup)
	}

	ginEngine struct {
		router *gin.Engine
		port   int64

		createProductController controllers.CreateProductController
	}
)

func NewGinServer() *ginEngine {
	return &ginEngine{
		router: gin.New(),
	}
}

func (engine *ginEngine) WithPort(port int64) *ginEngine {
	engine.port = port
	return engine
}

func (engine *ginEngine) WithControllers(prismaClient *db.PrismaClient, ctx context.Context) *ginEngine {
	productsRepository := repositories.NewPrismaProductsRepository(prismaClient, ctx)

	createProductUseCase := usecases.NewCreateProductUseCase(productsRepository)

	engine.createProductController = controllers.NewCreateProductController(createProductUseCase)

	return engine
}

func (engine *ginEngine) Listen(ctx context.Context, wg *sync.WaitGroup) {
	gin.Recovery()

	engine.setAppHandlers(engine.router)

	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Addr:         fmt.Sprintf(":%d", engine.port),
		Handler:      engine.router,
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := server.ListenAndServe(); err != nil {
			panic(fmt.Sprintf("Error starting HTTP server: %v", err))
		}
		if gin.Mode() == gin.DebugMode {
			pprof.Register(engine.router)
			go func() {
				http.ListenAndServe(":6060", nil)
			}()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(timeoutCtx); err != nil {
			panic(fmt.Sprintf("HTTP server forced to shutdown: %v", err))
		}
		fmt.Println("HTTP server exiting")
	}()
}

func (engine *ginEngine) setAppHandlers(router *gin.Engine) {
	router.GET("/health", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"status": "UP"}) })
	router.POST("/products", engine.handleCreateProduct())
}

func (engine *ginEngine) handleCreateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		engine.createProductController.Execute(ctx.Writer, ctx.Request)
	}
}
