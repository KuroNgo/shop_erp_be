package payments_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	payment_controller "shop_erp_mono/internal/api/controllers/sales_and_distributing_management/payments"
	"shop_erp_mono/internal/config"
	payments_domain "shop_erp_mono/internal/domain/sales_and_distribution_management/payments"
	sale_orders_domain "shop_erp_mono/internal/domain/sales_and_distribution_management/sale_orders"
	payment_repository "shop_erp_mono/internal/repository/sales_and_distribution_management/payment/repository"
	sales_order_repository "shop_erp_mono/internal/repository/sales_and_distribution_management/sale_order/repository"
	payment_usecase "shop_erp_mono/internal/usecase/sales_and_distribution_management/payment/usecase"
	"time"
)

func PaymentRouter(env *config.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup, cacheTTL time.Duration) {
	pa := payment_repository.NewPaymentRepository(db, payments_domain.CollectionPayment)
	so := sales_order_repository.NewSaleOrderRepository(db, sale_orders_domain.CollectionSalesOrder)
	payment := &payment_controller.PaymentController{
		PaymentUseCase: payment_usecase.NewPaymentUseCase(timeout, pa, so, cacheTTL),
		Database:       env,
	}

	router := group.Group("/payments")
	router.GET("/get/_id", payment.GetByID)
	router.GET("/get/status", payment.GetByStatus)
	router.GET("/get/order_id", payment.GetByOrder)
	router.POST("/create", payment.CreateOne)
	router.PUT("/update", payment.UpdateOne)
	router.DELETE("/delete", payment.DeleteOne)
}
