package usecase

import (
	"github.com/jeffersonayub/goexpert-clean-arch/internal/entity"
	"github.com/jeffersonayub/goexpert-clean-arch/pkg/events"
)

type ListOrdersOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrderCreated    events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewListOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrderCreated events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
		OrderCreated:    OrderCreated,
		EventDispatcher: EventDispatcher,
	}
}

func (l *ListOrderUseCase) Execute() ([]ListOrdersOutputDTO, error) {
	items, err := l.OrderRepository.List()
	if err != nil {
		return nil, err
	}

	var orders []ListOrdersOutputDTO
	for _, order := range items {
		orders = append(orders, ListOrdersOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	return orders, nil
}
