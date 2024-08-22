package role_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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
