package account_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	accountcontroller "shop_erp_mono/api/controllers/accounting_management/account"
	"shop_erp_mono/bootstrap"
	accountdomain "shop_erp_mono/domain/accounting_management/account"
	accountrepository "shop_erp_mono/repository/accounting_management/account/repository"
	accountusecase "shop_erp_mono/usecase/accounting_management/account/usecase"
	"time"
)

func AccountRouter(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
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
