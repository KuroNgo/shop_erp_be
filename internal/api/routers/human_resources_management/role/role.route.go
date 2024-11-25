package role_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	role_controller "shop_erp_mono/internal/api/controllers/human_resources_management/role"
	"shop_erp_mono/internal/config"
	employees_domain "shop_erp_mono/internal/domain/human_resource_management/employees"
	roledomain "shop_erp_mono/internal/domain/human_resource_management/role"
	userdomain "shop_erp_mono/internal/domain/human_resource_management/user"
	employee_repository "shop_erp_mono/internal/repository/human_resource_management/employee/repository"
	role_repository "shop_erp_mono/internal/repository/human_resource_management/role/repository"
	user_repository "shop_erp_mono/internal/repository/human_resource_management/user/repository"
	role_usecase "shop_erp_mono/internal/usecase/human_resource_management/role/usecase"
	"time"
)

func RoleRouter(env *config.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup, cacheTTL time.Duration) {
	ro := role_repository.NewRoleRepository(db, roledomain.CollectionRole)
	us := user_repository.NewUserRepository(db, userdomain.CollectionUser)
	em := employee_repository.NewEmployeeRepository(db, employees_domain.CollectionEmployee)
	role := &role_controller.RoleController{
		RoleUseCase: role_usecase.NewRoleUseCase(timeout, ro, us, em, cacheTTL),
		Database:    env,
	}

	router := group.Group("/roles")
	router.GET("/get/_id", role.GetByID)
	router.GET("/get/title", role.GetByTitle)
	router.GET("/get/all", role.GetAll)
	router.POST("/create", role.CreateOne)
	router.PUT("/update", role.UpdateOne)
	router.DELETE("/delete", role.DeleteOne)
}
