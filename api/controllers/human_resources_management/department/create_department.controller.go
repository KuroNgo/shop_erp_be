package department_controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	departmentsdomain "shop_erp_mono/domain/human_resource_management/departments"
	"time"
)

// CreateOneDepartment create the department's information
// @Summary Create Department Information
// @Description Create the department's information
// @Tags Department
// @Accept json
// @Produce json
// @@Success 200 {object} "status: success, message:update department success"
// @Failure 400 {object} map[string]interface{} "status: fail, message: detailed error message"
// @Failure 401 {object} map[string]interface{} "status: fail, message: You are not logged in!"
// @Router /api/v1/departments/create [post]
// @Security CookieAuth
func (d *DepartmentController) CreateOneDepartment(ctx *gin.Context) {
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

	err := d.DepartmentUseCase.CreateOne(ctx, &department)
	if err != nil {
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