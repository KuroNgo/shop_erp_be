package sale_orders_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ISalesOrderRepository interface {
	GetByID(ctx context.Context, id primitive.ObjectID) (*SalesOrder, error)
	GetByCustomerID(ctx context.Context, customerID primitive.ObjectID) ([]SalesOrder, error)
	GetByStatus(ctx context.Context, status string) ([]SalesOrder, error)
	GetAll(ctx context.Context) ([]SalesOrder, error)
	CreateOne(ctx context.Context, order SalesOrder) error
	UpdateOne(ctx context.Context, order SalesOrder) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
}

type ISalesOrderUseCase interface {
	GetByID(ctx context.Context, id string) (*SalesOrderResponse, error)
	GetByCustomerID(ctx context.Context, customerID string) ([]SalesOrderResponse, error)
	GetByStatus(ctx context.Context, status string) ([]SalesOrderResponse, error)
	GetAll(ctx context.Context) ([]SalesOrderResponse, error)
	CreateOne(ctx context.Context, input *Input) error
	UpdateOne(ctx context.Context, id string, input *Input) error
	DeleteOne(ctx context.Context, id string) error
}
