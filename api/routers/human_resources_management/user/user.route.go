package user_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	usercontroller "shop_erp_mono/api/controllers/human_resources_management/user"
	"shop_erp_mono/api/middlewares"
	"shop_erp_mono/bootstrap"
	userdomain "shop_erp_mono/domain/human_resource_management/user"
	userrepository "shop_erp_mono/repository/human_resource_management/user/repository"
	userusecase "shop_erp_mono/usecase/human_resource_management/user/usecase"
	"time"
)

func UserRouter(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	ur := userrepository.NewUserRepository(db, userdomain.CollectionUser)

	user := &usercontroller.UserController{
		UserUseCase: userusecase.NewUserUseCase(env, timeout, ur),
		Database:    env,
	}

	router := group.Group("/users")
	router.POST("/login", middlewares.RateLimiter(), user.LoginUser)
	//router.GET("/google/callback", user.GoogleLoginWithUser)
	//router.POST("/signup", user.SignUp)
	//router.PATCH("/update", middlewares.DeserializeUser(), user.UpdateUser)
	//router.PATCH("/verify", user.VerificationCode)
	//router.PATCH("/verify/password", user.VerificationCodeForChangePassword)
	//router.PATCH("/password/forget", user.ChangePassword)
	//router.POST("/forget", user.ForgetPasswordInUser)
	router.GET("/get/info", user.GetMe)
	//router.GET("/get/refresh", user.RefreshToken)
	router.DELETE("/current/delete", middlewares.DeserializeUser(), user.DeleteCurrentUser)
	router.GET("/logout", middlewares.DeserializeUser(), user.LogoutUser)
}
