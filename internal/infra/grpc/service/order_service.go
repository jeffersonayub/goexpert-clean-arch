package service

import (
	"context"

	"github.com/jeffersonayub/goexpert-clean-arch/internal/infra/grpc/pb"
	"github.com/jeffersonayub/goexpert-clean-arch/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrderUseCase   usecase.ListOrderUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, listOrderUseCase usecase.ListOrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		ListOrderUseCase:   listOrderUseCase,
	}
}

func (o *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := o.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (o *OrderService) ListOrders(ctx context.Context, in *pb.Blank) (*pb.OrdersList, error) {
	orders, err := o.ListOrderUseCase.Execute()
	if err != nil {
		return nil, err
	}

	var listOrdersResponse []*pb.Order

	for _, order := range orders {
		createOrderResponse := &pb.Order{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		}
		listOrdersResponse = append(listOrdersResponse, createOrderResponse)
	}

	return &pb.OrdersList{
		Orders: listOrdersResponse,
	}, nil
}
