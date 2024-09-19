package product_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IProductRepository interface {
	CreateOne(ctx context.Context, product Product) error
	UpdateOne(ctx context.Context, id primitive.ObjectID, product Product) error
	GetByID(ctx context.Context, id primitive.ObjectID) (*Product, error)
	GetByName(ctx context.Context, productName string) (*Product, error)
	GetAll(ctx context.Context) ([]Product, error)
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	CountCategory(ctx context.Context, categoryID primitive.ObjectID) (int64, error)
}

type IProductUseCase interface {
	CreateOne(ctx context.Context, input *Input) error
	UpdateOne(ctx context.Context, id string, input *Input) error
	GetByID(ctx context.Context, id string) (*ProductResponse, error)
	GetByName(ctx context.Context, productName string) (*ProductResponse, error)
	GetAll(ctx context.Context) ([]ProductResponse, error)
	DeleteOne(ctx context.Context, id string) error
}
