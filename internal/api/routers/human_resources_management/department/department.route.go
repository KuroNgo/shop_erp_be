package department_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	departmentcontroller "shop_erp_mono/internal/api/controllers/human_resources_management/department"
	"shop_erp_mono/internal/config"
	departmentsdomain "shop_erp_mono/internal/domain/human_resource_management/departments"
	employeesdomain "shop_erp_mono/internal/domain/human_resource_management/employees"
	roledomain "shop_erp_mono/internal/domain/human_resource_management/role"
	userdomain "shop_erp_mono/internal/domain/human_resource_management/user"
	departmentrepository "shop_erp_mono/internal/repository/human_resource_management/department/repository"
	employeerepository "shop_erp_mono/internal/repository/human_resource_management/employee/repository"
	rolerepository "shop_erp_mono/internal/repository/human_resource_management/role/repository"
	userrepository "shop_erp_mono/internal/repository/human_resource_management/user/repository"
	departmentusecase "shop_erp_mono/internal/usecase/human_resource_management/department/usecase"
	"time"
)

func DepartmentRouter(env *config.Database, timeout time.Duration, db *mongo.Database, client *mongo.Client, group *gin.RouterGroup, cacheTTL time.Duration) {
	de := departmentrepository.NewDepartmentRepository(db, departmentsdomain.CollectionDepartment)
	em := employeerepository.NewEmployeeRepository(db, employeesdomain.CollectionEmployee)
	ro := rolerepository.NewRoleRepository(db, roledomain.CollectionRole)
	us := userrepository.NewUserRepository(db, userdomain.CollectionUser)
	department := &departmentcontroller.DepartmentController{
		DepartmentUseCase: departmentusecase.NewDepartmentUseCase(timeout, de, em, us, ro, cacheTTL, client),
		Database:          env,
	}

	router := group.Group("/departments")
	router.GET("/get/_id", department.GetByID)
	router.GET("/get/status", department.GetByStatus)
	router.GET("/get/delete", department.GetAllSoftDelete)
	router.GET("/get/name", department.GetByName)
	router.GET("/get/all", department.GetAll)
	router.POST("/create", department.CreateOne)
	router.POST("/create/manager", department.CreateOneWithManager)
	router.PUT("/update", department.UpdateOne)
	router.PUT("/update/manager", department.UpdateManager)
	router.PUT("/update/status", department.UpdateStatus)
	router.DELETE("/delete/_id", department.DeleteOne)
	router.PATCH("/delete/_id", department.DeleteSoftOne)
}
