package role_controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	roledomain "shop_erp_mono/domain/human_resource_management/role"
	"time"
)

func (r *RoleController) CreateOneRole(ctx *gin.Context) {
	var role roledomain.Input
	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	roleData := &roledomain.Role{
		ID:          primitive.NewObjectID(),
		Title:       role.Title,
		Description: role.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := r.RoleUseCase.CreateOneRole(ctx, roleData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
