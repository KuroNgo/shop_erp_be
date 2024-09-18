package shipping_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	shipping_controller "shop_erp_mono/api/controllers/sales_and_distributing_management/shipping"
	"shop_erp_mono/bootstrap"
	shippingdomain "shop_erp_mono/domain/sales_and_distribution_management/shipping"
	shipping_repository "shop_erp_mono/repository/sales_and_distribution_management/shipping/repository"
	shipping_usecase "shop_erp_mono/usecase/sales_and_distribution_management/shipping/usecase"
	"time"
)

func ShippingRouter(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	sh := shipping_repository.NewShippingRepository(db, shippingdomain.CollectionShipping)
	shipping := &shipping_controller.ShippingController{
		ShippingUseCase: shipping_usecase.NewShippingUseCase(timeout, sh),
		Database:        env,
	}

	router := group.Group("/shipping")
	router.GET("/get/_id", shipping.GetByID)
	router.GET("/get/status", shipping.GetByStatus)
	router.GET("/get/order_id", shipping.GetByOrderID)
	router.POST("/create", shipping.CreateOne)
	router.PUT("/update", shipping.UpdateOne)
	router.DELETE("/delete", shipping.DeleteOne)
}
