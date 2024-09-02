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
	departmentrepository "shop_erp_mono/repository/human_resource_management/department/repository"
	employeerepository "shop_erp_mono/repository/human_resource_management/employee/repository"
	rolerepository "shop_erp_mono/repository/human_resource_management/role/repository"
	salaryrepository "shop_erp_mono/repository/human_resource_management/salary/repository"
	employeeusecase "shop_erp_mono/usecase/human_resource_management/employee/usecase"
	"time"
)

func EmployeeRouter(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	em := employeerepository.NewEmployeeRepository(db, employeesdomain.CollectionEmployee)
	de := departmentrepository.NewDepartmentRepository(db, departmentsdomain.CollectionDepartment)
	sa := salaryrepository.NewSalaryRepository(db, salarydomain.CollectionSalary)
	ro := rolerepository.NewRoleRepository(db, roledomain.CollectionRole)
	employee := &employeecontroller.EmployeeController{
		EmployeeUseCase: employeeusecase.NewEmployeeUseCase(timeout, em, de, sa, ro),
		Database:        env,
	}

	router := group.Group("/employees")
	router.GET("/get/_id", employee.FetchByIDEmployee)
	router.GET("/get/email", employee.FetchByEmailEmployee)
	router.GET("/get/all", employee.FetchAllEmployee)
	router.POST("/create", employee.CreateOneEmployee)
	router.PUT("/update", employee.UpdateOneEmployee)
	router.DELETE("/delete", employee.DeleteOneEmployee)
}
