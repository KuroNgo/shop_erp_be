package benefit_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	benefit_controller "shop_erp_mono/api/controllers/human_resources_management/benefit"
	"shop_erp_mono/bootstrap"
	benefits_domain "shop_erp_mono/domain/human_resource_management/benefits"
	employeesdomain "shop_erp_mono/domain/human_resource_management/employees"
	benefit_repository "shop_erp_mono/repository/human_resource_management/benefit/repository"
	benefit_usecase "shop_erp_mono/usecase/human_resource_management/benefit/usecase"
	"time"
)

func BenefitRouter(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	be := benefit_repository.NewBenefitRepository(db, benefits_domain.CollectionBenefit, employeesdomain.CollectionEmployee)
	benefit := &benefit_controller.BenefitController{
		BenefitUseCase: benefit_usecase.NewBenefitUseCase(timeout, be),
		Database:       env,
	}

	router := group.Group("/benefits")
	router.GET("/get/_id", benefit.FetchOneBenefitByID)
	router.GET("/get/email", benefit.FetchOneBenefitByEmail)
	router.GET("/get/all", benefit.FetchAllBenefit)
	router.POST("/create", benefit.CreateOneBenefit)
	router.PUT("/update", benefit.UpdateOneBenefit)
	router.DELETE("/delete", benefit.DeleteOneBenefit)
}
