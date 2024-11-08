package budget_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"shop_erp_mono/internal/config"
	"time"
)

func BudgetRouter(env *config.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	//bu := budget_repository.NewBudgetRepository(db, budgets_domain.CollectionBudgets)
	//budget := &budget_controller.BudgetController{
	//	BudgetUseCase: budget_usecase.NewBudgetUseCase(timeout, bu),
	//	Database:      env,
	//}
	//
	//router := group.Group("/budgets")
	//router.GET("/get/_id", account.GetByIDAccount)
	//router.GET("/get/name", account.GetByNameAccount)
	//router.GET("/get/all", account.GetAll)
	//router.POST("/create", account.CreateAccount)
	//router.PUT("/update", account.UpdateAccount)
	//router.DELETE("/delete", account.DeleteAccount)
}
