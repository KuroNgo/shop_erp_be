package stock_adjustment

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"shop_erp_mono/repository"
	"time"
)

type IStockAdjustmentRepository interface {
	GetByID(ctx context.Context, id primitive.ObjectID) (*StockAdjustment, error)
	Create(ctx context.Context, adjustment *StockAdjustment) error
	Update(ctx context.Context, id primitive.ObjectID, adjustment *StockAdjustment) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]StockAdjustment, error)
	GetByProductID(ctx context.Context, productID primitive.ObjectID) ([]StockAdjustment, error)
	GetByWarehouseID(ctx context.Context, warehouseID primitive.ObjectID) ([]StockAdjustment, error)
	GetByAdjustmentDateRange(ctx context.Context, startDate, endDate time.Time) ([]StockAdjustment, error)
}

type IStockAdjustmentUseCase interface {
	GetByID(ctx context.Context, id primitive.ObjectID) (*StockAdjustment, error)
	Create(ctx context.Context, adjustment *StockAdjustment) error
	Update(ctx context.Context, id primitive.ObjectID, adjustment *StockAdjustment) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]StockAdjustment, error)
	GetByProductID(ctx context.Context, productID primitive.ObjectID) ([]StockAdjustment, error)
	GetByWarehouseID(ctx context.Context, warehouseID primitive.ObjectID) ([]StockAdjustment, error)
	GetByAdjustmentDateRange(ctx context.Context, startDate, endDate time.Time) ([]StockAdjustment, error)
}
