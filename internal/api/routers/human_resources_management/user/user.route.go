package user_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	usercontroller "shop_erp_mono/internal/api/controllers/human_resources_management/user"
	"shop_erp_mono/internal/api/middlewares"
	"shop_erp_mono/internal/config"
	userdomain "shop_erp_mono/internal/domain/human_resource_management/user"
	userrepository "shop_erp_mono/internal/repository/human_resource_management/user/repository"
	userusecase "shop_erp_mono/internal/usecase/human_resource_management/user/usecase"
	"shop_erp_mono/pkg/interface/cloud/cloudinary/middlewares"
	"time"
)

func UserRouter(env *config.Database, timeout time.Duration, db *mongo.Database, client *mongo.Client, group *gin.RouterGroup) {
	ur := userrepository.NewUserRepository(db, userdomain.CollectionUser)

	user := &usercontroller.UserController{
		UserUseCase: userusecase.NewUserUseCase(env, timeout, ur, client),
		Database:    env,
	}

	router := group.Group("/users")
	router.POST("/login", middlewares.RateLimiter(), user.LoginUser)
	router.POST("/signup", middlewares_cloudinary.FileUploadMiddleware(), user.SignUp)
	router.PATCH("/update", middlewares_cloudinary.FileUploadMiddleware(), middlewares.DeserializeUser(), user.UpdateUser)
	router.PATCH("/verify", user.VerificationCode)
	router.PATCH("/verify/password", user.VerificationCodeForChangePassword)
	router.PATCH("/password/forget", user.ChangePassword)
	router.POST("/forget", user.ForgetPasswordInUser)
	router.GET("/get/info", user.GetMe)
	router.GET("/get/refresh", user.RefreshToken)
	router.DELETE("/current/delete", middlewares.DeserializeUser(), user.DeleteCurrentUser)
	router.GET("/logout", middlewares.DeserializeUser(), user.LogoutUser)

	google := group.Group("/auth")
	google.GET("/google/callback", user.GoogleLoginWithUser)
}
