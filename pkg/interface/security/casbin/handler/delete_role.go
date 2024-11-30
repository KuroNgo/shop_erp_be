package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_erp_mono/pkg/interface/security/casbin/principle"
)

// DeleteRoleForUser godoc
// @Summary Delete role for a user
// @Description Remove a specific role from a user
// @Tags Casbin
// @Accept json
// @Produce json
// @Param data body UserRole true "User ID and Role data"
// @Router /api/v1/casbin/delete/user [delete]
func DeleteRoleForUser(ctx *gin.Context) {
	var data UserRole
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "can not get data",
		})
		return
	}

	for _, id := range data.UserID {
		_, err := principle.Rbac.RemoveGroupingPolicy(id, data.Role)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "can not delete role for user",
			})
		}
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
	})
}

// DeleteRole godoc
// @Summary Delete a role
// @Description Delete a role from the system
// @Tags Casbin
// @Accept json
// @Produce json
// @Param data body Role true "Role data"
// @Router /api/v1/casbin/delete [delete]
func DeleteRole(ctx *gin.Context) {
	var data Role

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "can not get data",
		})
		return
	}

	ok, err := principle.Rbac.DeleteRole(data.Role)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "fail to get: " + err.Error(),
		})
		return
	}

	// nếu không có role thì in ra
	if !ok {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{
				"message": "do not have role: " + data.Role,
			})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "success delete role: " + data.Role,
	})
}

// DeleteRoleForAPI godoc
// @Summary Delete role permissions for an API
// @Description Remove role-based permissions for an API and method
// @Tags Casbin
// @Accept json
// @Produce json
// @Param data body APIRole true "API and Role data"
// @Router /api/v1/casbin/delete/api/role [delete]
func DeleteRoleForAPI(ctx *gin.Context) {
	var data APIRole
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "can not get data",
		})
		return
	}

	allAction, err := principle.Rbac.GetAllActions()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "fail to get: " + err.Error(),
		})
		return
	}

	for _, role := range data.Role {
		for _, action := range allAction {
			_, err = principle.Rbac.RemovePolicy(role, data.API, action)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"status":  "error",
					"message": "can not delete role for user",
				})
			}
		}
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
	})
}

// DeleteAPIForRole godoc
// @Summary Delete API for a role
// @Description Remove APIs with specific methods for a role
// @Tags Casbin
// @Accept json
// @Produce json
// @Param data body RoleAPI true "Role and API data"
// @Router /api/v1/casbin/delete/role/api [delete]
func DeleteAPIForRole(ctx *gin.Context) {
	var data RoleAPI

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "can not get data",
		})
		return
	}

	allAction, err := principle.Rbac.GetAllActions()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "fail to get: " + err.Error(),
		})
		return
	}

	for _, api := range data.API {
		for _, action := range allAction {
			principle.Rbac.RemovePolicy(data.Role, api, action)
		}
	}

	// nếu không còn endpoint nào thì thêm http://localhost:8080
	filteredPolicy, err := principle.Rbac.GetFilteredPolicy(0, data.Role)
	if (len(filteredPolicy)) == 0 {
		_, err := principle.Rbac.AddPolicy(data.Role, "http://localhost:8080", "GET")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "can not delete role for user",
			})
		}
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
	})
}
