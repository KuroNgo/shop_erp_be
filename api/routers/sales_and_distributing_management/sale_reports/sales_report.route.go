package sale_reports_route

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	sales_report_controller "shop_erp_mono/api/controllers/sales_and_distributing_management/sale_reports"
	"shop_erp_mono/bootstrap"
	sale_reports_domain "shop_erp_mono/domain/sales_and_distribution_management/sale_reports"
	sales_report_repository "shop_erp_mono/repository/sales_and_distribution_management/sale_report/repository"
	sale_report_usecase "shop_erp_mono/usecase/sales_and_distribution_management/sale_report/usecase"
	"time"
)

func SaleReportRouter(env *bootstrap.Database, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	sr := sales_report_repository.NewSaleReportRepository(db, sale_reports_domain.CollectionSalesReport)
	salesReport := &sales_report_controller.SalesReportController{
		SalesReportUseCase: sale_report_usecase.NewSaleReportUseCase(timeout, sr),
		Database:           env,
	}

	router := group.Group("/sales_reports")
	router.GET("/get/_id", salesReport.GetByID)
	router.GET("/get/date", salesReport.GetByDate)
	router.GET("/get/summary", salesReport.GetBySummary)
	router.POST("/create", salesReport.CreateOne)
	router.PUT("/update", salesReport.UpdateOne)
	router.DELETE("/delete", salesReport.DeleteOne)
}
