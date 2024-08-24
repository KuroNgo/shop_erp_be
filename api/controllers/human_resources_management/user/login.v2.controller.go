package user_controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"net/http"
	userdomain "shop_erp_mono/domain/human_resource_management/user"
	googleoauth2 "shop_erp_mono/pkg/oauth2/google"
	"shop_erp_mono/pkg/token"
	"time"
)

func (auth *UserController) GoogleLoginWithUser(c *gin.Context) {
	var googleOauthConfig = &oauth2.Config{}
	googleOauthConfig = &oauth2.Config{
		ClientID:     auth.Database.GoogleClientID,
		ClientSecret: auth.Database.GoogleClientSecret,
		RedirectURL:  auth.Database.GoogleOAuthRedirectUrl,
		Scopes:       []string{"profile", "email"}, // Adjust scopes as needed
		Endpoint:     google.Endpoint,
	}

	code := c.Query("code")
	tokenData, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Println("Error exchanging code: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userInfo, err := googleoauth2.GetUserInfo(tokenData.AccessToken)
	if err != nil {
		fmt.Println("Error getting user info: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Giả sử userInfo là một map[string]interface{}
	email := userInfo["email"].(string)
	fullName := userInfo["name"].(string)
	avatarURL := userInfo["picture"].(string)
	verifiedEmail := userInfo["verified_email"].(bool)
	resBody := &userdomain.User{
		ID:        primitive.NewObjectID(),
		Email:     email,
		Username:  fullName,
		AvatarURL: avatarURL,
		Provider:  "google",
		Verified:  verifiedEmail,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	updateUser, err := auth.UserUseCase.UpsertOne(c, resBody.Email, resBody)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	signedToken, err := googleoauth2.SignJWT(userInfo)
	if err != nil {
		fmt.Println("Error signing token: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessTokenCh := make(chan string)
	refreshTokenCh := make(chan string)

	go func() {
		defer close(accessTokenCh)
		// Generate token
		accessToken, err := token.CreateToken(auth.Database.AccessTokenExpiresIn, updateUser.ID, auth.Database.AccessTokenPrivateKey)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": err.Error()},
			)
			return
		}
		accessTokenCh <- accessToken
	}()

	go func() {
		defer close(refreshTokenCh)
		refreshToken, err := token.CreateToken(auth.Database.RefreshTokenExpiresIn, updateUser.ID, auth.Database.RefreshTokenPrivateKey)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": err.Error()},
			)
			return
		}
		refreshTokenCh <- refreshToken
	}()

	accessToken := <-accessTokenCh
	refreshToken := <-refreshTokenCh

	c.SetCookie("access_token", accessToken, auth.Database.AccessTokenMaxAge*1000, "/", "localhost", false, true)
	c.SetCookie("refresh_token", refreshToken, auth.Database.AccessTokenMaxAge*1000, "/", "localhost", false, true)
	c.SetCookie("logged_in", "true", auth.Database.AccessTokenMaxAge*1000, "/", "localhost", false, false)

	c.JSON(http.StatusOK, gin.H{"token": signedToken})
}
