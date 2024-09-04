package financial_reports_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionFinancialReports = "financial_reports"
)

type FinancialReports struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	ReportName  string             `bson:"report_name" json:"reportName"`
	StartDate   time.Time          `bson:"start_date" json:"startDate"`
	EndDate     time.Time          `bson:"end_date" json:"endDate"`
	GeneratedAt time.Time          `bson:"generated_at" json:"generatedAt"`
	Data        []string           `bson:"data" json:"data"`
}

type Input struct {
	ReportName string `bson:"report_name" json:"reportName"`
}

type FinancialReportsResponse struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	ReportName  string             `bson:"report_name" json:"reportName"`
	StartDate   time.Time          `bson:"start_date" json:"startDate"`
	EndDate     time.Time          `bson:"end_date" json:"endDate"`
	GeneratedAt time.Time          `bson:"generated_at" json:"generatedAt"`
	Data        []string           `bson:"data" json:"data"`
}
