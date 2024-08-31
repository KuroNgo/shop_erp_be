package contract_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	contractcontroller "shop_erp_mono/api/controllers/human_resources_management/contract"
	"shop_erp_mono/bootstrap"
	contractsdomain "shop_erp_mono/domain/human_resource_management/contracts"
	employeesdomain "shop_erp_mono/domain/human_resource_management/employees"
	contractrepository "shop_erp_mono/repository/human_resource_management/contracts/repository"
	contractusecase "shop_erp_mono/usecase/human_resource_management/contract/usecase"
	"time"
)

func ContractRouter(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	co := contractrepository.NewContractRepository(db, contractsdomain.CollectionContract, employeesdomain.CollectionEmployee)
	contract := &contractcontroller.ContractController{
		ContractUseCase: contractusecase.NewContractUseCase(timeout, co),
		Database:        env,
	}

	router := group.Group("/contracts")
	router.GET("/get/_id", contract.FetchOneByIDContract)
	router.GET("/get/email", contract.FetchOneByEmailContract)
	router.GET("/get/all", contract.FetchAllContract)
	router.POST("/create", contract.CreateOneContract)
	router.PUT("/update", contract.UpdateOneContract)
	router.DELETE("/delete", contract.DeleteOneContract)
}
