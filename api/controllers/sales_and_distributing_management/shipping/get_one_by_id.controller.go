package shipping_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *ShippingController) GetByID(ctx *gin.Context) {
	_id := ctx.Param("id")

	data, err := s.ShippingUseCase.GetByID(ctx, _id)
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
