package usecase

import (
	"github.com/AntonioSchappo/goexpert-clean-arch-challenge/internal/entity"
	"github.com/AntonioSchappo/goexpert-clean-arch-challenge/pkg/events"
)

type OrderCreateUseCase struct {
	OrderRepository entity.IOrderRepository
	OrderCreated    events.IEvent
	EventDispatcher events.IEventDispatcher
}

func NewOrderCreateUseCase(repository entity.IOrderRepository,
	orderCreated events.IEvent,
	eventDispatcher events.IEventDispatcher) *OrderCreateUseCase {
	return &OrderCreateUseCase{
		OrderRepository: repository,
		OrderCreated:    orderCreated,
		EventDispatcher: eventDispatcher,
	}
}

func (u *OrderCreateUseCase) Execute(input OrderInputDTO) (OrderOutputDTO, error) {
	order := entity.Order{ID: input.ID, Price: input.Price, Tax: input.Tax}
	order.CalculateFinalPrice()
	if err := u.OrderRepository.Save(&order); err != nil {
		return OrderOutputDTO{}, err
	}
	dto := OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}

	u.OrderCreated.SetPayload(dto)
	u.EventDispatcher.Dispatch(u.OrderCreated)

	return dto, nil
}
