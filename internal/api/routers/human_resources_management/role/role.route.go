package role_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	role_controller "shop_erp_mono/internal/api/controllers/human_resources_management/role"
	"shop_erp_mono/internal/config"
	roledomain "shop_erp_mono/internal/domain/human_resource_management/role"
	role_repository "shop_erp_mono/internal/repository/human_resource_management/role/repository"
	role_usecase "shop_erp_mono/internal/usecase/human_resource_management/role/usecase"
	"time"
)

func RoleRouter(env *config.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup, cacheTTL time.Duration) {
	ro := role_repository.NewRoleRepository(db, roledomain.CollectionRole)
	role := &role_controller.RoleController{
		RoleUseCase: role_usecase.NewRoleUseCase(timeout, ro, cacheTTL),
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
