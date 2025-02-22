package customer_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	customercontroller "shop_erp_mono/internal/api/controllers/sales_and_distributing_management/customer"
	"shop_erp_mono/internal/config"
	customerdomain "shop_erp_mono/internal/domain/sales_and_distribution_management/customer"
	customerrepository "shop_erp_mono/internal/repository/sales_and_distribution_management/customer/repository"
	customerusecase "shop_erp_mono/internal/usecase/sales_and_distribution_management/customer/usecase"
	"time"
)

func CustomerRouter(env *config.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup, cacheTTL time.Duration) {
	cu := customerrepository.NewCustomerRepository(db, customerdomain.CollectionCustomer)
	customer := &customercontroller.CustomerController{
		CustomerUseCase: customerusecase.NewCustomerUseCase(timeout, cu, cacheTTL),
		Database:        env,
	}

	router := group.Group("/customers")
	router.GET("/get/_id", customer.GetOneByID)
	router.GET("/get/name", customer.GetOneByName)
	router.GET("/get/all", customer.GetAll)
	router.POST("/create", customer.CreateOne)
	router.PUT("/update", customer.UpdateOne)
	router.DELETE("/delete", customer.DeleteOne)
}
