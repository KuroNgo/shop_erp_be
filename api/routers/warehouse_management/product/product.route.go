package product_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	productcontroller "shop_erp_mono/api/controllers/warehouse_management/product"
	"shop_erp_mono/bootstrap"
	productdomain "shop_erp_mono/domain/warehouse_management/product"
	categorydomain "shop_erp_mono/domain/warehouse_management/product_category"
	productrepository "shop_erp_mono/repository/warehouse_management/product/repository"
	categoryrepository "shop_erp_mono/repository/warehouse_management/product_category/repository"
	productusecase "shop_erp_mono/usecase/warehouse_management/product/usecase"
	"time"
)

func ProductRouter(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	pr := productrepository.NewProductRepository(db, productdomain.CollectionProduct)
	ca := categoryrepository.NewCategoryRepository(db, categorydomain.CollectionCategory)
	product := &productcontroller.ProductController{
		ProductUseCase: productusecase.NewProductUseCase(timeout, pr, ca),
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
