package user_controller

// UpdateUser updates the user's information
// @Summary Update User Information
// @Description Updates the user's first name, last name, and username
// @Tags User
// @Accept json
// @Produce json
// @Router /api/v1/users/update [put]
// @Security CookieAuth
//func (u *UserController) UpdateUser(ctx *gin.Context) {
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
//	fullName := ctx.Request.FormValue("full_name")
//
//	file, err := ctx.FormFile("file")
//	if err != nil {
//		userResponse := userdomain.UpdateUser{
//			ID:        user.ID,
//			Username:  fullName,
//			UpdatedAt: time.Now(),
//		}
//
//		err = u.UserUseCase.Update(ctx, &userResponse)
//		if err != nil {
//			ctx.JSON(http.StatusInternalServerError, gin.H{
//				"status":  "error",
//				"message": err.Error(),
//			})
//			return
//		}
//
//		ctx.JSON(http.StatusOK, gin.H{
//			"status": "success",
//		})
//		return
//	}
//
//	if !helper.IsImage(file.Filename) {
//		ctx.JSON(http.StatusBadRequest, gin.H{
//			"error": "Invalid file format. Only images are allowed.",
//		})
//		return
//	}
//
//	// Mở file để đọc dữ liệu
//	f, err := file.Open()
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{
//			"error": "Error opening uploaded file",
//		})
//		return
//	}
//	defer func(f multipart.File) {
//		err = f.Close()
//		if err != nil {
//			return
//		}
//	}(f)
//
//	resultString, err := json.Marshal(user)
//	userResponse := userdomain.UpdateUser{
//		ID:        user.ID,
//		Username:  fullName,
//		UpdatedAt: time.Now(),
//	}
//
//	err = u.UserUseCase.Update(ctx, &userResponse)
//	if err != nil {
//		ctx.JSON(http.StatusForbidden, gin.H{
//			"status":  "fail",
//			"message": string(resultString) + "the user belonging to this token no logger exists",
//		})
//		return
//	}
//
//	ctx.JSON(http.StatusOK, gin.H{
//		"status":  "success",
//		"message": "Updated user",
//	})
//
//}
