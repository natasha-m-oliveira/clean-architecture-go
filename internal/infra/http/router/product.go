package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/app/repositories"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/app/usecases"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/http/controllers"
)

func NewProductRoutes(productsRepository repositories.ProductsRepository) func(r chi.Router) {
	return func(r chi.Router) {
		createProductUseCase := usecases.NewCreateProductUseCase(productsRepository)
		getProductByIdUseCase := usecases.NewGetProductByIdUseCase(productsRepository)
		deleteProductUseCase := usecases.NewDeleteProductUseCase(productsRepository)
		listProductsUseCase := usecases.NewListProductsUseCase(productsRepository)

		productController := controllers.NewCreateProductController(createProductUseCase)
		getProductByIdController := controllers.NewGetProductByIdController(getProductByIdUseCase)
		deleteProductController := controllers.NewDeleteProductController(deleteProductUseCase)
		listProductsController := controllers.NewListProductsController(listProductsUseCase)

		r.Post("/", productController.Execute)
		r.Get("/{id}", getProductByIdController.Execute)
		r.Delete("/{id}", deleteProductController.Execute)
		r.Get("/", listProductsController.Execute)
	}
}
