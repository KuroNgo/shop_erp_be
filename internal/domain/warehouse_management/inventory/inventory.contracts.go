package inventory_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InventoryRepository interface {
	GetByID(ctx context.Context, id primitive.ObjectID) (*Inventory, error)
	GetByProduct(ctx context.Context, productID primitive.ObjectID) ([]Inventory, error)
	GetByWarehouse(ctx context.Context, warehouseID primitive.ObjectID) ([]Inventory, error)
	GetAll(ctx context.Context) ([]Inventory, error)

	CreateOne(ctx context.Context, inventory Inventory) error
	UpdateOne(ctx context.Context, inventory Inventory) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error

	CheckAvailability(ctx context.Context, productID primitive.ObjectID, warehouseID primitive.ObjectID, requiredQuantity int) (bool, error)
	AdjustQuantity(ctx context.Context, id primitive.ObjectID, adjustment int) (*Inventory, error)
}

type InventoryUseCase interface {
	GetInventoryByID(ctx context.Context, id string) (*InventoryResponse, error)
	GetInventoryByProduct(ctx context.Context, productID string) ([]InventoryResponse, error)
	GetInventoryByWarehouse(ctx context.Context, warehouseID string) ([]InventoryResponse, error)
	ListAllInventories(ctx context.Context) ([]InventoryResponse, error)

	CreateInventory(ctx context.Context, input *Input) error
	UpdateInventory(ctx context.Context, id string, input *Input) error
	DeleteInventory(ctx context.Context, id string) error

	AdjustInventoryQuantity(ctx context.Context, id string, adjustment int) (*InventoryResponse, error)
	CheckInventoryAvailability(ctx context.Context, productID string, warehouseID string, requiredQuantity int) (bool, error)
}
