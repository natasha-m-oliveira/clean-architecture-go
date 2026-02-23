package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/app/repositories"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/app/usecases"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/http/controllers"
)

func NewCartRoutes(cartsRepository repositories.CartsRepository, productsRepository repositories.ProductsRepository) func(r chi.Router) {
	return func(r chi.Router) {
		createCartUseCase := usecases.NewCreateCartUseCase(cartsRepository, productsRepository)
		getCartByIdUseCase := usecases.NewGetCartByIdUseCase(cartsRepository)
		deleteCartUseCase := usecases.NewDeleteCartUseCase(cartsRepository)
		updateCartItemsUseCase := usecases.NewUpdateCartItemsUseCase(cartsRepository, productsRepository)

		createCartController := controllers.NewCreateCartController(createCartUseCase)
		getCartByIdController := controllers.NewGetCartByIdController(getCartByIdUseCase)
		deleteCartController := controllers.NewDeleteCartController(deleteCartUseCase)
		updateCartItemsController := controllers.NewUpdateCartItemsController(updateCartItemsUseCase)

		r.Post("/", createCartController.Execute)
		r.Get("/{id}", getCartByIdController.Execute)
		r.Delete("/{id}", deleteCartController.Execute)
		r.Put("/{id}/items", updateCartItemsController.Execute)
	}
}
