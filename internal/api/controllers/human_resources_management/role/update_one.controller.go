package role_controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	roledomain "shop_erp_mono/internal/domain/human_resource_management/role"
	"shop_erp_mono/pkg/shared/constant"
)

// UpdateOne updates the role's information
// @Summary Update Role Information
// @Description Updates the role's information
// @Tags Role
// @Accept json
// @Produce json
// @Router /api/v1/roles/update [put]
// @Security CookieAuth
func (r *RoleController) UpdateOne(ctx *gin.Context) {
	currentUser, exists := ctx.Get("currentUser")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "fail",
			"message": constant.MsgAPIUnauthorized,
		})
		return
	}

	var role roledomain.Input
	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	roleID := ctx.Query("_id")

	if err := r.RoleUseCase.UpdateOne(ctx, roleID, &role, fmt.Sprintf("%s", currentUser)); err != nil {
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
