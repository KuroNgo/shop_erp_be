package stock_adjustment_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"shop_erp_mono/repository"
	"time"
)

type IStockAdjustmentRepository interface {
	GetByID(ctx context.Context, id primitive.ObjectID) (*StockAdjustment, error)
	CreateOne(ctx context.Context, adjustment *StockAdjustment) error
	UpdateOne(ctx context.Context, adjustment *StockAdjustment) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]StockAdjustment, error)
	GetByProductID(ctx context.Context, productID primitive.ObjectID) ([]StockAdjustment, error)
	GetByWarehouseID(ctx context.Context, warehouseID primitive.ObjectID) ([]StockAdjustment, error)
	GetByAdjustmentDateRange(ctx context.Context, startDate, endDate time.Time) ([]StockAdjustment, error)
}

type IStockAdjustmentUseCase interface {
	GetByID(ctx context.Context, id string) (*StockAdjustment, error)
	CreateOne(ctx context.Context, input *Input) error
	UpdateOne(ctx context.Context, id string, input *Input) error
	DeleteOne(ctx context.Context, id string) error
	GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]StockAdjustment, error)
	GetByProductID(ctx context.Context, productID string) ([]StockAdjustment, error)
	GetByWarehouseID(ctx context.Context, warehouseID string) ([]StockAdjustment, error)
	GetByAdjustmentDateRange(ctx context.Context, startDate, endDate time.Time) ([]StockAdjustment, error)
}
