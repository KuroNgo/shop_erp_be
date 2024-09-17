package unit

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	inventorydomain "shop_erp_mono/domain/warehouse_management/inventory"
	productdomain "shop_erp_mono/domain/warehouse_management/product"
	categorydomain "shop_erp_mono/domain/warehouse_management/product_category"
	warehousedomain "shop_erp_mono/domain/warehouse_management/warehouse"
	"shop_erp_mono/infrastructor"
	inventoryrepository "shop_erp_mono/repository/warehouse_management/inventory/repository"
	productrepository "shop_erp_mono/repository/warehouse_management/product/repository"
	category_repository "shop_erp_mono/repository/warehouse_management/product_category/repository"
	warehouserepository "shop_erp_mono/repository/warehouse_management/warehourse/repository"
	"testing"
	"time"
)

func TestCreateOneInventory(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	mockCategory := categorydomain.Category{
		ID:           primitive.NewObjectID(),
		CategoryName: "cà phê",
		Description:  "cà phê sữa ngon ngon",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	ca := category_repository.NewCategoryRepository(database, "product_category")
	err := ca.Create(context.Background(), mockCategory)

	mockProduct := productdomain.Product{
		ID:              primitive.NewObjectID(),
		ProductName:     "cà phê sữa",
		Description:     "cà phê sữa ngon ngon",
		Price:           12000,
		QuantityInStock: 120,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	pr := productrepository.NewProductRepository(database, "product")
	err = pr.CreateProduct(context.Background(), mockProduct)
	productData, err := pr.GetProductByName(context.Background(), "cà phê sữa")
	if err != nil || productData == nil {
		return
	}

	mockWarehouse := warehousedomain.Warehouse{
		ID:            primitive.NewObjectID(),
		WarehouseName: "trụ sở A",
		Location:      "Ho Chi Minh",
		Capacity:      12000,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	wr := warehouserepository.NewWarehouseRepository(database, "warehouse")
	err = wr.CreateOne(context.Background(), mockWarehouse)
	warehouseData, err := wr.GetByName(context.Background(), "trụ sở A")
	if err != nil || warehouseData == nil {
		return
	}

	mockInventory := inventorydomain.Inventory{
		ID:          primitive.NewObjectID(),
		ProductID:   productData.ID,
		WarehouseID: warehouseData.ID,
		Quantity:    100,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	mockEmptyInventory := inventorydomain.Inventory{}

	t.Run("success", func(t *testing.T) {
		ur := inventoryrepository.NewInventoryRepository(database, "inventory")
		err := ur.CreateOne(context.Background(), mockInventory)
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		ur := inventoryrepository.NewInventoryRepository(database, "inventory")
		err := ur.CreateOne(context.Background(), mockEmptyInventory)
		assert.Error(t, err)
	})
}
