package salary_base_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"shop_erp_mono/internal/api/controllers/human_resources_management/salary_base"
	"shop_erp_mono/internal/config"
	roledomain "shop_erp_mono/internal/domain/human_resource_management/role"
	base_salary_domain "shop_erp_mono/internal/domain/human_resource_management/salary_base"
	rolerepository "shop_erp_mono/internal/repository/human_resource_management/role/repository"
	base_salary_repository "shop_erp_mono/internal/repository/human_resource_management/salary_base/repository"
	base_salary_usecase "shop_erp_mono/internal/usecase/human_resource_management/salary_base/usecase"
	"time"
)

func BaseSalaryRouter(env *config.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup, cacheTTL time.Duration) {
	sa := base_salary_repository.NewBaseSalaryRepository(db, base_salary_domain.CollectionBaseSalary)
	ro := rolerepository.NewRoleRepository(db, roledomain.CollectionRole)
	salary := &salary_base.BaseSalaryController{
		BaseSalaryUseCase: base_salary_usecase.NewBaseSalaryUseCase(timeout, sa, ro, cacheTTL),
		Database:          env,
	}

	router := group.Group("/base-salaries")
	router.GET("/get/_id", salary.GetByID)
	router.GET("/get/title", salary.GetByID)
	router.GET("/get/all", salary.GetAll)
	router.POST("/create", salary.CreateOne)
	router.PUT("/update", salary.UpdateOne)
	router.DELETE("/delete", salary.DeleteOne)

}
