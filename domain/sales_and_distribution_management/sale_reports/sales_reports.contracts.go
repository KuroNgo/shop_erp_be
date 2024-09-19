package sale_reports_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ISalesReportRepository interface {
	CreateOne(ctx context.Context, report SalesReport) error
	GetByID(ctx context.Context, id primitive.ObjectID) (*SalesReport, error)
	GetByDate(ctx context.Context, reportDate time.Time) (*SalesReport, error)
	GetReportSummary(ctx context.Context, startDate, endDate time.Time) (SalesReportReport, error)
	GetTopSellingProducts(ctx context.Context, startDate, endDate time.Time) ([]TopSellingProduct, error)
	UpdateOne(ctx context.Context, updatedReport SalesReport) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	List(ctx context.Context) ([]SalesReport, error)
}

type ISalesReportUseCase interface {
	CreateOne(ctx context.Context, input *Input) error
	GetByID(ctx context.Context, id string) (*SalesReport, error)
	GetByDate(ctx context.Context, reportDate string) (*SalesReport, error)
	GetReportSummary(ctx context.Context, startDate, endDate string) (SalesReportReport, error)
	GetTopSellingProducts(ctx context.Context, startDate, endDate time.Time) ([]TopSellingProduct, error)
	UpdateOne(ctx context.Context, id string, updatedReport *Input) error
	DeleteOne(ctx context.Context, id string) error
	List(ctx context.Context) ([]SalesReport, error)
}
