package order_details_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	orderdetailcontroller "shop_erp_mono/api/controllers/sales_and_distributing_management/order_details"
	"shop_erp_mono/bootstrap"
	orderdetailsdomain "shop_erp_mono/domain/sales_and_distribution_management/order_details"
	orderdetailrepository "shop_erp_mono/repository/sales_and_distribution_management/order_details/repository"
	orderdetailusecase "shop_erp_mono/usecase/sales_and_distribution_management/order_details/usecase"
	"time"
)

func OrderDetailRouter(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	od := orderdetailrepository.NewOrderDetailRepository(db, orderdetailsdomain.CollectionOrderDetail)
	orderDetail := &orderdetailcontroller.OrderDetailController{
		OrderDetailUseCase: orderdetailusecase.NewOrderDetailUseCase(timeout, od),
		Database:           env,
	}

	router := group.Group("/order_details")
	router.GET("/get/_id", orderDetail.GetByID)
	router.GET("/get/product_id", orderDetail.GetByOrderID)
	router.GET("/get/order_id", orderDetail.GetByProductID)
	router.POST("/create", orderDetail.CreateOne)
	router.PUT("/update", orderDetail.UpdateOne)
	router.DELETE("/delete", orderDetail.DeleteOne)
}
