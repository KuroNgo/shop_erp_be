package sales_report_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	salereportsdomain "shop_erp_mono/domain/sales_and_distribution_management/sale_reports"
	"time"
)

type saleReportRepository struct {
	database             *mongo.Database
	saleReportCollection string
}

func NewSaleReportRepository(database *mongo.Database, saleReportCollection string) salereportsdomain.ISalesReportRepository {
	return &saleReportRepository{database: database, saleReportCollection: saleReportCollection}
}

func (s *saleReportRepository) CreateOne(ctx context.Context, report salereportsdomain.SalesReport) error {
	//TODO implement me
	panic("implement me")
}

func (s *saleReportRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*salereportsdomain.SalesReport, error) {
	//TODO implement me
	panic("implement me")
}

func (s *saleReportRepository) GetByDate(ctx context.Context, reportDate time.Time) (*salereportsdomain.SalesReport, error) {
	//TODO implement me
	panic("implement me")
}

func (s *saleReportRepository) GetReportSummary(ctx context.Context, startDate, endDate time.Time) (*salereportsdomain.SalesReportReport, error) {
	//TODO implement me
	panic("implement me")
}

func (s *saleReportRepository) GetTopSellingProducts(ctx context.Context, startDate, endDate time.Time) ([]salereportsdomain.TopSellingProduct, error) {
	//TODO implement me
	panic("implement me")
}

func (s *saleReportRepository) UpdateOne(ctx context.Context, id primitive.ObjectID, updatedReport salereportsdomain.SalesReport) error {
	//TODO implement me
	panic("implement me")
}

func (s *saleReportRepository) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	//TODO implement me
	panic("implement me")
}

func (s *saleReportRepository) List(ctx context.Context) ([]salereportsdomain.SalesReport, error) {
	//TODO implement me
	panic("implement me")
}
