package user_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	usercontroller "shop_erp_mono/internal/api/controllers/human_resources_management/user"
	"shop_erp_mono/internal/config"
	userdomain "shop_erp_mono/internal/domain/human_resource_management/user"
	userrepository "shop_erp_mono/internal/repository/human_resource_management/user/repository"
	userusecase "shop_erp_mono/internal/usecase/human_resource_management/user/usecase"
	"time"
)

func UserGoogleRouter(env *config.Database, timeout time.Duration, db *mongo.Database, client *mongo.Client, group *gin.RouterGroup) {
	ur := userrepository.NewUserRepository(db, userdomain.CollectionUser)

	user := &usercontroller.UserController{
		UserUseCase: userusecase.NewUserUseCase(env, timeout, ur, client),
		Database:    env,
	}

	google := group.Group("/auth")
	google.GET("/google/callback", user.GoogleLoginWithUser)
}
