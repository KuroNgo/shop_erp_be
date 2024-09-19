package purchase_order_detail_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"shop_erp_mono/repository"
)

type IPurchaseOrderDetailRepository interface {
	GetByID(ctx context.Context, id primitive.ObjectID) (*PurchaseOrderDetail, error)
	GetByPurchaseOrderID(ctx context.Context, purchaseOrderID primitive.ObjectID) ([]PurchaseOrderDetail, error)
	CreateOne(ctx context.Context, detail *PurchaseOrderDetail) error
	UpdateOne(ctx context.Context, detail *PurchaseOrderDetail) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]PurchaseOrderDetail, error)
	GetAll(ctx context.Context) ([]PurchaseOrderDetail, error)
}

type IPurchaseOrderDetailUseCase interface {
	GetByID(ctx context.Context, id string) (*PurchaseOrderDetailResponse, error)
	GetByPurchaseOrderID(ctx context.Context, purchaseOrderID string) ([]PurchaseOrderDetailResponse, error)
	CreateOne(ctx context.Context, input *Input) error
	UpdateOne(ctx context.Context, id string, input *Input) error
	DeleteOne(ctx context.Context, id string) error
	GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]PurchaseOrderDetailResponse, error)
	GetAll(ctx context.Context) ([]PurchaseOrderDetailResponse, error)
}
