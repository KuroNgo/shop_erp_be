package stockmovement_usecase

import (
	"context"
	stockmovementdomain "shop_erp_mono/domain/warehouse_management/stockmovement"
	"shop_erp_mono/repository"
	"time"
)

type stockMovementUseCase struct {
	contextTimeout          time.Duration
	stockMovementRepository stockmovementdomain.IStockMovementRepository
}

func NewStockMovementUseCase(contextTimeout time.Duration, stockMovementRepository stockmovementdomain.IStockMovementRepository) stockmovementdomain.IStockMovementUseCase {
	return &stockMovementUseCase{contextTimeout: contextTimeout, stockMovementRepository: stockMovementRepository}
}

func (s *stockMovementUseCase) GetByID(ctx context.Context, id string) (*stockmovementdomain.StockMovement, error) {
	//TODO implement me
	panic("implement me")
}

func (s *stockMovementUseCase) CreateOne(ctx context.Context, input *stockmovementdomain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (s *stockMovementUseCase) UpdateOne(ctx context.Context, id string, input *stockmovementdomain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (s *stockMovementUseCase) DeleteOne(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (s *stockMovementUseCase) GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]stockmovementdomain.StockMovement, error) {
	//TODO implement me
	panic("implement me")
}

func (s *stockMovementUseCase) GetByProductID(ctx context.Context, productID string) ([]stockmovementdomain.StockMovement, error) {
	//TODO implement me
	panic("implement me")
}

func (s *stockMovementUseCase) GetByWarehouseID(ctx context.Context, warehouseID string) ([]stockmovementdomain.StockMovement, error) {
	//TODO implement me
	panic("implement me")
}

func (s *stockMovementUseCase) GetByUserID(ctx context.Context, userID string) ([]stockmovementdomain.StockMovement, error) {
	//TODO implement me
	panic("implement me")
}

func (s *stockMovementUseCase) GetByMovementDateRange(ctx context.Context, startDate, endDate time.Time) ([]stockmovementdomain.StockMovement, error) {
	//TODO implement me
	panic("implement me")
}
