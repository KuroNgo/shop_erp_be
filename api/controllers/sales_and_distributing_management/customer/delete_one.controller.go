package customer_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *CustomerController) DeleteOne(ctx *gin.Context) {
	_id := ctx.Param("_id")

	if err := c.CustomerUseCase.DeleteOne(ctx, _id); err != nil {
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
