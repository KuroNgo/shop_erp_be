package role_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByTitle retrieves the role's information
// @Summary Get Role Information By Name
// @Description Retrieves the role's information name
// @Tags Role
// @Accept  json
// @Produce  json
// @Param title path string true "Employee ID"
// @Router /api/v1/roles/get/title [get]
// @Security CookieAuth
func (r *RoleController) GetByTitle(ctx *gin.Context) {
	title := ctx.Query("title")

	data, err := r.RoleUseCase.GetByTitle(ctx, title)
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
