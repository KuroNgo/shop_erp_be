package performance_review_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	performancereviewcontroller "shop_erp_mono/internal/api/controllers/human_resources_management/performance_review"
	"shop_erp_mono/internal/config"
	employeesdomain "shop_erp_mono/internal/domain/human_resource_management/employees"
	performancereviewdomain "shop_erp_mono/internal/domain/human_resource_management/performance_review"
	employeerepository "shop_erp_mono/internal/repository/human_resource_management/employee/repository"
	performancereviewrepository "shop_erp_mono/internal/repository/human_resource_management/performance_review/repository"
	performancereviewusecase "shop_erp_mono/internal/usecase/human_resource_management/performance_review/usecase"
	"time"
)

func PerformanceReviewRouterV2(env *config.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup, cacheTTL time.Duration) {
	pr := performancereviewrepository.NewPerformanceReviewRepository(db, performancereviewdomain.CollectionPerformanceReview)
	em := employeerepository.NewEmployeeRepository(db, employeesdomain.CollectionEmployee)
	performanceReview := &performancereviewcontroller.PerformanceReviewController{
		PerformanceReviewUseCase: performancereviewusecase.NewPerformanceReviewUseCase(timeout, pr, em, cacheTTL),
		Database:                 env,
	}

	router := group.Group("/performance-reviews")
	router.POST("/create", performanceReview.CreateOneWithIDEmployee)
	router.PUT("/update", performanceReview.CreateOneWithIDEmployee)
}
