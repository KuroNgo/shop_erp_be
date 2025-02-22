package shipping_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	shipping_controller "shop_erp_mono/internal/api/controllers/sales_and_distributing_management/shipping"
	"shop_erp_mono/internal/config"
	sale_orders_domain "shop_erp_mono/internal/domain/sales_and_distribution_management/sale_orders"
	shippingdomain "shop_erp_mono/internal/domain/sales_and_distribution_management/shipping"
	sales_order_repository "shop_erp_mono/internal/repository/sales_and_distribution_management/sale_order/repository"
	shipping_repository "shop_erp_mono/internal/repository/sales_and_distribution_management/shipping/repository"
	shipping_usecase "shop_erp_mono/internal/usecase/sales_and_distribution_management/shipping/usecase"
	"time"
)

func ShippingRouter(env *config.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup, cacheTTL time.Duration) {
	sh := shipping_repository.NewShippingRepository(db, shippingdomain.CollectionShipping)
	so := sales_order_repository.NewSaleOrderRepository(db, sale_orders_domain.CollectionSalesOrder)
	shipping := &shipping_controller.ShippingController{
		ShippingUseCase: shipping_usecase.NewShippingUseCase(timeout, sh, so, cacheTTL),
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
