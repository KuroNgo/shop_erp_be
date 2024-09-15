package stock_adjustment_controller

//func (s *StockAdjustmentController) GetByUserID(ctx *gin.Context) {
//	useriD := ctx.Param("user_id")
//
//	data, err := s.StockAdjustmentUseCase.Get(ctx, useriD)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{
//			"status":  "error",
//			"message": err.Error(),
//		})
//		return
//	}
//
//	ctx.JSON(http.StatusOK, gin.H{
//		"status": "success",
//		"data":   data,
//	})
//}
