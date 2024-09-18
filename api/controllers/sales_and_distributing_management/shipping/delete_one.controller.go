package shipping_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *ShippingController) DeleteOne(ctx *gin.Context) {
	_id := ctx.Param("_id")

	if err := s.ShippingUseCase.DeleteOne(ctx, _id); err != nil {
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
