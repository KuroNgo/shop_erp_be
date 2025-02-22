package sale_reports_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	salesreportcontroller "shop_erp_mono/internal/api/controllers/sales_and_distributing_management/sale_reports"
	"shop_erp_mono/internal/config"
	salereportsdomain "shop_erp_mono/internal/domain/sales_and_distribution_management/sale_reports"
	productdomain "shop_erp_mono/internal/domain/warehouse_management/product"
	salesreportrepository "shop_erp_mono/internal/repository/sales_and_distribution_management/sale_report/repository"
	productrepository "shop_erp_mono/internal/repository/warehouse_management/wm_product/repository"
	salereportusecase "shop_erp_mono/internal/usecase/sales_and_distribution_management/sale_report/usecase"
	"time"
)

func SaleReportRouter(env *config.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup, cacheTTL time.Duration) {
	sr := salesreportrepository.NewSaleReportRepository(db, salereportsdomain.CollectionSalesReport)
	pr := productrepository.NewProductRepository(db, productdomain.CollectionProduct)
	salesReport := &salesreportcontroller.SalesReportController{
		SalesReportUseCase: salereportusecase.NewSaleReportUseCase(timeout, sr, pr, cacheTTL),
		Database:           env,
	}

	router := group.Group("/sales-reports")
	router.GET("/get/_id", salesReport.GetByID)
	router.GET("/get/date", salesReport.GetByDate)
	router.GET("/get/summary", salesReport.GetBySummary)
	router.POST("/create", salesReport.CreateOne)
	router.PUT("/update", salesReport.UpdateOne)
	router.DELETE("/delete", salesReport.DeleteOne)
}
