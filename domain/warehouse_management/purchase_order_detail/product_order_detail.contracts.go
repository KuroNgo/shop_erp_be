package purchase_order_detail_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"shop_erp_mono/repository"
)

type IPurchaseOrderDetailRepository interface {
	GetByID(ctx context.Context, id primitive.ObjectID) (*PurchaseOrderDetail, error)
	GetByPurchaseOrderID(ctx context.Context, purchaseOrderID primitive.ObjectID) ([]PurchaseOrderDetail, error)
	Create(ctx context.Context, detail *PurchaseOrderDetail) error
	Update(ctx context.Context, id primitive.ObjectID, detail *PurchaseOrderDetail) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]PurchaseOrderDetail, error)
	GetAll(ctx context.Context) ([]PurchaseOrderDetail, error)
}

type IPurchaseOrderDetailUseCase interface {
	GetByID(ctx context.Context, id string) (*PurchaseOrderDetail, error)
	GetByPurchaseOrderID(ctx context.Context, purchaseOrderID string) ([]PurchaseOrderDetail, error)
	Create(ctx context.Context, input *Input) error
	Update(ctx context.Context, id string, input *Input) error
	Delete(ctx context.Context, id string) error
	GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]PurchaseOrderDetail, error)
	GetAll(ctx context.Context) ([]PurchaseOrderDetail, error)
}
