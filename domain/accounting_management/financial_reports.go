package accounting_management_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type FinancialReports[T interface{}] struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	ReportName  string             `bson:"report_name" json:"reportName"`
	StartDate   time.Time          `bson:"start_date" json:"startDate"`
	EndDate     time.Time          `bson:"end_date" json:"endDate"`
	GeneratedAt time.Time          `bson:"generated_at" json:"generatedAt"`
	Data        []T                `bson:"data" json:"data"`
}
