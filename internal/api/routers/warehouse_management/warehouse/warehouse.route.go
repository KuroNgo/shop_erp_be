package warehouse_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	warehousecontroller "shop_erp_mono/internal/api/controllers/warehouse_management/warehouse"
	"shop_erp_mono/internal/config"
	warehousedomain "shop_erp_mono/internal/domain/warehouse_management/warehouse"
	warehouserepository "shop_erp_mono/internal/repository/warehouse_management/warehourse/repository"
	warehouseusecase "shop_erp_mono/internal/usecase/warehouse_management/warehourse/usecase"
	"time"
)

func WarehouseRouter(env *config.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup, cacheTTL time.Duration) {
	wa := warehouserepository.NewWarehouseRepository(db, warehousedomain.CollectionWareHouse)
	warehouse := &warehousecontroller.WarehouseController{
		WarehouseUseCase: warehouseusecase.NewWarehouseUseCase(timeout, wa, cacheTTL),
		Database:         env,
	}

	router := group.Group("/warehouses")
	router.GET("/get/_id", warehouse.GetByID)
	router.GET("/get/name", warehouse.GetByName)
	router.GET("/get/all", warehouse.GetAll)
	router.POST("/create", warehouse.CreateOne)
	router.PUT("/update", warehouse.UpdateOne)
	router.DELETE("/delete", warehouse.DeleteOne)
}
