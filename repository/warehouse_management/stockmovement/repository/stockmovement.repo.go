package stockmovement_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	stockmovementdomain "shop_erp_mono/domain/warehouse_management/stockmovement"
	"shop_erp_mono/repository"
	"time"
)

type stockMovementRepository struct {
	database                *mongo.Database
	stockMovementCollection string
}

func NewStockMovementRepository(database *mongo.Database, stockMovementCollection string) stockmovementdomain.IStockMovementRepository {
	return &stockMovementRepository{database: database, stockMovementCollection: stockMovementCollection}
}

func (s *stockMovementRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*stockmovementdomain.StockMovement, error) {
	stockMovementCollection := s.database.Collection(s.stockMovementCollection)

	filter := bson.M{"_id": id}
	var stockMovement stockmovementdomain.StockMovement
	if err := stockMovementCollection.FindOne(ctx, filter).Decode(&stockMovement); err != nil {
		return nil, err
	}

	return &stockMovement, nil
}

func (s *stockMovementRepository) Create(ctx context.Context, movement *stockmovementdomain.StockMovement) error {
	stockMovementCollection := s.database.Collection(s.stockMovementCollection)

	_, err := stockMovementCollection.InsertOne(ctx, movement)
	if err != nil {
		return err
	}

	return nil
}

func (s *stockMovementRepository) Update(ctx context.Context, movement *stockmovementdomain.StockMovement) error {
	stockMovementCollection := s.database.Collection(s.stockMovementCollection)

	filter := bson.M{"_id": movement.ID}
	update := bson.M{"$set": bson.M{
		"product_id":    movement.ProductID,
		"warehouse_id":  movement.WarehouseID,
		"user_id":       movement.UserID,
		"movement_type": movement.MovementType,
		"quantity":      movement.Quantity,
		"movement_date": movement.MovementDate,
		"reference":     movement.Reference,
		"updated_at":    movement.UpdatedAt,
	}}

	_, err := stockMovementCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (s *stockMovementRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	stockMovementCollection := s.database.Collection(s.stockMovementCollection)

	filter := bson.M{"_id": id}
	_, err := stockMovementCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (s *stockMovementRepository) GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]stockmovementdomain.StockMovement, error) {
	stockMovementCollection := s.database.Collection(s.stockMovementCollection)

	filter := bson.M{}
	cursor, err := repository.Paginate(ctx, stockMovementCollection, filter, pagination)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	var stockMovements []stockmovementdomain.StockMovement
	stockMovements = make([]stockmovementdomain.StockMovement, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var stockMovement stockmovementdomain.StockMovement
		if err = cursor.Decode(&stockMovement); err != nil {
			return nil, err
		}

		stockMovements = append(stockMovements, stockMovement)
	}

	return stockMovements, nil
}

func (s *stockMovementRepository) GetByProductID(ctx context.Context, productID primitive.ObjectID) ([]stockmovementdomain.StockMovement, error) {
	stockMovementCollection := s.database.Collection(s.stockMovementCollection)

	filter := bson.M{"product_id": productID}
	cursor, err := stockMovementCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	var stockMovements []stockmovementdomain.StockMovement
	stockMovements = make([]stockmovementdomain.StockMovement, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var stockMovement stockmovementdomain.StockMovement
		if err = cursor.Decode(&stockMovement); err != nil {
			return nil, err
		}

		stockMovements = append(stockMovements, stockMovement)
	}

	return stockMovements, nil
}

func (s *stockMovementRepository) GetByWarehouseID(ctx context.Context, warehouseID primitive.ObjectID) ([]stockmovementdomain.StockMovement, error) {
	stockMovementCollection := s.database.Collection(s.stockMovementCollection)

	filter := bson.M{"warehouse_id": warehouseID}
	cursor, err := stockMovementCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	var stockMovements []stockmovementdomain.StockMovement
	stockMovements = make([]stockmovementdomain.StockMovement, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var stockMovement stockmovementdomain.StockMovement
		if err = cursor.Decode(&stockMovement); err != nil {
			return nil, err
		}

		stockMovements = append(stockMovements, stockMovement)
	}

	return stockMovements, nil
}

func (s *stockMovementRepository) GetByUserID(ctx context.Context, userID primitive.ObjectID) ([]stockmovementdomain.StockMovement, error) {
	stockMovementCollection := s.database.Collection(s.stockMovementCollection)

	filter := bson.M{"user_id": userID}
	cursor, err := stockMovementCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	var stockMovements []stockmovementdomain.StockMovement
	stockMovements = make([]stockmovementdomain.StockMovement, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var stockMovement stockmovementdomain.StockMovement
		if err = cursor.Decode(&stockMovement); err != nil {
			return nil, err
		}

		stockMovements = append(stockMovements, stockMovement)
	}

	return stockMovements, nil
}

func (s *stockMovementRepository) GetByMovementDateRange(ctx context.Context, startDate, endDate time.Time) ([]stockmovementdomain.StockMovement, error) {
	stockMovementCollection := s.database.Collection(s.stockMovementCollection)

	// Tạo filter theo khoảng thời gian MovementDate
	filter := bson.M{
		"movement_date": bson.M{
			"$gte": startDate,
			"$lte": endDate,
		},
	}
	cursor, err := stockMovementCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	var stockMovements []stockmovementdomain.StockMovement
	stockMovements = make([]stockmovementdomain.StockMovement, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var stockMovement stockmovementdomain.StockMovement
		if err = cursor.Decode(&stockMovement); err != nil {
			return nil, err
		}

		stockMovements = append(stockMovements, stockMovement)
	}

	return stockMovements, nil
}
