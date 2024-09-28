package candidate_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	candidatecontroller "shop_erp_mono/api/controllers/human_resources_management/candidate"
	"shop_erp_mono/bootstrap"
	candidatedomain "shop_erp_mono/domain/human_resource_management/candidate"
	employeesdomain "shop_erp_mono/domain/human_resource_management/employees"
	candidaterepository "shop_erp_mono/repository/human_resource_management/candidate/repository"
	employeerepository "shop_erp_mono/repository/human_resource_management/employee/repository"
	candidateusecase "shop_erp_mono/usecase/human_resource_management/candidate/usecase"
	"time"
)

func CandidateRouter(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup, cacheTTL time.Duration) {
	ca := candidaterepository.NewCandidateRepository(db, candidatedomain.CollectionCandidate)
	em := employeerepository.NewEmployeeRepository(db, employeesdomain.CollectionEmployee)
	candidate := &candidatecontroller.CandidateController{
		CandidateUseCase: candidateusecase.NewCandidateUseCase(timeout, ca, em, cacheTTL),
		Database:         env,
	}

	router := group.Group("/candidates")
	router.GET("/get/_id", candidate.GetByID)
	router.GET("/get/name", candidate.GetByName)
	router.GET("/get/all", candidate.GetAll)
	router.POST("/create", candidate.CreateOne)
	router.PUT("/update", candidate.UpdateOne)
	router.DELETE("/delete", candidate.DeleteOne)
}
