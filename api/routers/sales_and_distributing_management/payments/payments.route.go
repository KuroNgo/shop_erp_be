package payments_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	payment_controller "shop_erp_mono/api/controllers/sales_and_distributing_management/payments"
	"shop_erp_mono/bootstrap"
	payments_domain "shop_erp_mono/domain/sales_and_distribution_management/payments"
	payment_repository "shop_erp_mono/repository/sales_and_distribution_management/payment/repository"
	payment_usecase "shop_erp_mono/usecase/sales_and_distribution_management/payment/usecase"
	"time"
)

func PaymentRouter(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	pa := payment_repository.NewPaymentRepository(db, payments_domain.CollectionPayment)
	payment := &payment_controller.PaymentController{
		PaymentUseCase: payment_usecase.NewPaymentUseCase(timeout, pa),
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
