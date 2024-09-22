package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_erp_mono/pkg/casbin/principle"
)

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
