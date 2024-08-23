package role_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAllRole retrieves the role's information
// @Summary Get Role Information
// @Description Retrieves the role's information
// @Tags Role
// @Accept  json
// @Produce  json
// @Success 200 {object} role_domain.Role
// @Failure 400 {object} map[string]interface{} "status: fail, message: detailed error message"
// @Failure 401 {object} map[string]interface{} "status: fail, message: You are not logged in!"
// @Router /api/v1/all/roles [get]
// @Security CookieAuth
func (r *RoleController) GetAllRole(ctx *gin.Context) {
	data, err := r.RoleUseCase.GetAllRole(ctx)
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
