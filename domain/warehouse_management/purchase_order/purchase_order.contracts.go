package purchase_order_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"shop_erp_mono/repository"
)

type IPurchaseOrderRepository interface {
	GetByID(ctx context.Context, id primitive.ObjectID) (*PurchaseOrder, error)
	CreateOne(ctx context.Context, order *PurchaseOrder) error
	UpdateOne(ctx context.Context, order *PurchaseOrder) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]PurchaseOrder, error)
	GetBySupplierID(ctx context.Context, supplierID primitive.ObjectID) ([]PurchaseOrder, error)
	UpdateStatus(ctx context.Context, id primitive.ObjectID, status string) error
}

type IPurchaseOrderUseCase interface {
	GetByID(ctx context.Context, id string) (*PurchaseOrderResponse, error)
	CreateOne(ctx context.Context, input *Input) error
	UpdateOne(ctx context.Context, id string, input *Input) error
	DeleteOne(ctx context.Context, id string) error
	GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]PurchaseOrderResponse, error)
	GetBySupplierID(ctx context.Context, supplierID string) ([]PurchaseOrderResponse, error)
	UpdateStatus(ctx context.Context, id string, status string) error
}
