package stockmovement_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"shop_erp_mono/repository"
	"time"
)

type IStockMovementRepository interface {
	GetByID(ctx context.Context, id primitive.ObjectID) (*StockMovement, error)
	Create(ctx context.Context, movement *StockMovement) error
	Update(ctx context.Context, movement *StockMovement) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]StockMovement, error)
	GetByProductID(ctx context.Context, productID primitive.ObjectID) ([]StockMovement, error)
	GetByWarehouseID(ctx context.Context, warehouseID primitive.ObjectID) ([]StockMovement, error)
	GetByUserID(ctx context.Context, userID primitive.ObjectID) ([]StockMovement, error)
	GetByMovementDateRange(ctx context.Context, startDate, endDate time.Time) ([]StockMovement, error)
}

type IStockMovementUseCase interface {
	GetByID(ctx context.Context, id primitive.ObjectID) (*StockMovement, error)
	Create(ctx context.Context, movement *StockMovement) error
	Update(ctx context.Context, id primitive.ObjectID, movement *StockMovement) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]StockMovement, error)
	GetByProductID(ctx context.Context, productID primitive.ObjectID) ([]StockMovement, error)
	GetByWarehouseID(ctx context.Context, warehouseID primitive.ObjectID) ([]StockMovement, error)
	GetByUserID(ctx context.Context, userID primitive.ObjectID) ([]StockMovement, error)
	GetByMovementDateRange(ctx context.Context, startDate, endDate time.Time) ([]StockMovement, error)
}
