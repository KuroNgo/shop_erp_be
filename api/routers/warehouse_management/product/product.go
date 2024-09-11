package product_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	product_controller "shop_erp_mono/api/controllers/warehouse_management/product"
	"shop_erp_mono/bootstrap"
	productdomain "shop_erp_mono/domain/warehouse_management/product"
	categorydomain "shop_erp_mono/domain/warehouse_management/product_category"
	category_repository "shop_erp_mono/repository/warehouse_management/category/repository"
	product_repository "shop_erp_mono/repository/warehouse_management/product/repository"
	product_usecase "shop_erp_mono/usecase/warehouse_management/product/usecase"
	"time"
)

func ProductRouter(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	pr := product_repository.NewProductRepository(db, productdomain.CollectionProduct)
	ca := category_repository.NewInvoiceRepository(db, categorydomain.CollectionCategory)
	product := &product_controller.ProductController{
		ProductUseCase: product_usecase.NewProductUseCase(timeout, pr, ca),
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
