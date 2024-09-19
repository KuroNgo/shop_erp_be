package role_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	roledomain "shop_erp_mono/domain/human_resource_management/role"
)

// UpdateRole updates the role's information
// @Summary Update Role Information
// @Description Updates the role's information
// @Tags Role
// @Accept json
// @Produce json
// @Router /api/v1/roles/update [put]
// @Security CookieAuth
func (r *RoleController) UpdateRole(ctx *gin.Context) {
	var role roledomain.Input
	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	roleID := ctx.Query("_id")

	if err := r.RoleUseCase.UpdateOneRole(ctx, roleID, &role); err != nil {
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
