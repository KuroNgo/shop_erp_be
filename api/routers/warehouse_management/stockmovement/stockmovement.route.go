package stockmovement_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	stockmovementcontroller "shop_erp_mono/api/controllers/warehouse_management/stockmovement"
	"shop_erp_mono/bootstrap"
	userdomain "shop_erp_mono/domain/human_resource_management/user"
	productdomain "shop_erp_mono/domain/warehouse_management/product"
	stockmovementdomain "shop_erp_mono/domain/warehouse_management/stockmovement"
	warehousedomain "shop_erp_mono/domain/warehouse_management/warehouse"
	userrepository "shop_erp_mono/repository/human_resource_management/user/repository"
	productrepository "shop_erp_mono/repository/warehouse_management/product/repository"
	stockmovementrepository "shop_erp_mono/repository/warehouse_management/stockmovement/repository"
	warehouserepository "shop_erp_mono/repository/warehouse_management/warehourse/repository"
	stockmovementusecase "shop_erp_mono/usecase/warehouse_management/stockmovement/usecase"
	"time"
)

func StockMovementRouter(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	st := stockmovementrepository.NewStockMovementRepository(db, stockmovementdomain.CollectionStockMovement)
	pr := productrepository.NewProductRepository(db, productdomain.CollectionProduct)
	ur := userrepository.NewUserRepository(db, userdomain.CollectionUser)
	wa := warehouserepository.NewWarehouseRepository(db, warehousedomain.CollectionWareHouse)
	stockMovement := &stockmovementcontroller.StockMovementController{
		StockMovementUseCase: stockmovementusecase.NewStockMovementUseCase(timeout, st, pr, ur, wa),
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
