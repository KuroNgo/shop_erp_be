package financial_reports_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type IFinancialReportsRepository interface {
	CreateOne(ctx context.Context, financialReport *FinancialReports) error
	GetFinancialReportsByID(ctx context.Context, id primitive.ObjectID) (FinancialReports, error)
	GetFinancialReportsByName(ctx context.Context, name string) (FinancialReports, error)
	UpdateOne(ctx context.Context, financialReport *FinancialReports) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	ListFinancialReports(ctx context.Context) ([]FinancialReports, error)
	GetFinancialReportsByDateRange(ctx context.Context, startDate, endDate time.Time) ([]FinancialReports, error)
	ArchiveFinancialReport(ctx context.Context, id primitive.ObjectID) error
}

type IFinancialReportsUseCase interface {
	CreateOne(ctx context.Context, input *Input) error
	GetFinancialReportsByID(ctx context.Context, id string) (FinancialReportsResponse, error)
	GetFinancialReportsByName(ctx context.Context, name string) (FinancialReportsResponse, error)
	UpdateOne(ctx context.Context, id string, input *Input) error
	DeleteOne(ctx context.Context, id string) error
	ListFinancialReports(ctx context.Context) ([]FinancialReportsResponse, error)
	GetFinancialReportsByDateRange(ctx context.Context, startDate, endDate string) ([]FinancialReportsResponse, error)
	ArchiveFinancialReport(ctx context.Context, id string) error
	GenerateFinancialReport(ctx context.Context, input *GenerateReportInput) (FinancialReportsResponse, error)
}
