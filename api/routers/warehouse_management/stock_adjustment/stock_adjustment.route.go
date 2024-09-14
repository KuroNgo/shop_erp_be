package stock_adjustment_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	stock_adjustment_controller "shop_erp_mono/api/controllers/warehouse_management/stock_adjustment"
	"shop_erp_mono/bootstrap"
	stock_adjustment_domain "shop_erp_mono/domain/warehouse_management/stock_adjustment"
	stock_adjustment_repository "shop_erp_mono/repository/warehouse_management/stock_adjustment/repository"
	stock_adjustment_usecase "shop_erp_mono/usecase/warehouse_management/stock_adjustment/usecase"
	"time"
)

func StockAdjustmentRouter(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	st := stock_adjustment_repository.NewStockAdjustmentRepository(db, stock_adjustment_domain.CollectionStockAdjustment)
	stockAdjustment := &stock_adjustment_controller.StockAdjustmentController{
		StockAdjustmentUseCase: stock_adjustment_usecase.NewStockAdjustmentUseCase(timeout, st),
		Database:               env,
	}

	router := group.Group("/stock_adjustments")
	router.GET("/get/_id", stockAdjustment.GetByID)
	router.GET("/get/all", stockAdjustment.GetAll)
	router.POST("/create", stockAdjustment.CreateOne)
	router.PUT("/update", stockAdjustment.UpdateOne)
	router.DELETE("/delete", stockAdjustment.DeleteOne)
}
