package customer_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *CustomerController) GetAll(ctx *gin.Context) {
	data, err := c.CustomerUseCase.GetAll(ctx)
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
