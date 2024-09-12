package user_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	usercontroller "shop_erp_mono/api/controllers/human_resources_management/user"
	"shop_erp_mono/bootstrap"
	userdomain "shop_erp_mono/domain/human_resource_management/user"
	userrepository "shop_erp_mono/repository/human_resource_management/user/repository"
	userusecase "shop_erp_mono/usecase/human_resource_management/user/usecase"
	"time"
)

func UserGoogleRouter(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	ur := userrepository.NewUserRepository(db, userdomain.CollectionUser)

	user := &usercontroller.UserController{
		UserUseCase: userusecase.NewUserUseCase(env, timeout, ur),
		Database:    env,
	}

	google := group.Group("/auth")
	google.GET("/google/callback", user.GoogleLoginWithUser)
}
