package user_controller

// LogoutUser
// @Summary Logout user
// @Description Logout the current user
// @Tags User
// @Accept  json
// @Produce  json
// @Router /api/v1/users/get/logout [get]
//func (u *UserController) LogoutUser(ctx *gin.Context) {
//	currentUser, exists := ctx.Get("currentUser")
//	if !exists {
//		ctx.JSON(http.StatusUnauthorized, gin.H{
//			"status":  "fail",
//			"message": "You are not logged in!",
//		})
//		return
//	}
//
//	user, err := u.UserUseCase.GetByID(ctx, fmt.Sprint(currentUser))
//	if err != nil || user == nil {
//		ctx.JSON(http.StatusUnauthorized, gin.H{
//			"status":  "Unauthorized",
//			"message": "You are not authorized to perform this action!",
//		})
//		return
//	}
//
//	ctx.SetCookie("access_token", "", -1, "/", "localhost", false, true)
//	ctx.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
//	ctx.SetCookie("logged_in", "", -1, "/", "localhost", false, false)
//
//	ctx.JSON(http.StatusOK, gin.H{
//		"status": "success",
//	})
//}
