package benefit_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	benefitcontroller "shop_erp_mono/api/controllers/human_resources_management/benefit"
	"shop_erp_mono/bootstrap"
	benefitsdomain "shop_erp_mono/domain/human_resource_management/benefits"
	employeesdomain "shop_erp_mono/domain/human_resource_management/employees"
	benefitrepository "shop_erp_mono/repository/human_resource_management/benefit/repository"
	employeerepository "shop_erp_mono/repository/human_resource_management/employee/repository"
	benefitusecase "shop_erp_mono/usecase/human_resource_management/benefit/usecase"
	"time"
)

func BenefitRouter(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	be := benefitrepository.NewBenefitRepository(db, benefitsdomain.CollectionBenefit, employeesdomain.CollectionEmployee)
	em := employeerepository.NewEmployeeRepository(db, employeesdomain.CollectionEmployee)
	benefit := &benefitcontroller.BenefitController{
		BenefitUseCase: benefitusecase.NewBenefitUseCase(timeout, be, em),
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
