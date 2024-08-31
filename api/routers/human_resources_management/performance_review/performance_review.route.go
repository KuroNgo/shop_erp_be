package performance_review_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	performancereviewcontroller "shop_erp_mono/api/controllers/human_resources_management/performance_review"
	"shop_erp_mono/bootstrap"
	employeesdomain "shop_erp_mono/domain/human_resource_management/employees"
	performancereviewdomain "shop_erp_mono/domain/human_resource_management/performance_review"
	performancereviewrepository "shop_erp_mono/repository/human_resource_management/performance_review/repository"
	performancereviewusecase "shop_erp_mono/usecase/human_resource_management/performance_review/usecase"
	"time"
)

func PerformanceReviewRouter(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	pr := performancereviewrepository.NewPerformanceReviewRepository(db, performancereviewdomain.CollectionPerformanceReview, employeesdomain.CollectionEmployee)
	performanceReview := &performancereviewcontroller.PerformanceReviewController{
		PerformanceReviewUseCase: performancereviewusecase.NewPerformanceReviewUseCase(timeout, pr),
		Database:                 env,
	}

	router := group.Group("/performance_reviews")
	router.GET("/get/_id", performanceReview.GetOneByIDPerformanceReview)
	router.GET("/get/email", performanceReview.GetOneByEmailPerformanceReview)
	router.GET("/get/all", performanceReview.GetAllPerformanceReview)
	router.POST("/create", performanceReview.CreateOnePerformanceReview)
	router.PUT("/update", performanceReview.UpdateOnePerformanceReview)
	router.DELETE("/delete", performanceReview.DeleteOnePerformanceReview)
}
