package role_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	role_controller "shop_erp_mono/api/controllers/human_resources_management/role"
	"shop_erp_mono/bootstrap"
	roledomain "shop_erp_mono/domain/human_resource_management/role"
	role_repository "shop_erp_mono/repository/human_resource_management/role/repository"
	role_usecase "shop_erp_mono/usecase/human_resource_management/role/usecase"
	"time"
)

func RoleRouter(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	ro := role_repository.NewRoleRepository(db, roledomain.CollectionRole)
	role := &role_controller.RoleController{
		RoleUseCase: role_usecase.NewRoleUseCase(timeout, ro),
		Database:    env,
	}

	router := group.Group("/roles")
	router.GET("/_id", role.GetOneRoleByID)
	router.GET("/title", role.GetOneRoleByTitle)
	router.GET("/all", role.GetAllRole)
	router.POST("/create", role.CreateOneRole)
	router.PUT("/update", role.UpdateRole)
	router.DELETE("/delete", role.DeleteOneRole)
}
