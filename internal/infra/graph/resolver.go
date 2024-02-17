package graph

import "github.com/AntonioSchappo/goexpert-clean-arch-challenge/internal/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CreateOrderUsecase usecase.OrderCreateUseCase
	ListOrdersUsecase  usecase.OrderListUseCase
}
