package stockmovement_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	stockmovementcontroller "shop_erp_mono/internal/api/controllers/warehouse_management/stockmovement"
	"shop_erp_mono/internal/config"
	userdomain "shop_erp_mono/internal/domain/human_resource_management/user"
	productdomain "shop_erp_mono/internal/domain/warehouse_management/product"
	stockmovementdomain "shop_erp_mono/internal/domain/warehouse_management/stockmovement"
	warehousedomain "shop_erp_mono/internal/domain/warehouse_management/warehouse"
	userrepository "shop_erp_mono/internal/repository/human_resource_management/user/repository"
	stockmovementrepository "shop_erp_mono/internal/repository/warehouse_management/stockmovement/repository"
	warehouserepository "shop_erp_mono/internal/repository/warehouse_management/warehourse/repository"
	productrepository "shop_erp_mono/internal/repository/warehouse_management/wm_product/repository"
	stockmovementusecase "shop_erp_mono/internal/usecase/warehouse_management/stockmovement/usecase"
	"time"
)

func StockMovementRouter(env *config.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup, cacheTTL time.Duration) {
	st := stockmovementrepository.NewStockMovementRepository(db, stockmovementdomain.CollectionStockMovement)
	pr := productrepository.NewProductRepository(db, productdomain.CollectionProduct)
	ur := userrepository.NewUserRepository(db, userdomain.CollectionUser)
	wa := warehouserepository.NewWarehouseRepository(db, warehousedomain.CollectionWareHouse)
	stockMovement := &stockmovementcontroller.StockMovementController{
		StockMovementUseCase: stockmovementusecase.NewStockMovementUseCase(timeout, st, pr, ur, wa, cacheTTL),
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
