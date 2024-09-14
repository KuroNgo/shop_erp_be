package stockmovement_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	stockmovement_controller "shop_erp_mono/api/controllers/warehouse_management/stockmovement"
	"shop_erp_mono/bootstrap"
	stockmovementdomain "shop_erp_mono/domain/warehouse_management/stockmovement"
	stockmovement_repository "shop_erp_mono/repository/warehouse_management/stockmovement/repository"
	stockmovement_usecase "shop_erp_mono/usecase/warehouse_management/stockmovement/usecase"
	"time"
)

func StockMovementRouter(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	st := stockmovement_repository.NewStockMovementRepository(db, stockmovementdomain.CollectionStockMovement)
	stockMovement := &stockmovement_controller.StockMovementController{
		StockMovementUseCase: stockmovement_usecase.NewStockMovementUseCase(timeout, st),
		Database:             env,
	}

	router := group.Group("/stock_movements")
	router.GET("/get/_id", stockMovement.GetByID)
	router.GET("/get/warehouse_id", stockMovement.GetByWarehouseID)
	router.GET("/get/product_id", stockMovement.GetByProductID)
	router.GET("/get/all", stockMovement.GetAllPagination)
	router.POST("/create", stockMovement.CreateOne)
	router.PUT("/update", stockMovement.UpdateOne)
	router.DELETE("/delete", stockMovement.DeleteOne)
}
