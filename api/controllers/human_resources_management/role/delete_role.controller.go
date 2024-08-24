package role_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteOneRole delete the role's information
// @Summary Delete Role Information
// @Description Deletes the role's information
// @Tags Role
// @Accept json
// @Produce json
// @Router /api/v1/roles/delete [delete]
// @Security CookieAuth
func (r *RoleController) DeleteOneRole(ctx *gin.Context) {
	id := ctx.Param("_id")

	err := r.RoleUseCase.DeleteOneRole(ctx, id)
	if err != nil {
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
