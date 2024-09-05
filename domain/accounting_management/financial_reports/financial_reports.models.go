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

type GenerateReportInput struct {
	ReportName     string    `json:"report_name" validate:"required"`     // Tên báo cáo tài chính
	StartDate      time.Time `json:"start_date" validate:"required"`      // Ngày bắt đầu của dữ liệu báo cáo
	EndDate        time.Time `json:"end_date" validate:"required"`        // Ngày kết thúc của dữ liệu báo cáo
	RevenueSources []string  `json:"revenue_sources" validate:"required"` // Các nguồn doanh thu cần tổng hợp
	ExpenseSources []string  `json:"expense_sources" validate:"required"` // Các nguồn chi phí cần tổng hợp
	GeneratedBy    string    `json:"generated_by" validate:"required"`    // Tên người tạo báo cáo
	Notes          string    `json:"notes,omitempty"`                     // Ghi chú hoặc mô tả thêm (tùy chọn)
}
