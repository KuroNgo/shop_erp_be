package sale_report_usecase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	salereportsdomain "shop_erp_mono/domain/sales_and_distribution_management/sale_reports"
	productdomain "shop_erp_mono/domain/warehouse_management/product"
	"shop_erp_mono/usecase/sales_and_distribution_management/sale_report/validate"
	"time"
)

type saleReportUseCase struct {
	contextTimeout       time.Duration
	saleReportRepository salereportsdomain.ISalesReportRepository
	productRepository    productdomain.IProductRepository
}

func NewSaleReportUseCase(contextTimeout time.Duration, saleReportRepository salereportsdomain.ISalesReportRepository,
	productRepository productdomain.IProductRepository) salereportsdomain.ISalesReportUseCase {
	return &saleReportUseCase{contextTimeout: contextTimeout, saleReportRepository: saleReportRepository, productRepository: productRepository}
}

func (s *saleReportUseCase) CreateOne(ctx context.Context, input *salereportsdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	if err := validate.SaleReport(input); err != nil {
		return err
	}

	productData, err := s.productRepository.GetByName(ctx, input.ProductName)
	if err != nil {
		return err
	}

	saleReport := salereportsdomain.SalesReport{
		ID:           primitive.NewObjectID(),
		ReportDate:   input.ReportDate,
		TotalSales:   input.TotalSales,
		ProductID:    productData.ID,
		ProductName:  productData.ProductName,
		QuantitySold: input.QuantitySold,
		TotalRevenue: input.TotalRevenue,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	return s.saleReportRepository.CreateOne(ctx, saleReport)
}

func (s *saleReportUseCase) GetByID(ctx context.Context, id string) (*salereportsdomain.SalesReport, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	reportID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return s.saleReportRepository.GetByID(ctx, reportID)
}

func (s *saleReportUseCase) GetByDate(ctx context.Context, reportDate string) (*salereportsdomain.SalesReport, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	layout := "30-01-2006"
	date, err := time.Parse(layout, reportDate)
	if err != nil {
		return nil, err
	}

	return s.saleReportRepository.GetByDate(ctx, date)
}

func (s *saleReportUseCase) GetReportSummary(ctx context.Context, startDate, endDate string) (salereportsdomain.SalesReportReport, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	layout := "30-01-2006"
	start, err := time.Parse(layout, startDate)
	if err != nil {
		return salereportsdomain.SalesReportReport{}, err
	}

	end, err := time.Parse(layout, endDate)
	if err != nil {
		return salereportsdomain.SalesReportReport{}, err
	}

	data, err := s.saleReportRepository.GetReportSummary(ctx, start, end)
	if err != nil {
		return salereportsdomain.SalesReportReport{}, err
	}

	return data, nil
}

func (s *saleReportUseCase) GetTopSellingProducts(ctx context.Context, startDate, endDate time.Time) ([]salereportsdomain.TopSellingProduct, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	return s.saleReportRepository.GetTopSellingProducts(ctx, startDate, endDate)
}

func (s *saleReportUseCase) UpdateOne(ctx context.Context, id string, updatedReport *salereportsdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	if err := validate.SaleReport(updatedReport); err != nil {
		return err
	}

	reportID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	productData, err := s.productRepository.GetByName(ctx, updatedReport.ProductName)
	if err != nil {
		return err
	}

	saleReport := salereportsdomain.SalesReport{
		ID:           reportID,
		ReportDate:   updatedReport.ReportDate,
		TotalSales:   updatedReport.TotalSales,
		ProductID:    productData.ID,
		ProductName:  productData.ProductName,
		QuantitySold: updatedReport.QuantitySold,
		TotalRevenue: updatedReport.TotalRevenue,
		UpdatedAt:    time.Now(),
	}

	return s.saleReportRepository.UpdateOne(ctx, saleReport)
}

func (s *saleReportUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	reportID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return s.saleReportRepository.DeleteOne(ctx, reportID)
}

func (s *saleReportUseCase) GetAll(ctx context.Context) ([]salereportsdomain.SalesReport, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()
	return s.saleReportRepository.GetAll(ctx)
}
