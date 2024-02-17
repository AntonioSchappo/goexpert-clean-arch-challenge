package usecase

import "github.com/AntonioSchappo/goexpert-clean-arch-challenge/internal/entity"

type OrderListUseCase struct {
	OrderRepository entity.IOrderRepository
}

func NewOrderListUseCase(orderRepository entity.IOrderRepository) *OrderListUseCase {
	return &OrderListUseCase{
		OrderRepository: orderRepository,
	}
}

func (u *OrderListUseCase) Execute() ([]OrderOutputDTO, error) {
	orderList, err := u.OrderRepository.ListOrders()
	if err != nil {
		return []OrderOutputDTO{}, err
	}
	dtoList := []OrderOutputDTO{}
	for _, order := range orderList {
		dto := OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
		dtoList = append(dtoList, dto)
	}
	return dtoList, nil
}
