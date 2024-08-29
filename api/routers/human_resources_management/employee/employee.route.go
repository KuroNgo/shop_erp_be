package employee_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	employeecontroller "shop_erp_mono/api/controllers/human_resources_management/employee"
	"shop_erp_mono/bootstrap"
	departmentsdomain "shop_erp_mono/domain/human_resource_management/departments"
	employeesdomain "shop_erp_mono/domain/human_resource_management/employees"
	roledomain "shop_erp_mono/domain/human_resource_management/role"
	salarydomain "shop_erp_mono/domain/human_resource_management/salary"
	employee_repository "shop_erp_mono/repository/human_resource_management/employee/repository"
	employee_usecase "shop_erp_mono/usecase/human_resource_management/employee/usecase"
	"time"
)

func EmployeeRouter(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	em := employee_repository.NewEmployeeRepository(db, employeesdomain.CollectionEmployee, departmentsdomain.CollectionDepartment, roledomain.CollectionRole, salarydomain.CollectionSalary)
	employee := &employeecontroller.EmployeeController{
		EmployeeUseCase: employee_usecase.NewEmployeeUseCase(timeout, em),
		Database:        env,
	}

	router := group.Group("/employees")
	router.GET("/get/_id", employee.FetchByIDEmployee)
	router.GET("/get/name", employee.FetchByNameEmployee)
	router.GET("/get/all", employee.FetchAllEmployee)
	router.POST("/create", employee.CreateOneEmployee)
	router.PUT("/update", employee.UpdateOneEmployee)
	router.DELETE("/delete", employee.DeleteOneEmployee)
}
