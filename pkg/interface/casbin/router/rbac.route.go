package casbin_router

import (
	"github.com/gin-gonic/gin"
	"shop_erp_mono/internal/config"
	handler2 "shop_erp_mono/pkg/interface/casbin/handler"
	"shop_erp_mono/pkg/interface/casbin/principle"
)

func CasbinRouter(group *gin.RouterGroup, env *config.Database) {
	cbGroup := group.Group("/casbin")
	cbGroup.POST("/add", handler2.AddRole)
	cbGroup.POST("/add/user", handler2.AddRoleForUser)
	cbGroup.POST("/add/role/api", handler2.AddRoleForAPI)
	cbGroup.POST("/add/api/role", handler2.AddAPIForRole)
	cbGroup.DELETE("/delete", handler2.DeleteRole)
	cbGroup.DELETE("/delete/user", handler2.DeleteRoleForUser)
	cbGroup.DELETE("/delete/role/api", handler2.DeleteAPIForRole)
	cbGroup.DELETE("/delete/api/role", handler2.DeleteRoleForAPI)
	r := principle.SetUp(env)
	err := r.SavePolicy()
	if err != nil {
		return
	}
}
