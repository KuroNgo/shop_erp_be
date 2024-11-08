package supplier_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"shop_erp_mono/internal/repository"
)

type ISupplierRepository interface {
	CreateOne(ctx context.Context, supplier Supplier) error
	GetByID(ctx context.Context, id primitive.ObjectID) (*Supplier, error)
	GetByName(ctx context.Context, name string) (*Supplier, error)
	GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]Supplier, error)
	GetAll(ctx context.Context) ([]Supplier, error)
	UpdateOne(ctx context.Context, supplier *Supplier) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
}

type ISupplierUseCase interface {
	CreateOne(ctx context.Context, input *Input) error
	GetByID(ctx context.Context, id string) (*SupplierResponse, error)
	GetByName(ctx context.Context, name string) (*SupplierResponse, error)
	GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]SupplierResponse, error)
	GetAll(ctx context.Context) ([]SupplierResponse, error)
	UpdateOne(ctx context.Context, id string, input *Input) error
	DeleteOne(ctx context.Context, id string) error
}
