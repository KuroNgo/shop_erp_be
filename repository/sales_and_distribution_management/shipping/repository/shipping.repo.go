package shipping_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	shippingdomain "shop_erp_mono/domain/sales_and_distribution_management/shipping"
	"time"
)

type shippingRepository struct {
	database           *mongo.Database
	shippingCollection string
}

func NewShippingRepository(database *mongo.Database, shippingCollection string) shippingdomain.IShippingRepository {
	return &shippingRepository{database: database, shippingCollection: shippingCollection}
}

func (s *shippingRepository) CreateOne(ctx context.Context, shipping shippingdomain.Input) (primitive.ObjectID, error) {
	//TODO implement me
	panic("implement me")
}

func (s *shippingRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*shippingdomain.Shipping, error) {
	//TODO implement me
	panic("implement me")
}

func (s *shippingRepository) GetByOrderID(ctx context.Context, orderID primitive.ObjectID) (*shippingdomain.Shipping, error) {
	//TODO implement me
	panic("implement me")
}

func (s *shippingRepository) UpdateOne(ctx context.Context, id primitive.ObjectID, updatedShipping shippingdomain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (s *shippingRepository) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	//TODO implement me
	panic("implement me")
}

func (s *shippingRepository) List(ctx context.Context) ([]shippingdomain.Shipping, error) {
	//TODO implement me
	panic("implement me")
}

func (s *shippingRepository) GetByStatus(ctx context.Context, status string) ([]shippingdomain.Shipping, error) {
	//TODO implement me
	panic("implement me")
}

func (s *shippingRepository) UpdateDeliveryStatus(ctx context.Context, id primitive.ObjectID, status string, actualDelivery *time.Time) error {
	//TODO implement me
	panic("implement me")
}
