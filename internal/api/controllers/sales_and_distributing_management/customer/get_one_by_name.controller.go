package customer_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetOneByName godoc
// @Summary Get a customer by name
// @Description Retrieve a customer from the system using their name
// @Tags Customers
// @Accept json
// @Produce json
// @Param name path string true "Customer Name"
// @Router /api/v1/customers/get/name [get]
func (c *CustomerController) GetOneByName(ctx *gin.Context) {
	name := ctx.Query("name")

	data, err := c.CustomerUseCase.GetOneByName(ctx, name)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
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
