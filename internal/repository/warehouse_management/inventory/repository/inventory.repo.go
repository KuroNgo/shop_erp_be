package inventory_repository

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	inventorydomain "shop_erp_mono/internal/domain/warehouse_management/inventory"
	"time"
)

type inventoryRepository struct {
	database            *mongo.Database
	inventoryCollection string
}

func NewInventoryRepository(database *mongo.Database, inventoryCollection string) inventorydomain.InventoryRepository {
	return &inventoryRepository{database: database, inventoryCollection: inventoryCollection}
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
		"product_id":       inventory.ProductID,
		"warehouse_id":     inventory.WarehouseID,
		"quantity_warning": inventory.QuantityWarning,
		"quantity":         inventory.Quantity,
		"updated_at":       time.Now(),
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
	inventoryCollection := i.database.Collection(i.inventoryCollection)

	filter := bson.M{"_id": id}
	var inventory inventorydomain.Inventory
	if err := inventoryCollection.FindOne(ctx, filter).Decode(&inventory); err != nil {
		return nil, err
	}

	newQuantity := inventory.Quantity + adjustment // Điều chỉnh số lượng (adjust quantity)
	update := bson.M{"$set": bson.M{"quantity": newQuantity}}
	if _, err := inventoryCollection.UpdateOne(ctx, filter, update); err != nil {
		return nil, err
	}

	inventory.Quantity = newQuantity

	return &inventory, nil
}

func (i *inventoryRepository) CheckAvailability(ctx context.Context, productID primitive.ObjectID, warehouseID primitive.ObjectID, requiredQuantity int) (bool, error) {
	inventoryCollection := i.database.Collection(i.inventoryCollection)

	filter := bson.M{
		"product_id":   productID,
		"warehouse_id": warehouseID,
	}

	var inventory inventorydomain.Inventory
	if err := inventoryCollection.FindOne(ctx, filter).Decode(&inventory); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}
		return false, err
	}

	if inventory.Quantity >= requiredQuantity {
		return true, nil
	}

	return false, nil
}

func (i *inventoryRepository) CheckQuantity(ctx context.Context, productID primitive.ObjectID, quantity int) (bool, error) {
	inventoryCollection := i.database.Collection(i.inventoryCollection)

	filter := bson.M{
		"product_id": productID,
	}

	var inventory inventorydomain.Inventory
	if err := inventoryCollection.FindOne(ctx, filter).Decode(&inventory); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}
		return false, err
	}

	if inventory.Quantity < quantity {
		return false, nil
	}

	return true, nil
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

// WarningOutOfStock warning all products at risk are out of stock
func (i *inventoryRepository) WarningOutOfStock(ctx context.Context) ([]inventorydomain.Inventory, error) {
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

	var productWarnings []inventorydomain.Inventory
	productWarnings = make([]inventorydomain.Inventory, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var productWarning inventorydomain.Inventory
		if err = cursor.Decode(&productWarning); err != nil {
			return nil, err
		}

		if productWarning.Quantity <= productWarning.QuantityWarning {
			productWarnings = append(productWarnings, productWarning)
		}
	}

	return productWarnings, nil
}
