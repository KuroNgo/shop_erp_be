package benefit_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	benefitcontroller "shop_erp_mono/internal/api/controllers/human_resources_management/benefit"
	"shop_erp_mono/internal/config"
	benefitsdomain "shop_erp_mono/internal/domain/human_resource_management/benefits"
	employeesdomain "shop_erp_mono/internal/domain/human_resource_management/employees"
	benefitrepository "shop_erp_mono/internal/repository/human_resource_management/benefit/repository"
	employeerepository "shop_erp_mono/internal/repository/human_resource_management/employee/repository"
	benefitusecase "shop_erp_mono/internal/usecase/human_resource_management/benefit/usecase"
	"time"
)

func BenefitRouter(env *config.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup, cacheTTL time.Duration) {
	be := benefitrepository.NewBenefitRepository(db, benefitsdomain.CollectionBenefit)
	em := employeerepository.NewEmployeeRepository(db, employeesdomain.CollectionEmployee)
	benefit := &benefitcontroller.BenefitController{
		BenefitUseCase: benefitusecase.NewBenefitUseCase(timeout, be, em, cacheTTL),
		Database:       env,
	}

	router := group.Group("/benefits")
	router.GET("/get/_id", benefit.GetByID)
	router.GET("/get/email", benefit.GetByEmail)
	router.GET("/get/all", benefit.GetAll)
	router.POST("/create", benefit.CreateOne)
	router.PUT("/update", benefit.UpdateOne)
	router.DELETE("/delete", benefit.DeleteOne)
}
