package department_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	departmentcontroller "shop_erp_mono/api/controllers/human_resources_management/department"
	"shop_erp_mono/bootstrap"
	departmentsdomain "shop_erp_mono/domain/human_resource_management/departments"
	employeesdomain "shop_erp_mono/domain/human_resource_management/employees"
	departmentrepository "shop_erp_mono/repository/human_resource_management/department/repository"
	employeerepository "shop_erp_mono/repository/human_resource_management/employee/repository"
	departmentusecase "shop_erp_mono/usecase/human_resource_management/department/usecase"
	"time"
)

func DepartmentRouter(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, client *mongo.Client, group *gin.RouterGroup, cacheTTL time.Duration) {
	de := departmentrepository.NewDepartmentRepository(db, departmentsdomain.CollectionDepartment)
	em := employeerepository.NewEmployeeRepository(db, employeesdomain.CollectionEmployee)
	department := &departmentcontroller.DepartmentController{
		DepartmentUseCase: departmentusecase.NewDepartmentUseCase(timeout, de, em, cacheTTL, client),
		Database:          env,
	}

	router := group.Group("/departments")
	router.GET("/get/_id", department.GetByID)
	router.GET("/get/name", department.GetByName)
	router.GET("/get/all", department.GetAll)
	router.POST("/create", department.CreateOne)
	router.POST("/create/manager", department.CreateOneWithManager)
	router.PUT("/update", department.UpdateOne)
	router.DELETE("/delete/_id", department.DeleteOne)
}
