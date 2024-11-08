package warehouse_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	warehousedomain "shop_erp_mono/internal/domain/warehouse_management/warehouse"
	"time"
)

type warehouseRepository struct {
	database            *mongo.Database
	warehouseCollection string
}

func NewWarehouseRepository(database *mongo.Database, warehouseCollection string) warehousedomain.IWarehouseRepository {
	return &warehouseRepository{database: database, warehouseCollection: warehouseCollection}
}

func (w *warehouseRepository) CreateOne(ctx context.Context, warehouse warehousedomain.Warehouse) error {
	warehouseCollection := w.database.Collection(w.warehouseCollection)

	_, err := warehouseCollection.InsertOne(ctx, warehouse)
	if err != nil {
		return err
	}

	return nil
}

func (w *warehouseRepository) UpdateOne(ctx context.Context, warehouse warehousedomain.Warehouse) error {
	warehouseCollection := w.database.Collection(w.warehouseCollection)

	filter := bson.M{"_id": warehouse.ID}
	update := bson.M{"$set": bson.M{
		"warehouse_name": warehouse.WarehouseName,
		"location":       warehouse.Location,
		"capacity":       warehouse.Capacity,
		"updated_at":     time.Now(),
	}}
	_, err := warehouseCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (w *warehouseRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*warehousedomain.Warehouse, error) {
	warehouseCollection := w.database.Collection(w.warehouseCollection)

	filter := bson.M{"_id": id}
	var warehouse warehousedomain.Warehouse
	if err := warehouseCollection.FindOne(ctx, filter).Decode(&warehouse); err != nil {
		return nil, err
	}

	return &warehouse, nil
}

func (w *warehouseRepository) GetByName(ctx context.Context, name string) (*warehousedomain.Warehouse, error) {
	warehouseCollection := w.database.Collection(w.warehouseCollection)

	filter := bson.M{"name": name}
	var warehouse warehousedomain.Warehouse
	if err := warehouseCollection.FindOne(ctx, filter).Decode(&warehouse); err != nil {
		return nil, err
	}

	return &warehouse, nil
}

func (w *warehouseRepository) GetAll(ctx context.Context) ([]warehousedomain.Warehouse, error) {
	warehouseCollection := w.database.Collection(w.warehouseCollection)

	filter := bson.M{}
	cursor, err := warehouseCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	var warehouses []warehousedomain.Warehouse
	warehouses = make([]warehousedomain.Warehouse, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var warehouse warehousedomain.Warehouse
		if err = cursor.Decode(&warehouse); err != nil {
			return nil, err
		}

		warehouses = append(warehouses, warehouse)
	}

	return warehouses, nil
}

func (w *warehouseRepository) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	warehouseCollection := w.database.Collection(w.warehouseCollection)

	filter := bson.M{"_id": id}
	_, err := warehouseCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
