package shipping_usecase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	shippingdomain "shop_erp_mono/domain/sales_and_distribution_management/shipping"
	"time"
)

type shippingUseCase struct {
	contextTimeout     time.Duration
	shippingRepository shippingdomain.IShippingRepository
}

func NewShippingUseCase(contextTimeout time.Duration, shippingRepository shippingdomain.IShippingRepository) shippingdomain.IShippingRepository {
	return &shippingUseCase{contextTimeout: contextTimeout, shippingRepository: shippingRepository}
}

func (s *shippingUseCase) CreateOne(ctx context.Context, shipping shippingdomain.Shipping) error {
	//TODO implement me
	panic("implement me")
}

func (s *shippingUseCase) GetByID(ctx context.Context, id primitive.ObjectID) (*shippingdomain.Shipping, error) {
	//TODO implement me
	panic("implement me")
}

func (s *shippingUseCase) GetByOrderID(ctx context.Context, orderID primitive.ObjectID) (*shippingdomain.Shipping, error) {
	//TODO implement me
	panic("implement me")
}

func (s *shippingUseCase) UpdateOne(ctx context.Context, shipping shippingdomain.Shipping) error {
	//TODO implement me
	panic("implement me")
}

func (s *shippingUseCase) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	//TODO implement me
	panic("implement me")
}

func (s *shippingUseCase) List(ctx context.Context) ([]shippingdomain.Shipping, error) {
	//TODO implement me
	panic("implement me")
}

func (s *shippingUseCase) GetByStatus(ctx context.Context, status string) ([]shippingdomain.Shipping, error) {
	//TODO implement me
	panic("implement me")
}

func (s *shippingUseCase) UpdateDeliveryStatus(ctx context.Context, id primitive.ObjectID, status string, actualDelivery *time.Time) error {
	//TODO implement me
	panic("implement me")
}
