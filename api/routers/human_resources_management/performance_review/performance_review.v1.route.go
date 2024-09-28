package performance_review_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	performancereviewcontroller "shop_erp_mono/api/controllers/human_resources_management/performance_review"
	"shop_erp_mono/bootstrap"
	employeesdomain "shop_erp_mono/domain/human_resource_management/employees"
	performancereviewdomain "shop_erp_mono/domain/human_resource_management/performance_review"
	employeerepository "shop_erp_mono/repository/human_resource_management/employee/repository"
	performancereviewrepository "shop_erp_mono/repository/human_resource_management/performance_review/repository"
	performancereviewusecase "shop_erp_mono/usecase/human_resource_management/performance_review/usecase"
	"time"
)

func PerformanceReviewRouterV1(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup, cacheTTL time.Duration) {
	pr := performancereviewrepository.NewPerformanceReviewRepository(db, performancereviewdomain.CollectionPerformanceReview)
	em := employeerepository.NewEmployeeRepository(db, employeesdomain.CollectionEmployee)
	performanceReview := &performancereviewcontroller.PerformanceReviewController{
		PerformanceReviewUseCase: performancereviewusecase.NewPerformanceReviewUseCase(timeout, pr, em, cacheTTL),
		Database:                 env,
	}

	router := group.Group("/performance-reviews")
	router.GET("/get/_id", performanceReview.GetByID)
	router.GET("/get/email", performanceReview.GetByEmailEmployee)
	router.GET("/get/all", performanceReview.GetAll)
	router.POST("/create", performanceReview.CreateOneWithIDEmployee)
	router.PUT("/update", performanceReview.UpdateOneWithIDEmployee)
	router.DELETE("/delete", performanceReview.DeleteOne)
}
