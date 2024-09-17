package sale_orders_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ISalesOrderRepository interface {
	GetByID(ctx context.Context, id primitive.ObjectID) (*SalesOrderResponse, error)
	GetByCustomerID(ctx context.Context, customerID primitive.ObjectID) ([]SalesOrderResponse, error)
	GetByStatus(ctx context.Context, status string) ([]SalesOrderResponse, error)
	List(ctx context.Context, filters map[string]interface{}) ([]SalesOrderResponse, error)
	CreateOne(ctx context.Context, order SalesOrder) error
	UpdateOne(ctx context.Context, id primitive.ObjectID, updatedOrder Input) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
}

type ISalesOrderUseCase interface {
	GetByID(ctx context.Context, id primitive.ObjectID) (*SalesOrderResponse, error)
	GetByCustomerID(ctx context.Context, customerID primitive.ObjectID) ([]SalesOrderResponse, error)
	GetByStatus(ctx context.Context, status string) ([]SalesOrderResponse, error)
	List(ctx context.Context, filters map[string]interface{}) ([]SalesOrderResponse, error)
	CreateOne(ctx context.Context, order SalesOrder) error
	UpdateOne(ctx context.Context, id primitive.ObjectID, updatedOrder Input) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
}
