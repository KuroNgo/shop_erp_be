package supplier_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ISupplierRepository interface {
	CreateOne(ctx context.Context, supplier Supplier) (primitive.ObjectID, error)
	GetByID(ctx context.Context, id primitive.ObjectID) (*Supplier, error)
	GetAll(ctx context.Context) ([]Supplier, error)
	UpdateOne(ctx context.Context, id primitive.ObjectID, supplier Supplier) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
}

type SupplierUseCase interface {
	CreateSupplier(ctx context.Context, input Input) (*SupplierResponse, error)
	GetSupplierByID(ctx context.Context, id string) (*SupplierResponse, error)
	GetSuppliers(ctx context.Context) ([]SupplierResponse, error)
	UpdateSupplier(ctx context.Context, id string, input Input) error
	DeleteSupplier(ctx context.Context, id string) error
}
