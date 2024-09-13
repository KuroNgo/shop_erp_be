package inventory_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InventoryRepository interface {
	CreateOne(ctx context.Context, inventory Inventory) error
	UpdateOne(ctx context.Context, inventory Inventory) error
	GetByID(ctx context.Context, id primitive.ObjectID) (*Inventory, error)
	GetByProduct(ctx context.Context, productID primitive.ObjectID) ([]Inventory, error)
	GetByWarehouse(ctx context.Context, warehouseID primitive.ObjectID) ([]Inventory, error)
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	AdjustQuantity(ctx context.Context, id primitive.ObjectID, adjustment int) (*Inventory, error)
	CheckAvailability(ctx context.Context, productID primitive.ObjectID, warehouseID primitive.ObjectID, requiredQuantity int) (bool, error)
	GetAll(ctx context.Context) ([]Inventory, error)
}

type InventoryUseCase interface {
	CreateInventory(ctx context.Context, input *Input) (*InventoryResponse, error)
	UpdateInventory(ctx context.Context, id primitive.ObjectID, input *Input) (*InventoryResponse, error)
	GetInventoryByID(ctx context.Context, id primitive.ObjectID) (*InventoryResponse, error)
	GetInventoryByProduct(ctx context.Context, productID primitive.ObjectID) ([]InventoryResponse, error)
	GetInventoryByWarehouse(ctx context.Context, warehouseID primitive.ObjectID) ([]InventoryResponse, error)
	DeleteInventory(ctx context.Context, id primitive.ObjectID) error
	AdjustInventoryQuantity(ctx context.Context, id primitive.ObjectID, adjustment int) (*InventoryResponse, error)
	CheckInventoryAvailability(ctx context.Context, productID primitive.ObjectID, warehouseID primitive.ObjectID, requiredQuantity int) (bool, error)
	ListAllInventories(ctx context.Context) ([]InventoryResponse, error)
}
