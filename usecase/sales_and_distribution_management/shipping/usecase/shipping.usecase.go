package shipping_usecase

import (
	"context"
	shippingdomain "shop_erp_mono/domain/sales_and_distribution_management/shipping"
	"time"
)

type shippingUseCase struct {
	contextTimeout     time.Duration
	shippingRepository shippingdomain.IShippingRepository
}

func NewShippingUseCase(contextTimeout time.Duration, shippingRepository shippingdomain.IShippingRepository) shippingdomain.IShippingUseCase {
	return &shippingUseCase{contextTimeout: contextTimeout, shippingRepository: shippingRepository}
}

func (s *shippingUseCase) CreateOne(ctx context.Context, input *shippingdomain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (s *shippingUseCase) GetByID(ctx context.Context, id string) (*shippingdomain.ShippingResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *shippingUseCase) GetByOrderID(ctx context.Context, orderID string) (*shippingdomain.ShippingResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *shippingUseCase) UpdateOne(ctx context.Context, id string, updatedShipping *shippingdomain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (s *shippingUseCase) DeleteOne(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (s *shippingUseCase) List(ctx context.Context) ([]shippingdomain.ShippingResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *shippingUseCase) GetByStatus(ctx context.Context, status string) ([]shippingdomain.ShippingResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *shippingUseCase) UpdateDeliveryStatus(ctx context.Context, id string, status string, actualDelivery *time.Time) error {
	//TODO implement me
	panic("implement me")
}
