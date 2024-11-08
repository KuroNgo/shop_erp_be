package product_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	productcontroller "shop_erp_mono/internal/api/controllers/warehouse_management/product"
	"shop_erp_mono/internal/config"
	productdomain "shop_erp_mono/internal/domain/warehouse_management/product"
	categorydomain "shop_erp_mono/internal/domain/warehouse_management/product_category"
	categoryrepository "shop_erp_mono/internal/repository/warehouse_management/product_category/repository"
	productrepository "shop_erp_mono/internal/repository/warehouse_management/wm_product/repository"
	productusecase "shop_erp_mono/internal/usecase/warehouse_management/wm_product/usecase"
	"time"
)

func ProductRouter(env *config.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup, cacheTTL time.Duration) {
	pr := productrepository.NewProductRepository(db, productdomain.CollectionProduct)
	ca := categoryrepository.NewCategoryRepository(db, categorydomain.CollectionCategory)
	product := &productcontroller.ProductController{
		ProductUseCase: productusecase.NewProductUseCase(timeout, pr, ca, cacheTTL),
		Database:       env,
	}

	router := group.Group("/products")
	router.GET("/get/_id", product.GetOneByIDProduct)
	router.GET("/get/name", product.GetOneByNameProduct)
	router.GET("/get/all", product.GetAllProduct)
	router.POST("/create", product.CreateProduct)
	router.PUT("/update", product.UpdateProduct)
	router.DELETE("/delete", product.DeleteOneProduct)
}
