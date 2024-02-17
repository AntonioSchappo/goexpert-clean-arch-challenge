//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/AntonioSchappo/goexpert-clean-arch-challenge/internal/entity"
	"github.com/AntonioSchappo/goexpert-clean-arch-challenge/internal/event"
	"github.com/AntonioSchappo/goexpert-clean-arch-challenge/internal/infra/database"
	"github.com/AntonioSchappo/goexpert-clean-arch-challenge/internal/infra/web"
	"github.com/AntonioSchappo/goexpert-clean-arch-challenge/internal/usecase"
	"github.com/AntonioSchappo/goexpert-clean-arch-challenge/pkg/events"
	"github.com/google/wire"
)

var setRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.IOrderRepository), new(*database.OrderRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewCreatedOrder,
	wire.Bind(new(events.IEventDispatcher), new(*events.EventDispatcher)),
	wire.Bind(new(events.IEvent), new(*event.CreatedOrder)),
)

var setOrderCreatedEvent = wire.NewSet(
	event.NewCreatedOrder,
	wire.Bind(new(events.IEvent), new(*event.CreatedOrder)),
)

func NewOrderCreateUseCase(db *sql.DB, eventDispatcher events.IEventDispatcher) *usecase.OrderCreateUseCase {
	wire.Build(
		setRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewOrderCreateUseCase,
	)
	return &usecase.OrderCreateUseCase{}
}

func NewOrderListUseCase(db *sql.DB) *usecase.OrderListUseCase {
	wire.Build(
		setRepositoryDependency,
		usecase.NewOrderListUseCase,
	)
	return &usecase.OrderListUseCase{}
}

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.IEventDispatcher) *web.WebOrderHandler {
	wire.Build(
		setRepositoryDependency,
		setOrderCreatedEvent,
		web.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}
