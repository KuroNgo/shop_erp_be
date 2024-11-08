package account_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	accountcontroller "shop_erp_mono/internal/api/controllers/accounting_management/account"
	"shop_erp_mono/internal/config"
	accountdomain "shop_erp_mono/internal/domain/accounting_management/account"
	accountrepository "shop_erp_mono/internal/repository/accounting_management/account/repository"
	accountusecase "shop_erp_mono/internal/usecase/accounting_management/account/usecase"
	"time"
)

func AccountRouter(env *config.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	ac := accountrepository.NewAccountRepository(db, accountdomain.CollectionAccount)

	account := &accountcontroller.AccountController{
		AccountUseCase: accountusecase.NewAccountUseCase(timeout, ac),
		Database:       env,
	}

	router := group.Group("/accounts")
	router.GET("/get/_id", account.GetByID)
	router.GET("/get/name", account.GetByName)
	router.GET("/get/all", account.GetAll)
	router.POST("/create", account.CreateOne)
	router.PUT("/update", account.UpdateOne)
	router.DELETE("/delete", account.DeleteOne)
}
