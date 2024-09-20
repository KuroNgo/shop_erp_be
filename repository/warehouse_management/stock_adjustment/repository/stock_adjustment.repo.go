package stock_adjustment_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
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
	stockAdjustmentCollection := s.database.Collection(s.stockAdjustmentCollection)

	filter := bson.M{"_id": id}
	var stock stock_adjustment_domain.StockAdjustment
	if err := stockAdjustmentCollection.FindOne(ctx, filter).Decode(&stock); err != nil {
		return nil, err
	}

	return &stock, nil
}

func (s *stockAdjustmentRepository) CreateOne(ctx context.Context, adjustment *stock_adjustment_domain.StockAdjustment) error {
	stockAdjustmentCollection := s.database.Collection(s.stockAdjustmentCollection)

	_, err := stockAdjustmentCollection.InsertOne(ctx, adjustment)
	if err != nil {
		return err
	}

	return nil
}

func (s *stockAdjustmentRepository) UpdateOne(ctx context.Context, adjustment *stock_adjustment_domain.StockAdjustment) error {
	stockAdjustmentCollection := s.database.Collection(s.stockAdjustmentCollection)

	filter := bson.M{"_id": adjustment.ID}
	update := bson.M{"$set": bson.M{
		"product_id":      adjustment.ProductID,
		"warehouse_id":    adjustment.WarehouseID,
		"adjustment_type": adjustment.AdjustmentType,
		"adjustment_date": adjustment.AdjustmentDate,
		"quantity":        adjustment.Quantity,
		"reason":          adjustment.Reason,
		"updated_at":      time.Now(),
	}}
	_, err := stockAdjustmentCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (s *stockAdjustmentRepository) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	stockAdjustmentCollection := s.database.Collection(s.stockAdjustmentCollection)

	filter := bson.M{"_id": id}
	_, err := stockAdjustmentCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (s *stockAdjustmentRepository) GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]stock_adjustment_domain.StockAdjustment, error) {
	stockAdjustmentCollection := s.database.Collection(s.stockAdjustmentCollection)

	filter := bson.M{}
	cursor, err := repository.Paginate(ctx, stockAdjustmentCollection, filter, pagination)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var stocks []stock_adjustment_domain.StockAdjustment
	stocks = make([]stock_adjustment_domain.StockAdjustment, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var stock stock_adjustment_domain.StockAdjustment
		if err = cursor.Decode(&stock); err != nil {
			return nil, err
		}

		stocks = append(stocks, stock)
	}

	return stocks, nil
}

func (s *stockAdjustmentRepository) GetByProductID(ctx context.Context, productID primitive.ObjectID) ([]stock_adjustment_domain.StockAdjustment, error) {
	stockAdjustmentCollection := s.database.Collection(s.stockAdjustmentCollection)

	filter := bson.M{"product_id": productID}
	cursor, err := stockAdjustmentCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var stocks []stock_adjustment_domain.StockAdjustment
	stocks = make([]stock_adjustment_domain.StockAdjustment, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var stock stock_adjustment_domain.StockAdjustment
		if err = cursor.Decode(&stock); err != nil {
			return nil, err
		}

		stocks = append(stocks, stock)
	}

	return stocks, nil
}

func (s *stockAdjustmentRepository) GetByWarehouseID(ctx context.Context, warehouseID primitive.ObjectID) ([]stock_adjustment_domain.StockAdjustment, error) {
	stockAdjustmentCollection := s.database.Collection(s.stockAdjustmentCollection)

	filter := bson.M{"warehouse_id": warehouseID}
	cursor, err := stockAdjustmentCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var stocks []stock_adjustment_domain.StockAdjustment
	stocks = make([]stock_adjustment_domain.StockAdjustment, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var stock stock_adjustment_domain.StockAdjustment
		if err = cursor.Decode(&stock); err != nil {
			return nil, err
		}

		stocks = append(stocks, stock)
	}

	return stocks, nil
}

func (s *stockAdjustmentRepository) GetByAdjustmentDateRange(ctx context.Context, startDate, endDate time.Time) ([]stock_adjustment_domain.StockAdjustment, error) {
	stockAdjustmentCollection := s.database.Collection(s.stockAdjustmentCollection)

	filter := bson.M{
		"adjustment_date": bson.M{
			"$gte": startDate,
			"$lte": endDate,
		},
	}
	cursor, err := stockAdjustmentCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var stocks []stock_adjustment_domain.StockAdjustment
	stocks = make([]stock_adjustment_domain.StockAdjustment, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var stock stock_adjustment_domain.StockAdjustment
		if err = cursor.Decode(&stock); err != nil {
			return nil, err
		}

		stocks = append(stocks, stock)
	}

	return stocks, nil
}
