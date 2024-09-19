package inventory_usecase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	inventory_domain "shop_erp_mono/domain/warehouse_management/inventory"
	productdomain "shop_erp_mono/domain/warehouse_management/product"
	warehousedomain "shop_erp_mono/domain/warehouse_management/warehouse"
	"shop_erp_mono/usecase/warehouse_management/inventory/validate"
	"time"
)

type inventoryUseCase struct {
	contextTimeout      time.Duration
	inventoryRepository inventory_domain.InventoryRepository
	productRepository   productdomain.IProductRepository
	warehouseRepository warehousedomain.IWarehouseRepository
}

func NewInventoryRepository(contextTimeout time.Duration, inventoryRepository inventory_domain.InventoryRepository,
	productRepository productdomain.IProductRepository, warehouseRepository warehousedomain.IWarehouseRepository) inventory_domain.InventoryUseCase {
	return &inventoryUseCase{contextTimeout: contextTimeout, inventoryRepository: inventoryRepository,
		productRepository: productRepository, warehouseRepository: warehouseRepository}
}

func (i *inventoryUseCase) CreateInventory(ctx context.Context, input *inventory_domain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, i.contextTimeout)
	defer cancel()

	if err := validate.ValidateInventory(input); err != nil {
		return err
	}

	productData, err := i.productRepository.GetByName(ctx, input.ProductName)
	if err != nil {
		return err
	}

	warehouseData, err := i.warehouseRepository.GetByName(ctx, input.WarehouseName)
	if err != nil {
		return err
	}

	inventory := inventory_domain.Inventory{
		ID:          primitive.NewObjectID(),
		ProductID:   productData.ID,
		WarehouseID: warehouseData.ID,
		Quantity:    input.Quantity,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return i.inventoryRepository.CreateOne(ctx, inventory)
}

func (i *inventoryUseCase) UpdateInventory(ctx context.Context, id string, input *inventory_domain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, i.contextTimeout)
	defer cancel()

	inventoryID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	if err := validate.ValidateInventory(input); err != nil {
		return err
	}

	productData, err := i.productRepository.GetByName(ctx, input.ProductName)
	if err != nil {
		return err
	}

	warehouseData, err := i.warehouseRepository.GetByName(ctx, input.WarehouseName)
	if err != nil {
		return err
	}

	inventory := inventory_domain.Inventory{
		ID:          inventoryID,
		ProductID:   productData.ID,
		WarehouseID: warehouseData.ID,
		Quantity:    input.Quantity,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return i.inventoryRepository.UpdateOne(ctx, inventory)
}

func (i *inventoryUseCase) GetInventoryByID(ctx context.Context, id string) (*inventory_domain.InventoryResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, i.contextTimeout)
	defer cancel()

	inventoryID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	inventoryData, err := i.inventoryRepository.GetByID(ctx, inventoryID)
	if err != nil {
		return nil, err
	}

	productData, err := i.productRepository.GetByID(ctx, inventoryData.ProductID)
	if err != nil {
		return nil, err
	}

	warehouseData, err := i.warehouseRepository.GetByID(ctx, inventoryData.WarehouseID)
	if err != nil {
		return nil, err
	}

	response := &inventory_domain.InventoryResponse{
		Inventory: *inventoryData,
		Product:   *productData,
		Warehouse: *warehouseData,
	}

	return response, nil
}

func (i *inventoryUseCase) GetInventoryByProduct(ctx context.Context, productID string) ([]inventory_domain.InventoryResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, i.contextTimeout)
	defer cancel()

	idProduct, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return nil, err
	}

	inventoryData, err := i.inventoryRepository.GetByProduct(ctx, idProduct)
	if err != nil {
		return nil, err
	}

	var responses []inventory_domain.InventoryResponse
	responses = make([]inventory_domain.InventoryResponse, 0, len(inventoryData))
	for _, inventory := range inventoryData {
		productData, err := i.productRepository.GetByID(ctx, inventory.ProductID)
		if err != nil {
			return nil, err
		}

		warehouseData, err := i.warehouseRepository.GetByID(ctx, inventory.WarehouseID)
		if err != nil {
			return nil, err
		}

		response := inventory_domain.InventoryResponse{
			Inventory: inventory,
			Product:   *productData,
			Warehouse: *warehouseData,
		}

		responses = append(responses, response)

	}

	return responses, nil
}

func (i *inventoryUseCase) GetInventoryByWarehouse(ctx context.Context, warehouseID string) ([]inventory_domain.InventoryResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, i.contextTimeout)
	defer cancel()

	idWarehouse, err := primitive.ObjectIDFromHex(warehouseID)
	if err != nil {
		return nil, err
	}

	inventoryData, err := i.inventoryRepository.GetByWarehouse(ctx, idWarehouse)
	if err != nil {
		return nil, err
	}

	var responses []inventory_domain.InventoryResponse
	responses = make([]inventory_domain.InventoryResponse, 0, len(inventoryData))
	for _, inventory := range inventoryData {
		productData, err := i.productRepository.GetByID(ctx, inventory.ProductID)
		if err != nil {
			return nil, err
		}

		warehouseData, err := i.warehouseRepository.GetByID(ctx, inventory.WarehouseID)
		if err != nil {
			return nil, err
		}

		response := inventory_domain.InventoryResponse{
			Inventory: inventory,
			Product:   *productData,
			Warehouse: *warehouseData,
		}

		responses = append(responses, response)

	}

	return responses, nil
}

func (i *inventoryUseCase) DeleteInventory(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, i.contextTimeout)
	defer cancel()

	inventoryID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return i.inventoryRepository.DeleteOne(ctx, inventoryID)
}

func (i *inventoryUseCase) AdjustInventoryQuantity(ctx context.Context, id string, adjustment int) (*inventory_domain.InventoryResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, i.contextTimeout)
	defer cancel()

	warehouseID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	inventoryData, err := i.inventoryRepository.AdjustQuantity(ctx, warehouseID, adjustment)
	if err != nil {
		return nil, err
	}

	productData, err := i.productRepository.GetByID(ctx, inventoryData.ProductID)
	if err != nil {
		return nil, err
	}

	warehouseData, err := i.warehouseRepository.GetByID(ctx, inventoryData.WarehouseID)
	if err != nil {
		return nil, err
	}

	response := &inventory_domain.InventoryResponse{
		Inventory: *inventoryData,
		Product:   *productData,
		Warehouse: *warehouseData,
	}

	return response, nil
}

func (i *inventoryUseCase) CheckInventoryAvailability(ctx context.Context, productID string, warehouseID string, requiredQuantity int) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, i.contextTimeout)
	defer cancel()

	idWarehouse, err := primitive.ObjectIDFromHex(warehouseID)
	if err != nil {
		return false, err
	}

	idProduct, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return false, err
	}

	return i.inventoryRepository.CheckAvailability(ctx, idProduct, idWarehouse, requiredQuantity)
}

func (i *inventoryUseCase) ListAllInventories(ctx context.Context) ([]inventory_domain.InventoryResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, i.contextTimeout)
	defer cancel()

	inventoryData, err := i.inventoryRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []inventory_domain.InventoryResponse
	responses = make([]inventory_domain.InventoryResponse, 0, len(inventoryData))
	for _, inventory := range inventoryData {
		productData, err := i.productRepository.GetByID(ctx, inventory.ProductID)
		if err != nil {
			return nil, err
		}

		warehouseData, err := i.warehouseRepository.GetByID(ctx, inventory.WarehouseID)
		if err != nil {
			return nil, err
		}

		response := inventory_domain.InventoryResponse{
			Inventory: inventory,
			Product:   *productData,
			Warehouse: *warehouseData,
		}

		responses = append(responses, response)

	}

	return responses, nil
}
