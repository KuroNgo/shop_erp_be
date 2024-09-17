package customer_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *CustomerController) GetOneByName(ctx *gin.Context) {
	name := ctx.Param("name")

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
