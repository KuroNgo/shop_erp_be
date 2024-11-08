package role_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAll retrieves the role's information
// @Summary Get Role Information
// @Description Retrieves the role's information
// @Tags Role
// @Accept  json
// @Produce  json
// @Router /api/v1/roles/get/all [get]
// @Security CookieAuth
func (r *RoleController) GetAll(ctx *gin.Context) {
	data, err := r.RoleUseCase.GetAll(ctx)
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
