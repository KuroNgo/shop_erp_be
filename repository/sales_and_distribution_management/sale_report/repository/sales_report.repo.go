package sales_report_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
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
	saleReportCollection := s.database.Collection(s.saleReportCollection)

	_, err := saleReportCollection.InsertOne(ctx, report)
	if err != nil {
		return err
	}

	return nil
}

func (s *saleReportRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*salereportsdomain.SalesReport, error) {
	saleReportCollection := s.database.Collection(s.saleReportCollection)

	filter := bson.M{"_id": id}
	var report salereportsdomain.SalesReport
	if err := saleReportCollection.FindOne(ctx, filter).Decode(&report); err != nil {
		return nil, err
	}

	return &report, nil
}

func (s *saleReportRepository) GetByDate(ctx context.Context, reportDate time.Time) (*salereportsdomain.SalesReport, error) {
	saleReportCollection := s.database.Collection(s.saleReportCollection)

	filter := bson.M{"report_date": reportDate}
	var report salereportsdomain.SalesReport
	if err := saleReportCollection.FindOne(ctx, filter).Decode(&report); err != nil {
		return nil, err
	}

	return &report, nil
}

func (s *saleReportRepository) GetReportSummary(ctx context.Context, startDate, endDate time.Time) ([]salereportsdomain.SalesReport, error) {
	saleReportCollection := s.database.Collection(s.saleReportCollection)

	filter := bson.M{
		"report_date": bson.M{
			"$gte": startDate,
			"$lte": endDate,
		},
	}

	cursor, err := saleReportCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var reports []salereportsdomain.SalesReport
	if err := cursor.All(ctx, &reports); err != nil {
		return nil, err
	}

	return reports, nil
}

// need fix
func (s *saleReportRepository) GetTopSellingProducts(ctx context.Context, startDate, endDate time.Time) ([]salereportsdomain.TopSellingProduct, error) {
	saleReportCollection := s.database.Collection(s.saleReportCollection)

	filter := bson.M{
		"report_date": bson.M{
			"$gte": startDate,
			"$lte": endDate,
		},
	}

	cursor, err := saleReportCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var reports []salereportsdomain.TopSellingProduct
	if err := cursor.All(ctx, &reports); err != nil {
		return nil, err
	}

	return reports, nil
}

func (s *saleReportRepository) UpdateOne(ctx context.Context, updatedReport salereportsdomain.SalesReport) error {
	saleReportCollection := s.database.Collection(s.saleReportCollection)

	filter := bson.M{"_id": updatedReport.ID}
	update := bson.M{"$set": bson.M{
		"report_date":   updatedReport.ReportDate,
		"total_sales":   updatedReport.TotalSales,
		"product_id":    updatedReport.ProductID,
		"product_name":  updatedReport.ProductName,
		"quantity_sold": updatedReport.QuantitySold,
		"total_revenue": updatedReport.TotalRevenue,
		"updated_at":    time.Now(),
	}}
	_, err := saleReportCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (s *saleReportRepository) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	saleReportCollection := s.database.Collection(s.saleReportCollection)

	filter := bson.M{"_id": id}
	_, err := saleReportCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (s *saleReportRepository) List(ctx context.Context) ([]salereportsdomain.SalesReport, error) {
	saleReportCollection := s.database.Collection(s.saleReportCollection)

	filter := bson.M{}

	cursor, err := saleReportCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var reports []salereportsdomain.SalesReport
	if err := cursor.All(ctx, &reports); err != nil {
		return nil, err
	}

	return reports, nil
}
