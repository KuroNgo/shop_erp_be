package supplier_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	suppliercontroller "shop_erp_mono/api/controllers/warehouse_management/supplier"
	"shop_erp_mono/bootstrap"
	supplierdomain "shop_erp_mono/domain/warehouse_management/supplier"
	supplierrepository "shop_erp_mono/repository/warehouse_management/supplier/repository"
	supplierusecase "shop_erp_mono/usecase/warehouse_management/supplier/usecase"
	"time"
)

func SupplierRouter(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	su := supplierrepository.NewSupplierRepository(db, supplierdomain.CollectionSupplier)

	supplier := &suppliercontroller.SupplierController{
		SupplierUseCase: supplierusecase.NewSupplierUseCase(timeout, su),
		Database:        env,
	}

	router := group.Group("/suppliers")
	router.GET("/get/_id", supplier.GetByIDSupplier)
	router.GET("/get/name", supplier.GetByNameSupplier)
	router.GET("/get/all/pagination", supplier.GetAllSupplierWithPagination)
	router.GET("/get/all", supplier.GetAll)
	router.POST("/create", supplier.CreateSupplier)
	router.PUT("/update", supplier.UpdateSupplier)
	router.DELETE("/delete", supplier.DeleteSupplier)
}
