package user_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhpk/randstr"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	userdomain "shop_erp_mono/domain/human_resource_management/user"
	imagescloudinary "shop_erp_mono/pkg/cloudinary/utils/images"
	"shop_erp_mono/pkg/helper"
	"shop_erp_mono/pkg/mail"
	"shop_erp_mono/pkg/password"
	"time"
)

// SignUp Create a new user
// @Summary Register user
// @Description Create new user
// @Tags User
// @Accept json
// @Produce json
// @Param RegisterUserRequestDto body user_domain.User true "User data"
// @Success 201 {object} map[string]interface{} "status: success, message: created a new user"
// @Security ApiKeyAuth
// @Router /api/v1/users/register [post]
func (u *UserController) SignUp(ctx *gin.Context) {
	emailForm := ctx.Request.FormValue("email")
	passwordForm := ctx.Request.FormValue("password")
	fullName := ctx.Request.FormValue("fullName")
	avatarUrlForm := ctx.Request.FormValue("avatarUrl")
	phoneForm := ctx.Request.FormValue("phone")

	if !helper.EmailValid(emailForm) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Email Invalid !",
		})
		return
	}

	// Bên phía client sẽ phải so sánh password thêm một lần nữa đã đúng chưa
	if !helper.PasswordStrong(passwordForm) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"message": "Password must have at least 8 characters, " +
				"including uppercase letters, lowercase letters and numbers!",
		})
		return
	}

	// Băm mật khẩu
	hashedPassword, err := password.HashPassword(passwordForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error()},
		)
		return
	}

	passwordForm = hashedPassword
	passwordForm = password.Sanitize(passwordForm)
	emailForm = password.Sanitize(emailForm)
	file, err := ctx.FormFile("file")
	if err != nil {
		newUser := &userdomain.User{
			ID:           primitive.NewObjectID(),
			Username:     fullName,
			AvatarURL:    avatarUrlForm,
			Email:        emailForm,
			PasswordHash: hashedPassword,
			Verified:     false,
			Provider:     "fe-it",
			Role:         "user",
			Phone:        phoneForm,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		// thực hiện đăng ký người dùng
		err = u.UserUseCase.Create(ctx, newUser)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": err.Error()},
			)
			return
		}

		var code string
		for {
			code = randstr.Dec(6)
			if u.UserUseCase.UniqueVerificationCode(ctx, code) {
				break
			}
		}

		updUser := userdomain.User{
			ID:               newUser.ID,
			VerificationCode: code,
			Verified:         false,
			UpdatedAt:        time.Now(),
		}

		// Update User in Database
		_, err = u.UserUseCase.UpdateVerify(ctx, &updUser)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": err.Error()},
			)
			return
		}

		emailData := mail.EmailData{
			Code:      code,
			FirstName: newUser.Username,
			Subject:   "Your account verification code",
		}

		err = mail.SendEmail(&emailData, newUser.Email, "sign_in_first_time.html")
		if err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{
				"status":  "success",
				"message": "There was an error sending email",
			})
			return
		}

		// Trả về phản hồi thành công
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "We sent an email with a verification code to your email",
		})
		return
	}

	// Kiểm tra xem file có phải là hình ảnh không
	if !helper.IsImage(file.Filename) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid file format. Only images are allowed.",
		})
		return
	}

	// Mở file để đọc dữ liệu
	f, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error opening uploaded file",
		})
		return
	}
	defer f.Close()

	imageURL, err := imagescloudinary.UploadImageToCloudinary(f, file.Filename, u.Database.CloudinaryUploadFolderUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	newUser := userdomain.User{
		ID:           primitive.NewObjectID(),
		Username:     fullName,
		AvatarURL:    imageURL.ImageURL,
		AssetURL:     imageURL.AssetID,
		Email:        emailForm,
		PasswordHash: hashedPassword,
		Verified:     false,
		Provider:     "fe-it",
		Role:         "user",
		Phone:        phoneForm,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// thực hiện đăng ký người dùng
	err = u.UserUseCase.Create(ctx, &newUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error()},
		)
		return
	}

	// Trả về phản hồi thành công
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
	return
}

func (u *UserController) VerificationCode(ctx *gin.Context) {
	var verificationCode userdomain.VerificationCode
	if err := ctx.ShouldBindJSON(&verificationCode); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error()},
		)
		return
	}

	user, err := u.UserUseCase.GetByVerificationCode(ctx, verificationCode.VerificationCode)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error()},
		)
		return
	}

	res := u.UserUseCase.CheckVerify(ctx, verificationCode.VerificationCode)
	if res != true {
		ctx.JSON(http.StatusNotModified, gin.H{
			"status":  "error",
			"message": err.Error()},
		)
		return
	}

	updUser := userdomain.User{
		ID:        user.ID,
		Verified:  true,
		UpdatedAt: time.Now(),
	}

	// Update User in Database
	_, err = u.UserUseCase.UpdateVerify(ctx, &updUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error()},
		)
		return
	}

	// Trả về phản hồi thành công
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}