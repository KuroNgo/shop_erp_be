package employee_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	employeecontroller "shop_erp_mono/internal/api/controllers/human_resources_management/employee"
	"shop_erp_mono/internal/config"
	departmentsdomain "shop_erp_mono/internal/domain/human_resource_management/departments"
	employeesdomain "shop_erp_mono/internal/domain/human_resource_management/employees"
	roledomain "shop_erp_mono/internal/domain/human_resource_management/role"
	salarydomain "shop_erp_mono/internal/domain/human_resource_management/salary"
	departmentrepository "shop_erp_mono/internal/repository/human_resource_management/department/repository"
	employeerepository "shop_erp_mono/internal/repository/human_resource_management/employee/repository"
	rolerepository "shop_erp_mono/internal/repository/human_resource_management/role/repository"
	salaryrepository "shop_erp_mono/internal/repository/human_resource_management/salary/repository"
	employeeusecase "shop_erp_mono/internal/usecase/human_resource_management/employee/usecase"
	"time"
)

func EmployeeRouter(env *config.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup, cacheTTL time.Duration) {
	em := employeerepository.NewEmployeeRepository(db, employeesdomain.CollectionEmployee)
	de := departmentrepository.NewDepartmentRepository(db, departmentsdomain.CollectionDepartment)
	sa := salaryrepository.NewSalaryRepository(db, salarydomain.CollectionSalary)
	ro := rolerepository.NewRoleRepository(db, roledomain.CollectionRole)
	employee := &employeecontroller.EmployeeController{
		EmployeeUseCase: employeeusecase.NewEmployeeUseCase(timeout, em, de, sa, ro, cacheTTL),
		Database:        env,
	}

	router := group.Group("/employees")
	router.GET("/get/_id", employee.GetByID)
	router.GET("/get/name", employee.GetByName)
	router.GET("/get/status", employee.GetByStatus)
	router.GET("/get/email", employee.GetByEmail)
	router.GET("/get/all", employee.GetAll)
	router.POST("/create", employee.CreateOne)
	router.PUT("/update", employee.UpdateOne)
	router.PATCH("/update/status", employee.UpdateStatus)
	router.PATCH("/delete-soft/_id", employee.DeleteSoft)
	router.DELETE("/delete/_id", employee.DeleteOne)
}
