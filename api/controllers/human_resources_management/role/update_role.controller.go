package role_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	roledomain "shop_erp_mono/domain/human_resource_management/role"
	"time"
)

func (r *RoleController) UpdateRole(ctx *gin.Context) {
	var role roledomain.Input
	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	roleData := &roledomain.Role{
		Title:       role.Title,
		Description: role.Description,
		UpdatedAt:   time.Now(),
	}

	if err := r.RoleUseCase.UpdateOneRole(ctx, roleData); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
