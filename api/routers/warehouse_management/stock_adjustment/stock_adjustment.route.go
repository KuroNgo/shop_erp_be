package stock_adjustment_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	stockadjustmentcontroller "shop_erp_mono/api/controllers/warehouse_management/stock_adjustment"
	"shop_erp_mono/bootstrap"
	productdomain "shop_erp_mono/domain/warehouse_management/product"
	stockadjustmentdomain "shop_erp_mono/domain/warehouse_management/stock_adjustment"
	warehousedomain "shop_erp_mono/domain/warehouse_management/warehouse"
	stockadjustmentrepository "shop_erp_mono/repository/warehouse_management/stock_adjustment/repository"
	warehouserepository "shop_erp_mono/repository/warehouse_management/warehourse/repository"
	productrepository "shop_erp_mono/repository/warehouse_management/wm_product/repository"
	stockadjustmentusecase "shop_erp_mono/usecase/warehouse_management/stock_adjustment/usecase"
	"time"
)

func StockAdjustmentRouter(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup, cacheTTL time.Duration) {
	st := stockadjustmentrepository.NewStockAdjustmentRepository(db, stockadjustmentdomain.CollectionStockAdjustment)
	pr := productrepository.NewProductRepository(db, productdomain.CollectionProduct)
	wa := warehouserepository.NewWarehouseRepository(db, warehousedomain.CollectionWareHouse)
	stockAdjustment := &stockadjustmentcontroller.StockAdjustmentController{
		StockAdjustmentUseCase: stockadjustmentusecase.NewStockAdjustmentUseCase(timeout, st, pr, wa, cacheTTL),
		Database:               env,
	}

	router := group.Group("/stock-adjustments")
	router.GET("/get/_id", stockAdjustment.GetByID)
	router.GET("/get/all", stockAdjustment.GetAll)
	router.POST("/create", stockAdjustment.CreateOne)
	router.PUT("/update", stockAdjustment.UpdateOne)
	router.DELETE("/delete", stockAdjustment.DeleteOne)
}
