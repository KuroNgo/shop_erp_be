package inventory_usecase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	inventory_domain "shop_erp_mono/domain/warehouse_management/inventory"
	"time"
)

type inventoryUseCase struct {
	contextTimeout      time.Duration
	inventoryRepository inventory_domain.InventoryRepository
}

func NewInventoryRepository(contextTimeout time.Duration, inventoryRepository inventory_domain.InventoryRepository) inventory_domain.InventoryUseCase {
	return &inventoryUseCase{contextTimeout: contextTimeout, inventoryRepository: inventoryRepository}
}

func (i *inventoryUseCase) CreateInventory(ctx context.Context, input *inventory_domain.Input) (*inventory_domain.InventoryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (i *inventoryUseCase) UpdateInventory(ctx context.Context, id primitive.ObjectID, input *inventory_domain.Input) (*inventory_domain.InventoryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (i *inventoryUseCase) GetInventoryByID(ctx context.Context, id primitive.ObjectID) (*inventory_domain.InventoryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (i *inventoryUseCase) GetInventoryByProduct(ctx context.Context, productID primitive.ObjectID) ([]inventory_domain.InventoryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (i *inventoryUseCase) GetInventoryByWarehouse(ctx context.Context, warehouseID primitive.ObjectID) ([]inventory_domain.InventoryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (i *inventoryUseCase) DeleteInventory(ctx context.Context, id primitive.ObjectID) error {
	//TODO implement me
	panic("implement me")
}

func (i *inventoryUseCase) AdjustInventoryQuantity(ctx context.Context, id primitive.ObjectID, adjustment int) (*inventory_domain.InventoryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (i *inventoryUseCase) CheckInventoryAvailability(ctx context.Context, productID primitive.ObjectID, warehouseID primitive.ObjectID, requiredQuantity int) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (i *inventoryUseCase) ListAllInventories(ctx context.Context) ([]inventory_domain.InventoryResponse, error) {
	//TODO implement me
	panic("implement me")
}
