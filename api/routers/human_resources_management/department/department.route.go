package department_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	departmentcontroller "shop_erp_mono/api/controllers/human_resources_management/department"
	"shop_erp_mono/bootstrap"
	departmentsdomain "shop_erp_mono/domain/human_resource_management/departments"
	employees_domain "shop_erp_mono/domain/human_resource_management/employees"
	departmentrepository "shop_erp_mono/repository/human_resource_management/department/repository"
	employeerepository "shop_erp_mono/repository/human_resource_management/employee/repository"
	departmentusecase "shop_erp_mono/usecase/human_resource_management/department/usecase"
	"time"
)

func DepartmentRouter(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	de := departmentrepository.NewDepartmentRepository(db, departmentsdomain.CollectionDepartment)
	em := employeerepository.NewEmployeeRepository(db, employees_domain.CollectionEmployee)
	department := &departmentcontroller.DepartmentController{
		DepartmentUseCase: departmentusecase.NewDepartmentUseCase(timeout, de, em),
		Database:          env,
	}

	router := group.Group("/departments")
	router.GET("/get/_id", department.FetchOneDepartmentByID)
	router.GET("/get/name", department.FetchOneDepartmentByName)
	router.GET("/get/all", department.FetchAllDepartment)
	router.POST("/create", department.CreateOneDepartment)
	router.PUT("/update", department.UpdateOneDepartment)
	router.DELETE("/delete", department.DeleteOneDepartment)
}
