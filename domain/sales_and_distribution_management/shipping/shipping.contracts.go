package shipping_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type IShippingRepository interface {
	CreateOne(ctx context.Context, shipping Input) (primitive.ObjectID, error)
	GetByID(ctx context.Context, id primitive.ObjectID) (*Shipping, error)
	GetByOrderID(ctx context.Context, orderID primitive.ObjectID) (*Shipping, error)
	UpdateOne(ctx context.Context, id primitive.ObjectID, updatedShipping Input) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	List(ctx context.Context) ([]Shipping, error)
	GetByStatus(ctx context.Context, status string) ([]Shipping, error)
	UpdateDeliveryStatus(ctx context.Context, id primitive.ObjectID, status string, actualDelivery *time.Time) error
}

type IShippingUseCase interface {
	CreateOne(ctx context.Context, input *Input) error
	GetByID(ctx context.Context, id string) (*ShippingResponse, error)
	GetByOrderID(ctx context.Context, orderID string) (*ShippingResponse, error)
	UpdateOne(ctx context.Context, id string, updatedShipping Input) error
	DeleteOne(ctx context.Context, id string) error
	List(ctx context.Context) ([]ShippingResponse, error)
	GetByStatus(ctx context.Context, status string) ([]ShippingResponse, error)
	UpdateDeliveryStatus(ctx context.Context, id string, status string, actualDelivery *time.Time) error
}
