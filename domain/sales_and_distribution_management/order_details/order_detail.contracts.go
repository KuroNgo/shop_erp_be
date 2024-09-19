package order_details_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IOrderDetailRepository interface {
	CreateOne(ctx context.Context, detail OrderDetail) error
	GetByID(ctx context.Context, id primitive.ObjectID) (*OrderDetail, error)
	GetByOrderID(ctx context.Context, orderID primitive.ObjectID) ([]OrderDetail, error)
	GetByProductID(ctx context.Context, productID primitive.ObjectID) ([]OrderDetail, error)
	UpdateOne(ctx context.Context, detail OrderDetail) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	GetAll(ctx context.Context) ([]OrderDetail, error)
}

type IOrderDetailUseCase interface {
	CreateOne(ctx context.Context, input *Input) error
	GetByID(ctx context.Context, id string) (*OrderDetailResponse, error)
	GetByOrderID(ctx context.Context, orderID string) ([]OrderDetailResponse, error)
	GetByProductID(ctx context.Context, productID string) ([]OrderDetailResponse, error)
	UpdateOne(ctx context.Context, id string, input *Input) error
	DeleteOne(ctx context.Context, id string) error
	GetAll(ctx context.Context) ([]OrderDetailResponse, error)
}
