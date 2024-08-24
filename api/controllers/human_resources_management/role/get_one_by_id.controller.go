package role_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetOneRoleByID retrieves the role's information
// @Summary Get Role Information By Name
// @Description Retrieves the role's information name
// @Tags Role
// @Accept  json
// @Produce  json
// @Router /api/v1/roles/get/one/_id [get]
// @Security CookieAuth
func (r *RoleController) GetOneRoleByID(ctx *gin.Context) {
	id := ctx.Param("_id")

	data, err := r.RoleUseCase.GetByIDRole(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   data,
	})
}
