package casbin_router

import (
	"github.com/gin-gonic/gin"
	"shop_erp_mono/pkg/casbin/handler"
	"shop_erp_mono/pkg/casbin/principle"
)

func CasbinRouter(group *gin.RouterGroup) {
	r := principle.SetUp()
	cbGroup := group.Group("/casbin")
	cbGroup.POST("/add", handler.AddRole)
	cbGroup.POST("/add/user", handler.AddRoleForUser)
	cbGroup.POST("/add/role/api", handler.AddRoleForAPI)
	cbGroup.POST("/add/api/role", handler.AddAPIForRole)
	cbGroup.DELETE("/delete", handler.DeleteRole)
	cbGroup.DELETE("/delete/user", handler.DeleteRoleForUser)
	cbGroup.DELETE("/delete/role/api", handler.AddRoleForAPI)
	cbGroup.DELETE("/delete/api/role", handler.AddAPIForRole)
	err := r.SavePolicy()
	if err != nil {
		return
	}
}
