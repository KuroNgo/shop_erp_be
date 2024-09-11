package product_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IProductRepository interface {
	CreateProduct(ctx context.Context, product Product) error
	UpdateProduct(ctx context.Context, id primitive.ObjectID, product Product) error
	GetProductByID(ctx context.Context, id primitive.ObjectID) (*Product, error)
	GetProductByName(ctx context.Context, productName string) (*Product, error)
	GetAllProducts(ctx context.Context) ([]Product, error)
	DeleteProduct(ctx context.Context, id primitive.ObjectID) error
	CountCategory(ctx context.Context, categoryID primitive.ObjectID) (int64, error)
}

type IProductUseCase interface {
	CreateProduct(ctx context.Context, input *Input) error
	UpdateProduct(ctx context.Context, id string, input *Input) error
	GetProductByID(ctx context.Context, id string) (*ProductResponse, error)
	GetProductByName(ctx context.Context, productName string) (*ProductResponse, error)
	GetAllProducts(ctx context.Context) ([]ProductResponse, error)
	DeleteProduct(ctx context.Context, id string) error
}
