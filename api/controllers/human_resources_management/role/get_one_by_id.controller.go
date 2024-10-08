package role_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetByID retrieves the role's information
// @Summary Get Role Information By Name
// @Description Retrieves the role's information name
// @Tags Role
// @Accept  json
// @Produce  json
// @Param _id path string true "Employee ID"
// @Router /api/v1/roles/get/_id [get]
// @Security CookieAuth
func (r *RoleController) GetByID(ctx *gin.Context) {
	id := ctx.Query("_id")

	data, err := r.RoleUseCase.GetByID(ctx, id)
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
