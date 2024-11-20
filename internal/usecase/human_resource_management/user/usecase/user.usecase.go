package user_usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/thanhpk/randstr"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongo_driven "go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"mime/multipart"
	"shop_erp_mono/internal/config"
	userdomain "shop_erp_mono/internal/domain/human_resource_management/user"
	"shop_erp_mono/internal/usecase/human_resource_management/user/validate"
	"shop_erp_mono/pkg/interface/cloudinary/utils/images"
	"shop_erp_mono/pkg/interface/oauth2/google"
	helper2 "shop_erp_mono/pkg/shared/helper"
	"shop_erp_mono/pkg/shared/mail/handles"
	"shop_erp_mono/pkg/shared/password"
	"shop_erp_mono/pkg/shared/token"
	"time"
)

type userUseCase struct {
	database       *config.Database
	contextTimeout time.Duration
	userRepository userdomain.IUserRepository
	client         *mongo_driven.Client
}

func NewUserUseCase(database *config.Database, contextTimeout time.Duration, userRepository userdomain.IUserRepository,
	client *mongo_driven.Client) userdomain.IUserUseCase {
	return &userUseCase{database: database, contextTimeout: contextTimeout, userRepository: userRepository, client: client}
}

// SignUp Create a new user
func (u *userUseCase) SignUp(ctx context.Context, file *multipart.FileHeader, input *userdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	session, err := u.client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	callback := func(sessionCtx mongo_driven.SessionContext) (interface{}, error) {
		if err = validate.User(input); err != nil {
			return nil, err
		}

		// Bên phía client sẽ phải so sánh password thêm một lần nữa đã đúng chưa
		if !helper2.PasswordStrong(input.PasswordHash) {
			return nil, errors.New("password must have at least 8 characters including uppercase letters, lowercase letters and numbers")
		}

		// Băm mật khẩu
		hashedPassword, err := password.HashPassword(input.PasswordHash)
		if err != nil {
			return nil, err
		}

		if file == nil {
			newUser := &userdomain.User{
				ID:           primitive.NewObjectID(),
				Username:     input.Username,
				Email:        input.Email,
				PasswordHash: hashedPassword,
				Verified:     false,
				Provider:     "inside",
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			}

			err = u.userRepository.CreateOne(sessionCtx, newUser)
			if err != nil {
				return nil, err
			}

			var code string
			code = randstr.Dec(6)

			updUser := userdomain.User{
				ID:               newUser.ID,
				VerificationCode: code,
				Verified:         false,
				UpdatedAt:        time.Now(),
			}

			// Update User in Database
			err = u.userRepository.UpdateVerificationCode(sessionCtx, &updUser)
			if err != nil {
				return nil, err
			}

			emailData := handles.EmailData{
				Code:     code,
				FullName: newUser.Username,
				Subject:  "Your account verification code: " + code,
			}

			err = handles.SendEmail(&emailData, newUser.Email, "user.sign_up.html")
			if err != nil {
				return nil, err
			}

			return nil, nil
		}

		// Kiểm tra xem file có phải là hình ảnh không
		if !helper2.IsImage(file.Filename) {
			return nil, err
		}

		f, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer f.Close()

		imageURL, err := images_cloudinary.UploadImageToCloudinary(f, file.Filename, u.database.CloudinaryUploadFolderUser)
		if err != nil {
			return nil, err
		}

		// Đảm bảo xóa ảnh trên Cloudinary nếu xảy ra lỗi sau khi tải lên thành công
		defer func() {
			if err != nil {
				_, _ = images_cloudinary.DeleteToCloudinary(imageURL.AssetID)
			}
		}()

		newUser := userdomain.User{
			ID:           primitive.NewObjectID(),
			Username:     input.Username,
			AvatarURL:    imageURL.ImageURL,
			AssetURL:     imageURL.AssetID,
			Email:        input.Email,
			PasswordHash: hashedPassword,
			Verified:     false,
			Provider:     "app",
			Role:         "user",
			Phone:        input.Phone,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		// thực hiện đăng ký người dùng
		err = u.userRepository.CreateOne(sessionCtx, &newUser)
		if err != nil {
			return nil, err
		}

		var code string
		code = randstr.Dec(6)

		updUser := userdomain.User{
			ID:               newUser.ID,
			VerificationCode: code,
			Verified:         false,
			UpdatedAt:        time.Now(),
		}

		// Update User in Database
		err = u.userRepository.UpdateVerificationCode(sessionCtx, &updUser)
		if err != nil {
			return nil, err
		}

		emailData := handles.EmailData{
			Code:     code,
			FullName: newUser.Username,
			Subject:  "Your account verification code: " + code,
		}

		err = handles.SendEmail(&emailData, newUser.Email, "user.sign_up.html")
		if err != nil {
			return nil, err
		}

		return nil, nil
	}

	// Run the transaction
	_, err = session.WithTransaction(ctx, callback)
	if err != nil {
		return err
	}

	return session.CommitTransaction(ctx)
}

// GetByVerificationCode authentication for Create a new user
func (u *userUseCase) GetByVerificationCode(ctx context.Context, verificationCode string) (userdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	user, err := u.userRepository.GetByVerificationCode(ctx, verificationCode)
	if err != nil {
		return userdomain.Output{}, err
	}

	updUser := userdomain.User{
		ID:        user.ID,
		Verified:  true,
		UpdatedAt: time.Now(),
	}

	// Update User in Database
	err = u.userRepository.UpdateVerificationCode(ctx, &updUser)
	if err != nil {
		return userdomain.Output{}, err
	}

	response := userdomain.Output{
		User: user,
	}

	return response, nil
}

func (u *userUseCase) LoginUser(ctx context.Context, signIn *userdomain.SignIn) (userdomain.OutputLogin, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	userData, err := u.userRepository.GetByEmail(ctx, signIn.Email)
	if err != nil || userData.Verified == false {
		return userdomain.OutputLogin{}, err
	}

	err = password.VerifyPassword(userData.PasswordHash, signIn.Password)
	if err != nil {
		return userdomain.OutputLogin{}, err
	}

	accessToken, err := token.CreateToken(u.database.AccessTokenExpiresIn, userData.ID, u.database.AccessTokenPrivateKey)
	if err != nil {
		return userdomain.OutputLogin{}, err
	}

	refreshToken, err := token.CreateToken(u.database.RefreshTokenExpiresIn, userData.ID, u.database.RefreshTokenPrivateKey)
	if err != nil {
		return userdomain.OutputLogin{}, err
	}

	response := userdomain.OutputLogin{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		IsLogged:     "1",
	}

	return response, nil
}

func (u *userUseCase) UpdateOne(ctx context.Context, userID string, input *userdomain.Input, file *multipart.FileHeader) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	idUser, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	userData, err := u.userRepository.GetByID(ctx, idUser)
	if err != nil {
		return err
	}

	if file == nil {
		user := userdomain.User{
			ID:        userData.ID,
			Username:  input.Username,
			UpdatedAt: time.Now(),
		}

		err = u.userRepository.UpdateOne(ctx, &user)
		if err != nil {
			return err
		}

		return nil
	}

	if !helper2.IsImage(file.Filename) {
		return err
	}

	f, err := file.Open()
	if err != nil {
		return err
	}
	defer func(f multipart.File) {
		err = f.Close()
		if err != nil {
			return
		}
	}(f)

	user := userdomain.User{
		ID:        idUser,
		Username:  input.Username,
		UpdatedAt: time.Now(),
	}

	err = u.userRepository.UpdateOne(ctx, &user)
	if err != nil {
		return err
	}

	return nil
}

func (u *userUseCase) UpdateVerify(ctx context.Context, id string, input *userdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	idUser, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	if err = validate.User(input); err != nil {
		return err
	}

	user := userdomain.User{
		ID:               idUser,
		VerificationCode: input.VerificationCode,
		Verified:         false,
		UpdatedAt:        time.Now(),
	}

	return u.userRepository.UpdateVerify(ctx, &user)
}

func (u *userUseCase) UpdatePassword(ctx context.Context, id string, input *userdomain.ChangePasswordInput) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	if input.Password != input.PasswordCompare {
		return errors.New("the passwords provided do not match")
	}

	user, err := u.userRepository.GetByVerificationCode(ctx, id)
	if err != nil {
		return err
	}

	input.Password, err = password.HashPassword(input.Password)
	if err != nil {
		return err
	}

	updateUser := &userdomain.User{
		ID:           user.ID,
		PasswordHash: input.Password,
		UpdatedAt:    time.Now(),
	}

	return u.userRepository.UpdatePassword(ctx, updateUser)
}

func (u *userUseCase) UpdateVerifyForChangePassword(ctx context.Context, verificationCode string) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	user, err := u.userRepository.GetByVerificationCode(ctx, verificationCode)
	if err != nil {
		return err
	}

	if user.Verified == false {
		return errors.New("verification code check failed")
	}

	updUser := userdomain.User{
		ID:       user.ID,
		Verified: true,
	}

	return u.userRepository.UpdateVerify(ctx, &updUser)
}

func (u *userUseCase) ForgetPassword(ctx context.Context, email string) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	session, err := u.client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	callback := func(sessionCtx mongo_driven.SessionContext) (interface{}, error) {
		user, err := u.userRepository.GetByEmail(sessionCtx, email)
		if err != nil {
			return nil, err
		}

		var code string
		code = randstr.Dec(6)

		updUser := &userdomain.User{
			ID:       user.ID,
			Verified: true,
		}

		// Update User in Database
		err = u.userRepository.UpdateVerify(sessionCtx, updUser)
		if err != nil {
			return nil, err
		}

		emailData := handles.EmailData{
			Code:     code,
			FullName: user.Username,
			Subject:  "Khôi phục mật khẩu",
		}

		err = handles.SendEmail(&emailData, user.Email, "user.forget_password.html")
		if err != nil {
			return nil, err
		}

		return nil, nil
	}

	// Run the transaction
	_, err = session.WithTransaction(ctx, callback)
	if err != nil {
		return err
	}

	return session.CommitTransaction(ctx)
}

func (u *userUseCase) LoginGoogle(ctx context.Context, code string) (*userdomain.Output, *userdomain.OutputLoginGoogle, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	googleOauthConfig := &oauth2.Config{
		ClientID:     u.database.GoogleClientID,
		ClientSecret: u.database.GoogleClientSecret,
		RedirectURL:  u.database.GoogleOAuthRedirectUrl,
		Scopes:       []string{"profile", "email"}, // Adjust scopes as needed
		Endpoint:     google.Endpoint,
	}

	tokenData, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Println("Error exchanging code: " + err.Error())
		return nil, nil, err
	}

	userInfo, err := google_oauth2.GetUserInfo(tokenData.AccessToken)
	if err != nil {
		fmt.Println("Error getting user info: " + err.Error())
		return nil, nil, err
	}

	// Giả sử userInfo là một map[string]interface{}
	email := userInfo["email"].(string)
	phone := userInfo["phone"].(string)
	fullName := userInfo["name"].(string)
	avatarURL := userInfo["picture"].(string)
	verifiedEmail := userInfo["verified_email"].(bool)

	user := &userdomain.User{
		ID:        primitive.NewObjectID(),
		Email:     email,
		Phone:     phone,
		Username:  fullName,
		AvatarURL: avatarURL,
		Provider:  "google",
		Verified:  verifiedEmail,
		Role:      "guess",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	updateUser, err := u.userRepository.UpsertOne(ctx, user)
	if err != nil {
		return nil, nil, err
	}

	signedToken, err := google_oauth2.SignJWT(userInfo)
	if err != nil {
		fmt.Println("Error signing token: " + err.Error())
		return nil, nil, err
	}

	accessToken, err := token.CreateToken(u.database.AccessTokenExpiresIn, updateUser.ID, u.database.AccessTokenPrivateKey)
	if err != nil {
		return nil, nil, err
	}

	refreshToken, err := token.CreateToken(u.database.RefreshTokenExpiresIn, updateUser.ID, u.database.RefreshTokenPrivateKey)
	if err != nil {
		return nil, nil, err
	}

	output := &userdomain.Output{
		User: *updateUser,
	}

	output2 := &userdomain.OutputLoginGoogle{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		IsLogged:     "1",
		SignedToken:  signedToken,
	}

	return output, output2, nil
}

func (u *userUseCase) UpdateImage(ctx context.Context, id string, file *multipart.FileHeader) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	idUser, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	session, err := u.client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	callback := func(sessionCtx mongo_driven.SessionContext) (interface{}, error) {
		if file == nil {
			return nil, errors.New("images not nil")
		}

		// Kiểm tra xem file có phải là hình ảnh không
		if !helper2.IsImage(file.Filename) {
			return nil, err
		}

		f, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer f.Close()

		imageURL, err := images_cloudinary.UploadImageToCloudinary(f, file.Filename, u.database.CloudinaryUploadFolderUser)
		if err != nil {
			return nil, err
		}

		// Đảm bảo xóa ảnh trên Cloudinary nếu xảy ra lỗi sau khi tải lên thành công
		defer func() {
			if err != nil {
				_, _ = images_cloudinary.DeleteToCloudinary(imageURL.AssetID)
			}
		}()

		user := userdomain.User{
			ID:        idUser,
			AvatarURL: imageURL.ImageURL,
			AssetURL:  imageURL.AssetID,
			UpdatedAt: time.Now(),
		}

		err = u.userRepository.UpdateImage(sessionCtx, &user)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}

	// Run the transaction
	_, err = session.WithTransaction(ctx, callback)
	if err != nil {
		return err
	}

	return session.CommitTransaction(ctx)
}

func (u *userUseCase) DeleteOne(ctx context.Context, idUser string) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	userID, err := primitive.ObjectIDFromHex(idUser)
	if err != nil {
		return err
	}

	return u.userRepository.DeleteOne(ctx, userID)
}

func (u *userUseCase) FetchMany(ctx context.Context) ([]userdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	userData, err := u.userRepository.FetchMany(ctx)
	if err != nil {
		return nil, err
	}

	var outputs []userdomain.Output
	outputs = make([]userdomain.Output, 0, len(userData))
	for _, user := range userData {
		output := userdomain.Output{
			User: user,
		}

		outputs = append(outputs, output)
	}

	return outputs, nil
}

func (u *userUseCase) GetByEmail(ctx context.Context, email string) (userdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	userData, err := u.userRepository.GetByEmail(ctx, email)
	if err != nil {
		return userdomain.Output{}, err
	}

	output := userdomain.Output{
		User: userData,
	}

	return output, nil
}

func (u *userUseCase) GetByIDForCheckCookie(ctx context.Context, accessToken string) (*userdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	sub, err := token.ValidateToken(accessToken, u.database.AccessTokenPublicKey)
	if err != nil {
		return nil, err
	}

	userID, err := primitive.ObjectIDFromHex(fmt.Sprint(sub))
	if err != nil {
		return nil, err
	}

	userData, err := u.userRepository.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	output := &userdomain.Output{
		User: userData,
	}

	return output, nil
}

func (u *userUseCase) GetByID(ctx context.Context, id string) (*userdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	userData, err := u.userRepository.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	output := &userdomain.Output{
		User: userData,
	}

	return output, nil
}

func (u *userUseCase) RefreshToken(ctx context.Context, refreshToken string) (*userdomain.OutputLogin, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	sub, err := token.ValidateToken(refreshToken, u.database.RefreshTokenPublicKey)
	if err != nil {
		return nil, err
	}

	userID, err := primitive.ObjectIDFromHex(fmt.Sprint(sub))
	if err != nil {
		return nil, err
	}

	userData, err := u.userRepository.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	accessToken, err := token.CreateToken(u.database.AccessTokenExpiresIn, userData.ID, u.database.AccessTokenPrivateKey)
	if err != nil {
		return nil, err
	}

	refresh, err := token.CreateToken(u.database.RefreshTokenExpiresIn, userData.ID, u.database.RefreshTokenPrivateKey)
	if err != nil {
		return nil, err
	}

	response := &userdomain.OutputLogin{
		AccessToken:  accessToken,
		RefreshToken: refresh,
		IsLogged:     "1",
	}

	return response, nil
}
