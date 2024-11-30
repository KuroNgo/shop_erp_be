package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_erp_mono/pkg/interface/security/casbin/principle"
)

// AddRole godoc
// @Summary Add a role to the system
// @Description Add a new role with API and method to the system
// @Tags Casbin
// @Accept json
// @Produce json
// @Param data body RoleData true "Role data"
// @Router /api/v1/casbin/add/role [post]
func AddRole(ctx *gin.Context) {
	var data RoleData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, "can not get data")
		return
	}

	if data.API == nil {
		data.API = append(data.API, "http://localhost:8080")
	}

	// Add policy rules
	for _, api := range data.API {
		for _, method := range data.Method {
			// Assuming rbac is already initialized and AddPolicy is defined
			_, err := principle.Rbac.AddPolicy(data.Role, api, method)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "ERROR"})
				return
			}
		}
	}

	ctx.JSON(http.StatusCreated, "success added role: "+data.Role)
}

// AddRoleForUser godoc
// @Summary Add a role to a user
// @Description Assign a role to a user
// @Tags Casbin
// @Accept json
// @Produce json
// @Param data body UserRole true "User and Role data"
// @Router /api/v1/casbin/add/user [post]
func AddRoleForUser(ctx *gin.Context) {
	var data UserRole

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "can not get data",
		})
	}

	for _, id := range data.UserID {
		_, err := principle.Rbac.AddGroupingPolicy(id, data.Role)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "can not add role for this user",
			})
		}
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
	})
}

// AddRoleForAPI godoc
// @Summary Add role permissions for an API
// @Description Add role-based permissions for a specific API and method
// @Tags Casbin
// @Accept json
// @Produce json
// @Param data body APIData true "API data with Role and Methods"
// @Router /api/v1/casbin/add/role/api [post]
func AddRoleForAPI(ctx *gin.Context) {
	var data APIData

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "success",
			"message": "can not get data",
		})
	}

	for _, role := range data.Role {
		for _, method := range data.Method {
			_, err := principle.Rbac.AddPolicy(role, data.API, method)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"status":  "fail",
					"message": "can not add role for this user",
				})
			}
		}
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
	})
}

// AddAPIForRole godoc
// @Summary Assign API permissions to a role
// @Description Add APIs with specific methods to a role
// @Tags Casbin
// @Accept json
// @Produce json
// @Param data body RoleData true "Role and API data"
// @Router /api/v1/casbin/add/api/role [post]
func AddAPIForRole(ctx *gin.Context) {
	var data RoleData

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "success",
			"message": "can not get data",
		})
	}

	// nếu chưa có method nào thì thêm method mặc định GET
	if data.Method == nil {
		data.Method = append(data.Method, "GET")
	}

	for _, api := range data.API {
		for _, method := range data.Method {
			_, err := principle.Rbac.AddPolicy(data.Role, api, method)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"status":  "fail",
					"message": "can not add role for this user",
				})
			}
		}
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
	})
}
