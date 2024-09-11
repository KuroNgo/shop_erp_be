package customer_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ICustomerRepository interface {
	CreateOne(ctx context.Context, customer *Customer) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	UpdateOne(ctx context.Context, id primitive.ObjectID, customer *Customer) error
	GetOneByID(ctx context.Context, id primitive.ObjectID) (*Customer, error)
	GetAll(ctx context.Context) ([]Customer, error)
}

type ICustomerUseCase interface {
	CreateOne(ctx context.Context, customer *Customer) error
	DeleteOne(ctx context.Context, id string) error
	UpdateOne(ctx context.Context, id string, customer *Customer) error
	GetOneByID(ctx context.Context, id string) (CustomerResponse, error)
	GetAll(ctx context.Context) ([]CustomerResponse, error)
}
