package service

import (
	"context"

	"github.com/AntonioSchappo/goexpert-clean-arch-challenge/internal/infra/grpc/pb"
	"github.com/AntonioSchappo/goexpert-clean-arch-challenge/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUsecase usecase.OrderCreateUseCase
	ListOrdersUsecase  usecase.OrderListUseCase
}

func NewOrderService(createOrderUsecase usecase.OrderCreateUseCase, listOrdersUsecase usecase.OrderListUseCase) *OrderService {
	return &OrderService{
		CreateOrderUsecase: createOrderUsecase,
		ListOrdersUsecase:  listOrdersUsecase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUsecase.Execute(dto)
	if err != nil {
		return nil, err
	}

	return &pb.OrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrders(ctx context.Context, in *pb.Blank) (*pb.OrderListResponse, error) {
	ordersListDto, err := s.ListOrdersUsecase.Execute()
	if err != nil {
		return nil, err
	}

	var orderListResponse []*pb.OrderResponse

	for _, order := range ordersListDto {
		orderResponse := &pb.OrderResponse{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		}
		orderListResponse = append(orderListResponse, orderResponse)
	}

	return &pb.OrderListResponse{Orders: orderListResponse}, nil
}
