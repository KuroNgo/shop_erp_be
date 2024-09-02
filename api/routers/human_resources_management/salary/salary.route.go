package salary_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	salarycontroller "shop_erp_mono/api/controllers/human_resources_management/salary"
	"shop_erp_mono/bootstrap"
	roledomain "shop_erp_mono/domain/human_resource_management/role"
	salarydomain "shop_erp_mono/domain/human_resource_management/salary"
	rolerepository "shop_erp_mono/repository/human_resource_management/role/repository"
	salaryrepository "shop_erp_mono/repository/human_resource_management/salary/repository"
	salaryusecase "shop_erp_mono/usecase/human_resource_management/salary/usecase"
	"time"
)

func SalaryRouter(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	sa := salaryrepository.NewSalaryRepository(db, salarydomain.CollectionSalary)
	ro := rolerepository.NewRoleRepository(db, roledomain.CollectionRole)
	salary := &salarycontroller.SalaryController{
		SalaryUseCase: salaryusecase.NewSalaryUseCase(timeout, sa, ro),
		Database:      env,
	}

	router := group.Group("/salaries")
	router.GET("/get/_id", salary.GetOneSalaryByID)
	router.GET("/get/title", salary.GetOneSalaryByRole)
	router.GET("/get/all", salary.GetAllSalary)
	router.POST("/create", salary.CreateOneSalary)
	router.PUT("/update", salary.UpdateOneSalary)
	router.DELETE("/delete", salary.DeleteOneSalary)

}
