package sale_report_usecase

import (
	"context"
	salereportsdomain "shop_erp_mono/domain/sales_and_distribution_management/sale_reports"
	"time"
)

type saleReportUseCase struct {
	contextTimeout       time.Duration
	saleReportRepository salereportsdomain.ISalesReportRepository
}

func NewSaleReportUseCase(contextTimeout time.Duration, saleReportRepository salereportsdomain.ISalesReportRepository) salereportsdomain.ISalesReportUseCase {
	return &saleReportUseCase{contextTimeout: contextTimeout, saleReportRepository: saleReportRepository}
}

func (s *saleReportUseCase) CreateOne(ctx context.Context, input *salereportsdomain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (s *saleReportUseCase) GetByID(ctx context.Context, id string) (*salereportsdomain.SalesReport, error) {
	//TODO implement me
	panic("implement me")
}

func (s *saleReportUseCase) GetByDate(ctx context.Context, reportDate string) (*salereportsdomain.SalesReport, error) {
	//TODO implement me
	panic("implement me")
}

func (s *saleReportUseCase) GetReportSummary(ctx context.Context, startDate, endDate string) (*salereportsdomain.SalesReportReport, error) {
	//TODO implement me
	panic("implement me")
}

func (s *saleReportUseCase) GetTopSellingProducts(ctx context.Context, startDate, endDate time.Time) ([]salereportsdomain.TopSellingProduct, error) {
	//TODO implement me
	panic("implement me")
}

func (s *saleReportUseCase) UpdateOne(ctx context.Context, id string, updatedReport *salereportsdomain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (s *saleReportUseCase) DeleteOne(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (s *saleReportUseCase) List(ctx context.Context) ([]salereportsdomain.SalesReport, error) {
	//TODO implement me
	panic("implement me")
}
