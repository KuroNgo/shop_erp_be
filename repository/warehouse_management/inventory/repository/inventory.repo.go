package inventory_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	inventorydomain "shop_erp_mono/domain/warehouse_management/inventory"
	"time"
)

type inventoryRepository struct {
	database            *mongo.Database
	inventoryCollection string
}

func (i *inventoryRepository) CreateOne(ctx context.Context, inventory inventorydomain.Inventory) error {
	inventoryCollection := i.database.Collection(i.inventoryCollection)

	_, err := inventoryCollection.InsertOne(ctx, inventory)
	if err != nil {
		return err
	}

	return nil
}

func (i *inventoryRepository) UpdateOne(ctx context.Context, inventory inventorydomain.Inventory) error {
	inventoryCollection := i.database.Collection(i.inventoryCollection)

	filter := bson.M{"_id": inventory.ID}
	update := bson.M{"$set": bson.M{
		"product_id":   inventory.ProductID,
		"warehouse_id": inventory.WarehouseID,
		"quantity":     inventory.Quantity,
		"updated_at":   time.Now(),
	}}
	_, err := inventoryCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (i *inventoryRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*inventorydomain.Inventory, error) {
	inventoryCollection := i.database.Collection(i.inventoryCollection)

	filter := bson.M{"_id": id}
	var inventory inventorydomain.Inventory
	if err := inventoryCollection.FindOne(ctx, filter).Decode(&inventory); err != nil {
		return nil, err
	}

	return &inventory, nil
}

func (i *inventoryRepository) GetByProduct(ctx context.Context, productID primitive.ObjectID) ([]inventorydomain.Inventory, error) {
	inventoryCollection := i.database.Collection(i.inventoryCollection)

	filter := bson.M{"product_id": productID}
	cursor, err := inventoryCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	var inventories []inventorydomain.Inventory
	inventories = make([]inventorydomain.Inventory, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var inventory inventorydomain.Inventory
		if err = cursor.Decode(&inventory); err != nil {
			return nil, err
		}

		inventories = append(inventories, inventory)
	}

	return inventories, nil
}

func (i *inventoryRepository) GetByWarehouse(ctx context.Context, warehouseID primitive.ObjectID) ([]inventorydomain.Inventory, error) {
	inventoryCollection := i.database.Collection(i.inventoryCollection)

	filter := bson.M{"warehouse_id": warehouseID}
	cursor, err := inventoryCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	var inventories []inventorydomain.Inventory
	inventories = make([]inventorydomain.Inventory, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var inventory inventorydomain.Inventory
		if err = cursor.Decode(&inventory); err != nil {
			return nil, err
		}

		inventories = append(inventories, inventory)
	}

	return inventories, nil
}

func (i *inventoryRepository) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	inventoryCollection := i.database.Collection(i.inventoryCollection)

	filter := bson.M{"_id": id}
	_, err := inventoryCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (i *inventoryRepository) AdjustQuantity(ctx context.Context, id primitive.ObjectID, adjustment int) (*inventorydomain.Inventory, error) {
	//TODO implement me
	panic("implement me")
}

func (i *inventoryRepository) CheckAvailability(ctx context.Context, productID primitive.ObjectID, warehouseID primitive.ObjectID, requiredQuantity int) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (i *inventoryRepository) GetAll(ctx context.Context) ([]inventorydomain.Inventory, error) {
	inventoryCollection := i.database.Collection(i.inventoryCollection)

	filter := bson.M{}
	cursor, err := inventoryCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	var inventories []inventorydomain.Inventory
	inventories = make([]inventorydomain.Inventory, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var inventory inventorydomain.Inventory
		if err = cursor.Decode(&inventory); err != nil {
			return nil, err
		}

		inventories = append(inventories, inventory)
	}

	return inventories, nil
}

func NewInventoryRepository(database *mongo.Database, inventoryCollection string) inventorydomain.InventoryRepository {
	return &inventoryRepository{database: database, inventoryCollection: inventoryCollection}
}
