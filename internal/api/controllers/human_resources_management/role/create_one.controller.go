package role_controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	roledomain "shop_erp_mono/internal/domain/human_resource_management/role"
	"shop_erp_mono/pkg/shared/constant"
)

// CreateOne Create a new role
// @Summary Create role
// @Description Create new role
// @Tags Role
// @Accept json
// @Produce json
// @Param Role body role_domain.Input true "Role data"
// @Security ApiKeyAuth
// @Router /api/v1/roles/create [post]
func (r *RoleController) CreateOne(ctx *gin.Context) {
	currentUser, exists := ctx.Get("currentUser")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "fail",
			"message": constant.MsgAPIUnauthorized,
		})
		return
	}

	var input roledomain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	if err := r.RoleUseCase.CreateOne(ctx, &input, fmt.Sprintf("%s", currentUser)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": constant.MsgAPIBadRequest,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
