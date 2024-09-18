package product_category_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	categorycontroller "shop_erp_mono/api/controllers/warehouse_management/product_category"
	"shop_erp_mono/bootstrap"
	productdomain "shop_erp_mono/domain/warehouse_management/product"
	categorydomain "shop_erp_mono/domain/warehouse_management/product_category"
	productrepository "shop_erp_mono/repository/warehouse_management/product/repository"
	categoryrepository "shop_erp_mono/repository/warehouse_management/product_category/repository"
	categoryusecase "shop_erp_mono/usecase/warehouse_management/product_category/usecase"
	"time"
)

func ProductCategoryRouter(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	pr := productrepository.NewProductRepository(db, productdomain.CollectionProduct)
	ca := categoryrepository.NewCategoryRepository(db, categorydomain.CollectionCategory)
	category := &categorycontroller.CategoryController{
		CategoryUseCase: categoryusecase.NewCategoryUseCase(timeout, ca, pr),
		Database:        env,
	}

	router := group.Group("/product_categories")
	router.GET("/get/_id", category.GetByIDCategories)
	router.GET("/get/name", category.GetByNameCategories)
	router.GET("/get/all", category.GetAllCategories)
	router.POST("/create", category.CreateCategory)
	router.PUT("/update", category.UpdateCategory)
	router.DELETE("/delete", category.DeleteCategory)
}
