package casbin_router

import (
	"github.com/gin-gonic/gin"
	"shop_erp_mono/internal/config"
	"shop_erp_mono/pkg/interface/security/casbin/handler"
	"shop_erp_mono/pkg/interface/security/casbin/principle"
)

func CasbinRouter(group *gin.RouterGroup, env *config.Database) {
	cbGroup := group.Group("/casbin")
	cbGroup.POST("/add", handler.AddRole)
	cbGroup.POST("/add/user", handler.AddRoleForUser)
	cbGroup.POST("/add/role/api", handler.AddRoleForAPI)
	cbGroup.POST("/add/api/role", handler.AddAPIForRole)
	cbGroup.DELETE("/delete", handler.DeleteRole)
	cbGroup.DELETE("/delete/user", handler.DeleteRoleForUser)
	cbGroup.DELETE("/delete/role/api", handler.DeleteAPIForRole)
	cbGroup.DELETE("/delete/api/role", handler.DeleteRoleForAPI)
	r := principle.SetUp(env)
	err := r.SavePolicy()
	if err != nil {
		return
	}
}
