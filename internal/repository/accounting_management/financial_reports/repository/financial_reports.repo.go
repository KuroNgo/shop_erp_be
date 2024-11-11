package financial_reports_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	financialreportsdomain "shop_erp_mono/internal/domain/accounting_management/financial_reports"
	"time"
)

type financialRepository struct {
	database            *mongo.Database
	collectionFinancial string
}

func NewFinancialRepository(database *mongo.Database, collectionFinancial string) financialreportsdomain.IFinancialReportsRepository {
	return &financialRepository{database: database, collectionFinancial: collectionFinancial}
}

func (f *financialRepository) CreateOne(ctx context.Context, financialReport *financialreportsdomain.FinancialReports) error {
	//TODO implement me
	panic("implement me")
}

func (f *financialRepository) GetFinancialReportsByID(ctx context.Context, id primitive.ObjectID) (financialreportsdomain.FinancialReports, error) {
	//TODO implement me
	panic("implement me")
}

func (f *financialRepository) GetFinancialReportsByName(ctx context.Context, name string) (financialreportsdomain.FinancialReports, error) {
	//TODO implement me
	panic("implement me")
}

func (f *financialRepository) UpdateOne(ctx context.Context, financialReport *financialreportsdomain.FinancialReports) error {
	//TODO implement me
	panic("implement me")
}

func (f *financialRepository) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	//TODO implement me
	panic("implement me")
}

func (f *financialRepository) ListFinancialReports(ctx context.Context) ([]financialreportsdomain.FinancialReports, error) {
	//TODO implement me
	panic("implement me")
}

func (f *financialRepository) GetFinancialReportsByDateRange(ctx context.Context, startDate, endDate time.Time) ([]financialreportsdomain.FinancialReports, error) {
	//TODO implement me
	panic("implement me")
}

func (f *financialRepository) ArchiveFinancialReport(ctx context.Context, id primitive.ObjectID) error {
	//TODO implement me
	panic("implement me")
}
