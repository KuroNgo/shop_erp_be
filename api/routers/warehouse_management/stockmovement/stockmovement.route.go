package stockmovement_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	stockmovement_controller "shop_erp_mono/api/controllers/warehouse_management/stockmovement"
	"shop_erp_mono/bootstrap"
	userdomain "shop_erp_mono/domain/human_resource_management/user"
	productdomain "shop_erp_mono/domain/warehouse_management/product"
	stockmovementdomain "shop_erp_mono/domain/warehouse_management/stockmovement"
	warehousedomain "shop_erp_mono/domain/warehouse_management/warehouse"
	userrepository "shop_erp_mono/repository/human_resource_management/user/repository"
	product_repository "shop_erp_mono/repository/warehouse_management/product/repository"
	stockmovement_repository "shop_erp_mono/repository/warehouse_management/stockmovement/repository"
	warehouserepository "shop_erp_mono/repository/warehouse_management/warehourse/repository"
	stockmovement_usecase "shop_erp_mono/usecase/warehouse_management/stockmovement/usecase"
	"time"
)

func StockMovementRouter(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	st := stockmovement_repository.NewStockMovementRepository(db, stockmovementdomain.CollectionStockMovement)
	pr := product_repository.NewProductRepository(db, productdomain.CollectionProduct)
	ur := userrepository.NewUserRepository(db, userdomain.CollectionUser)
	wa := warehouserepository.NewWarehouseRepository(db, warehousedomain.CollectionWareHouse)
	stockMovement := &stockmovement_controller.StockMovementController{
		StockMovementUseCase: stockmovement_usecase.NewStockMovementUseCase(timeout, st, pr, ur, wa),
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
