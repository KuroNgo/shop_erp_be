package contract_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	contractcontroller "shop_erp_mono/api/controllers/human_resources_management/contract"
	"shop_erp_mono/bootstrap"
	contractsdomain "shop_erp_mono/domain/human_resource_management/contracts"
	employeesdomain "shop_erp_mono/domain/human_resource_management/employees"
	contractrepository "shop_erp_mono/repository/human_resource_management/contracts/repository"
	employeerepository "shop_erp_mono/repository/human_resource_management/employee/repository"
	contractusecase "shop_erp_mono/usecase/human_resource_management/contract/usecase"
	"time"
)

func ContractRouter(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	co := contractrepository.NewContractRepository(db, contractsdomain.CollectionContract)
	em := employeerepository.NewEmployeeRepository(db, employeesdomain.CollectionEmployee)
	contract := &contractcontroller.ContractController{
		ContractUseCase: contractusecase.NewContractUseCase(timeout, co, em),
		Database:        env,
	}

	router := group.Group("/contracts")
	router.GET("/get/_id", contract.GetByID)
	router.GET("/get/email", contract.GetByEmail)
	router.GET("/get/all", contract.GetAll)
	router.POST("/create", contract.CreateOne)
	router.PUT("/update", contract.UpdateOne)
	router.DELETE("/delete", contract.DeleteOne)
}
