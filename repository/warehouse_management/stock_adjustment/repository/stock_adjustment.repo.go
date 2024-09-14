package stock_adjustment_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"shop_erp_mono/domain/warehouse_management/stock_adjustment"
	"shop_erp_mono/repository"
	"time"
)

type stockAdjustmentRepository struct {
	database                  *mongo.Database
	stockAdjustmentCollection string
}

func NewStockAdjustmentRepository(database *mongo.Database, stockAdjustmentCollection string) stock_adjustment_domain.IStockAdjustmentRepository {
	return &stockAdjustmentRepository{database: database, stockAdjustmentCollection: stockAdjustmentCollection}
}

func (s *stockAdjustmentRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*stock_adjustment_domain.StockAdjustment, error) {
	//TODO implement me
	panic("implement me")
}

func (s *stockAdjustmentRepository) CreateOne(ctx context.Context, adjustment *stock_adjustment_domain.StockAdjustment) error {
	//TODO implement me
	panic("implement me")
}

func (s *stockAdjustmentRepository) UpdateOne(ctx context.Context, adjustment *stock_adjustment_domain.StockAdjustment) error {
	//TODO implement me
	panic("implement me")
}

func (s *stockAdjustmentRepository) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	//TODO implement me
	panic("implement me")
}

func (s *stockAdjustmentRepository) GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]stock_adjustment_domain.StockAdjustment, error) {
	//TODO implement me
	panic("implement me")
}

func (s *stockAdjustmentRepository) GetByProductID(ctx context.Context, productID primitive.ObjectID) ([]stock_adjustment_domain.StockAdjustment, error) {
	//TODO implement me
	panic("implement me")
}

func (s *stockAdjustmentRepository) GetByWarehouseID(ctx context.Context, warehouseID primitive.ObjectID) ([]stock_adjustment_domain.StockAdjustment, error) {
	//TODO implement me
	panic("implement me")
}

func (s *stockAdjustmentRepository) GetByAdjustmentDateRange(ctx context.Context, startDate, endDate time.Time) ([]stock_adjustment_domain.StockAdjustment, error) {
	//TODO implement me
	panic("implement me")
}
