package stock_adjustment_usecase

import (
	"context"
	stock_adjustment_domain "shop_erp_mono/domain/warehouse_management/stock_adjustment"
	"shop_erp_mono/repository"
	"time"
)

type stockAdjustmentUseCase struct {
	contextTimeout            time.Duration
	stockAdjustmentRepository stock_adjustment_domain.IStockAdjustmentRepository
}

func (s *stockAdjustmentUseCase) GetByID(ctx context.Context, id string) (*stock_adjustment_domain.StockAdjustment, error) {
	//TODO implement me
	panic("implement me")
}

func (s *stockAdjustmentUseCase) CreateOne(ctx context.Context, input *stock_adjustment_domain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (s *stockAdjustmentUseCase) UpdateOne(ctx context.Context, id string, input *stock_adjustment_domain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (s *stockAdjustmentUseCase) DeleteOne(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (s *stockAdjustmentUseCase) GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]stock_adjustment_domain.StockAdjustment, error) {
	//TODO implement me
	panic("implement me")
}

func (s *stockAdjustmentUseCase) GetByProductID(ctx context.Context, productID string) ([]stock_adjustment_domain.StockAdjustment, error) {
	//TODO implement me
	panic("implement me")
}

func (s *stockAdjustmentUseCase) GetByWarehouseID(ctx context.Context, warehouseID string) ([]stock_adjustment_domain.StockAdjustment, error) {
	//TODO implement me
	panic("implement me")
}

func (s *stockAdjustmentUseCase) GetByAdjustmentDateRange(ctx context.Context, startDate, endDate time.Time) ([]stock_adjustment_domain.StockAdjustment, error) {
	//TODO implement me
	panic("implement me")
}

func NewStockAdjustmentUseCase(contextTimeout time.Duration, stockAdjustmentRepository stock_adjustment_domain.IStockAdjustmentRepository) stock_adjustment_domain.IStockAdjustmentUseCase {
	return &stockAdjustmentUseCase{contextTimeout: contextTimeout, stockAdjustmentRepository: stockAdjustmentRepository}
}
