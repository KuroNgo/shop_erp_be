package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_erp_mono/pkg/casbin/principle"
)

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
