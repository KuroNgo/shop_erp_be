package candidate_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	candidatecontroller "shop_erp_mono/internal/api/controllers/human_resources_management/candidate"
	"shop_erp_mono/internal/config"
	candidatedomain "shop_erp_mono/internal/domain/human_resource_management/candidate"
	employeesdomain "shop_erp_mono/internal/domain/human_resource_management/employees"
	candidaterepository "shop_erp_mono/internal/repository/human_resource_management/candidate/repository"
	employeerepository "shop_erp_mono/internal/repository/human_resource_management/employee/repository"
	candidateusecase "shop_erp_mono/internal/usecase/human_resource_management/candidate/usecase"
	"time"
)

func CandidateRouter(env *config.Database, timeout time.Duration, db *mongo.Database, client *mongo.Client, group *gin.RouterGroup, cacheTTL time.Duration) {
	ca := candidaterepository.NewCandidateRepository(db, candidatedomain.CollectionCandidate)
	em := employeerepository.NewEmployeeRepository(db, employeesdomain.CollectionEmployee)
	candidate := &candidatecontroller.CandidateController{
		CandidateUseCase: candidateusecase.NewCandidateUseCase(timeout, ca, em, cacheTTL, client),
		Database:         env,
	}

	router := group.Group("/candidates")
	router.GET("/get/_id", candidate.GetByID)
	router.GET("/get/name", candidate.GetByName)
	router.GET("/get/all", candidate.GetAll)
	//router.GET("/get/all/pagination", candidate.GetAll)
	router.POST("/create", candidate.CreateOne)
	router.PUT("/update", candidate.UpdateOne)
	router.PATCH("/update/status", candidate.UpdateStatus)
	//router.PATCH("/delete/_id", candidate.D)
	router.DELETE("/delete", candidate.DeleteOne)
}
