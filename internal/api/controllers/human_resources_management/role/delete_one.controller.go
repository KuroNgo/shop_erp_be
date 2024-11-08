package role_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteOne delete the role's information
// @Summary Delete Role Information
// @Description Deletes the role's information
// @Tags Role
// @Accept json
// @Produce json
// @Param _id path string true "Role ID"
// @Router /api/v1/roles/delete [delete]
// @Security CookieAuth
func (r *RoleController) DeleteOne(ctx *gin.Context) {
	id := ctx.Query("_id")

	err := r.RoleUseCase.DeleteOne(ctx, id)
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
