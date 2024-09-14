package inventory_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	inventorycontroller "shop_erp_mono/api/controllers/warehouse_management/inventory"
	"shop_erp_mono/bootstrap"
	inventorydomain "shop_erp_mono/domain/warehouse_management/inventory"
	productdomain "shop_erp_mono/domain/warehouse_management/product"
	warehousedomain "shop_erp_mono/domain/warehouse_management/warehouse"
	inventoryrepository "shop_erp_mono/repository/warehouse_management/inventory/repository"
	productrepository "shop_erp_mono/repository/warehouse_management/product/repository"
	warehouserepository "shop_erp_mono/repository/warehouse_management/warehourse/repository"
	inventoryusecase "shop_erp_mono/usecase/warehouse_management/inventory/usecase"
	"time"
)

func InventoryRouter(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	inv := inventoryrepository.NewInventoryRepository(db, inventorydomain.CollectionInventory)
	pr := productrepository.NewProductRepository(db, productdomain.CollectionProduct)
	wa := warehouserepository.NewWarehouseRepository(db, warehousedomain.CollectionWareHouse)
	inventory := &inventorycontroller.InventoryController{
		InventoryUseCase: inventoryusecase.NewInventoryRepository(timeout, inv, pr, wa),
		Database:         env,
	}

	router := group.Group("/inventories")
	router.GET("/get/_id", inventory.GetByIDInventory)
	router.GET("/get/warehouse_id", inventory.GetByWarehouseID)
	router.GET("/get/product_id", inventory.GetByProductID)
	router.GET("/get/all", inventory.GetByProductID)
	router.GET("/get/check/availability", inventory.CheckInventoryAvailability)
	router.POST("/create", inventory.CreateInventory)
	router.PUT("/update", inventory.Update)
	router.PUT("/update/adjustment", inventory.AdjustmentQuantity)
	router.DELETE("/delete", inventory.DeleteInventory)
}
