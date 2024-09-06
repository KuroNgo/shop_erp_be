package product_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IProductRepository interface {
	CreateProduct(ctx context.Context, product Product) (*Product, error)
	UpdateProduct(ctx context.Context, id primitive.ObjectID, product Product) (*Product, error)
	GetProductByID(ctx context.Context, id primitive.ObjectID) (*ProductResponse, error)
	GetProductByName(ctx context.Context, productName string) ([]ProductResponse, error)
	GetAllProducts(ctx context.Context) ([]ProductResponse, error)
	DeleteProduct(ctx context.Context, id primitive.ObjectID) error
}

type IProductUseCase interface {
	CreateProduct(ctx context.Context, input *Input) (*ProductResponse, error)
	UpdateProduct(ctx context.Context, id primitive.ObjectID, input *Input) (*ProductResponse, error)
	GetProductByID(ctx context.Context, id primitive.ObjectID) (*ProductResponse, error)
	GetProductByName(ctx context.Context, productName string) ([]ProductResponse, error)
	GetAllProducts(ctx context.Context) ([]ProductResponse, error)
	DeleteProduct(ctx context.Context, id primitive.ObjectID) error
}
