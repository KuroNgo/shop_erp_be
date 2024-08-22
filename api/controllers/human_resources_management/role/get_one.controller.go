package role_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *RoleController) GetOneRoleByTitle(ctx *gin.Context) {
	title := ctx.Param("title")

	data, err := r.RoleUseCase.GetByTitleRole(ctx, title)
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
