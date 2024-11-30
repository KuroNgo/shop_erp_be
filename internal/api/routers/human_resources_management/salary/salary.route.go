package salary_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	salarycontroller "shop_erp_mono/internal/api/controllers/human_resources_management/salary"
	"shop_erp_mono/internal/config"
	employees_domain "shop_erp_mono/internal/domain/human_resource_management/employees"
	roledomain "shop_erp_mono/internal/domain/human_resource_management/role"
	salarydomain "shop_erp_mono/internal/domain/human_resource_management/salary"
	employeerepository "shop_erp_mono/internal/repository/human_resource_management/employee/repository"
	rolerepository "shop_erp_mono/internal/repository/human_resource_management/role/repository"
	salaryrepository "shop_erp_mono/internal/repository/human_resource_management/salary/repository"
	salaryusecase "shop_erp_mono/internal/usecase/human_resource_management/salary/usecase"
	"time"
)

func SalaryRouter(env *config.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup, cacheTTL time.Duration) {
	sa := salaryrepository.NewSalaryRepository(db, salarydomain.CollectionSalary)
	ro := rolerepository.NewRoleRepository(db, roledomain.CollectionRole)
	em := employeerepository.NewEmployeeRepository(db, employees_domain.CollectionEmployee)
	salary := &salarycontroller.SalaryController{
		SalaryUseCase: salaryusecase.NewSalaryUseCase(timeout, sa, ro, em, cacheTTL),
		Database:      env,
	}

	router := group.Group("/salaries")
	router.GET("/get/_id", salary.GetByID)
	router.GET("/get/title", salary.GetByRoleTitle)
	router.GET("/get/all", salary.GetAll)
	router.POST("/create", salary.CreateOne)
	router.PUT("/update", salary.UpdateOne)
	router.DELETE("/delete", salary.DeleteOne)

}
