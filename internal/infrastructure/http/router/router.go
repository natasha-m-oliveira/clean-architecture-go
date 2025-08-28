package router

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/adapter/database/repositories"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/adapter/http/controllers"
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

		createProductController  controllers.CreateProductController
		deleteProductController  controllers.DeleteProductController
		getProductByIdController controllers.GetProductByIdController
		listProductsController   controllers.ListProductsController
		createCartController     controllers.CreateCartController
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
	cartsRepository := repositories.NewPrismaCartsRepository(prismaClient, ctx)

	createProductUseCase := usecases.NewCreateProductUseCase(productsRepository)
	deleteProductUseCase := usecases.NewDeleteProductUseCase(productsRepository)
	getProductByIdUseCase := usecases.NewGetProductByIdUseCase(productsRepository)
	listProductsUseCase := usecases.NewListProductsUseCase(productsRepository)
	createCartUseCase := usecases.NewCreateCartUseCase(cartsRepository, productsRepository)

	engine.createProductController = controllers.NewCreateProductController(createProductUseCase)
	engine.deleteProductController = controllers.NewDeleteProductController(deleteProductUseCase)
	engine.getProductByIdController = controllers.NewGetProductByIdController(getProductByIdUseCase)
	engine.listProductsController = controllers.NewListProductsController(listProductsUseCase)
	engine.createCartController = controllers.NewCreateCartController(createCartUseCase)

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
			fmt.Printf("Error starting HTTP server: %v\n", err)
			return
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

	v1 := router.Group("/v1")
	{
		products := v1.Group("/products")
		{
			products.POST("/", engine.wrapper(engine.createProductController.Execute))
			products.DELETE("/:id", engine.wrapper(engine.deleteProductController.Execute))
			products.GET("/:id", engine.wrapper(engine.deleteProductController.Execute))
			products.GET("/", engine.wrapper(engine.listProductsController.Execute))
		}

		carts := v1.Group("/carts")
		{
			carts.POST("/", engine.wrapper(engine.createCartController.Execute))
		}
	}

}

func (engine *ginEngine) wrapper(callback func(w http.ResponseWriter, r *http.Request)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rw := ctx.Writer
		req := ctx.Request

		reconstructedPath := ctx.FullPath()
		for _, param := range ctx.Params {
			reconstructedPath = strings.Replace(reconstructedPath, ":"+param.Key, param.Value, 1)
		}

		req.URL.Path = reconstructedPath

		callback(rw, req)
	}
}
