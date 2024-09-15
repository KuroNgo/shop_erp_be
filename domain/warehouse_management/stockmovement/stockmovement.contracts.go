package stockmovement_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"shop_erp_mono/repository"
	"time"
)

type IStockMovementRepository interface {
	GetByID(ctx context.Context, id primitive.ObjectID) (*StockMovement, error)
	CreateOne(ctx context.Context, movement *StockMovement) error
	UpdateOne(ctx context.Context, movement *StockMovement) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]StockMovement, error)
	GetByProductID(ctx context.Context, productID primitive.ObjectID) ([]StockMovement, error)
	GetByWarehouseID(ctx context.Context, warehouseID primitive.ObjectID) ([]StockMovement, error)
	GetByUserID(ctx context.Context, userID primitive.ObjectID) ([]StockMovement, error)
	GetByMovementDateRange(ctx context.Context, startDate, endDate time.Time) ([]StockMovement, error)
}

type IStockMovementUseCase interface {
	GetByID(ctx context.Context, id string) (*StockMovementResponse, error)
	CreateOne(ctx context.Context, input *Input) error
	UpdateOne(ctx context.Context, id string, input *Input) error
	DeleteOne(ctx context.Context, id string) error
	GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]StockMovementResponse, error)
	GetByProductID(ctx context.Context, productID string) ([]StockMovementResponse, error)
	GetByWarehouseID(ctx context.Context, warehouseID string) ([]StockMovementResponse, error)
	GetByUserID(ctx context.Context, userID string) ([]StockMovementResponse, error)
	GetByMovementDateRange(ctx context.Context, startDate, endDate time.Time) ([]StockMovementResponse, error)
}
