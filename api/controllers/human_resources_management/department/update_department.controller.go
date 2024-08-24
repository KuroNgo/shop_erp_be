package department_controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	departmentsdomain "shop_erp_mono/domain/human_resource_management/departments"
	"time"
)

// UpdateOneDepartment updates the department's information
// @Summary Update Department Information
// @Description Updates the department's information
// @Tags Department
// @Accept json
// @Produce json
// @Router /api/v1/departments/update [patch]
// @Security CookieAuth
func (d *DepartmentController) UpdateOneDepartment(ctx *gin.Context) {
	var input departmentsdomain.Input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	department := departmentsdomain.Department{
		ID:          primitive.NewObjectID(),
		Name:        input.Name,
		Description: input.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := d.DepartmentUseCase.UpdateOne(ctx, &department); err != nil {
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
