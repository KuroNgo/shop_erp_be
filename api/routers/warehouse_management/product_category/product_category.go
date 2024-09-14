package product_category_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	category_controller "shop_erp_mono/api/controllers/warehouse_management/product_category"
	"shop_erp_mono/bootstrap"
	productdomain "shop_erp_mono/domain/warehouse_management/product"
	categorydomain "shop_erp_mono/domain/warehouse_management/product_category"
	product_repository "shop_erp_mono/repository/warehouse_management/product/repository"
	category_repository "shop_erp_mono/repository/warehouse_management/product_category/repository"
	category_usecase "shop_erp_mono/usecase/warehouse_management/product_category/usecase"
	"time"
)

func ProductCategoryRouter(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	pr := product_repository.NewProductRepository(db, productdomain.CollectionProduct)
	ca := category_repository.NewInvoiceRepository(db, categorydomain.CollectionCategory)
	category := &category_controller.CategoryController{
		CategoryUseCase: category_usecase.NewCategoryUseCase(timeout, ca, pr),
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
