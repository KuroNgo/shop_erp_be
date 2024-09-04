package user_controller

//func (u *UserController) ForgetPasswordInUser(ctx *gin.Context) {
//	var forgetInput userdomain.ForgetPassword
//	if err := ctx.ShouldBindJSON(&forgetInput); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{
//			"status":  "error",
//			"message": err.Error()},
//		)
//		return
//	}
//
//	// check email exist
//	user, err := u.UserUseCase.GetByEmail(ctx, forgetInput.Email)
//	if err != nil || user.User == nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{
//			"error":  err.Error(),
//			"status": "error",
//		})
//		return
//	}
//
//	if user.User.Provider != "fe-it" {
//		ctx.JSON(http.StatusBadRequest, gin.H{
//			"error":   "Unsupported provider",
//			"message": fmt.Sprintf("Sorry, the %s you provided is not supported.", user.Provider),
//			"status":  "error",
//		})
//		return
//	}
//
//	// send mail
//	var code string
//	for {
//		code = randstr.Dec(6)
//		if u.UserUseCase.UniqueVerificationCode(ctx, code) {
//			break
//		}
//	}
//
//	updUser := userdomain.Verify{
//		Verified: true,
//	}
//
//	// Update User in Database
//	err = u.UserUseCase.UpdateVerify(ctx, user.User.ID.Hex(), updUser)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{
//			"status":  "error",
//			"message": err.Error()},
//		)
//		return
//	}
//
//	emailData := mail.EmailData{
//		Code:      code,
//		FirstName: user.User.Username,
//		Subject:   "Khôi phục mật khẩu",
//	}
//
//	err = mail.SendEmail(&emailData, user.User.Email, "forget_password.html")
//	if err != nil {
//		ctx.JSON(http.StatusBadGateway, gin.H{
//			"status":  "success",
//			"message": "There was an error sending email",
//		})
//		return
//	}
//
//	// Trả về phản hồi thành công
//	ctx.JSON(http.StatusOK, gin.H{
//		"status":  "success",
//		"message": "We sent an email with a verification code to your email",
//	})
//}
//
//func (u *UserController) VerificationCodeForChangePassword(ctx *gin.Context) {
//	var verificationCode userdomain.VerificationCode
//	if err := ctx.ShouldBindJSON(&verificationCode); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{
//			"status":  "error",
//			"message": err.Error(),
//		})
//		return
//	}
//
//	user, err := u.UserUseCase.GetByVerificationCode(ctx, verificationCode.VerificationCode)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{
//			"status":  "error",
//			"message": err.Error(),
//		})
//		return
//	}
//
//	res := u.UserUseCase.CheckVerify(ctx, verificationCode.VerificationCode)
//	if !res {
//		ctx.JSON(http.StatusNotModified, gin.H{
//			"status":  "error",
//			"message": "Verification code check failed",
//		})
//		return
//	}
//
//	updUser := userdomain.Verify{
//		Verified: true,
//	}
//
//	// Update User in Database
//	if err = u.UserUseCase.UpdateVerifyForChangePassword(ctx, user.User.ID.Hex(), updUser); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{
//			"status":  "error",
//			"message": err.Error(),
//		})
//		return
//	}
//
//	// Set cookie
//	ctx.SetCookie("verification_code", verificationCode.VerificationCode, u.Database.AccessTokenMaxAge*60, "/", "localhost", false, false)
//	ctx.SetSameSite(http.SameSiteStrictMode)
//
//	// Trả về phản hồi thành công
//	ctx.JSON(http.StatusOK, gin.H{
//		"status": "success",
//	})
//}
//
//func (u *UserController) ChangePassword(ctx *gin.Context) {
//	cookie, err := ctx.Cookie("verification_code")
//	if err != nil {
//		ctx.JSON(http.StatusForbidden, gin.H{
//			"status":  "fail",
//			"message": "Verification code is missing!",
//		})
//		return
//	}
//
//	var changePasswordInput userdomain.ChangePassword
//	if err := ctx.ShouldBindJSON(&changePasswordInput); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{
//			"status":  "error",
//			"message": err.Error()},
//		)
//		return
//	}
//
//	if changePasswordInput.Password != changePasswordInput.PasswordCompare {
//		ctx.JSON(http.StatusBadRequest, gin.H{
//			"status":  "error",
//			"message": "The passwords provided do not match.",
//		})
//		return
//	}
//
//	// Đối với change password, không clear giá trị verification Code ở phía client và backend
//	user, err := u.UserUseCase.GetByVerificationCode(ctx, cookie)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{
//			"status":  "error",
//			"message": "verification code do not match"},
//		)
//		return
//	}
//
//	changePasswordInput.Password, err = password.HashPassword(changePasswordInput.Password)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{
//			"status":  "error",
//			"message": err.Error()},
//		)
//		return
//	}
//
//	updateUser := &userdomain.User{
//		ID:           user.User.ID,
//		PasswordHash: changePasswordInput.Password,
//	}
//
//	err = u.UserUseCase.UpdatePassword(ctx,, updateUser)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{
//			"status":  "error",
//			"message": err.Error()},
//		)
//		return
//	}
//
//	ctx.SetCookie("verification_code", "", -1, "/", "localhost", false, false)
//
//	// Trả về phản hồi thành công
//	ctx.JSON(http.StatusOK, gin.H{
//		"status": "success",
//	})
//}
