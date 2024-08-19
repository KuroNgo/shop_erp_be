package sale_reports

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionSalesReport = "sales_report"
)

// SalesReport represents sales report data.
type SalesReport struct {
	ID                 primitive.ObjectID  `bson:"_id,omitempty" json:"id,omitempty"`
	ReportDate         time.Time           `bson:"report_date" json:"report_date"`
	TotalSales         float64             `bson:"total_sales" json:"total_sales"`
	TopSellingProducts []TopSellingProduct `bson:"top_selling_products" json:"top_selling_products"`
	CreatedAt          time.Time           `bson:"created_at" json:"created_at"`
	UpdatedAt          time.Time           `bson:"updated_at" json:"updated_at"`
}

// TopSellingProduct represents a top-selling product in a sales report.
type TopSellingProduct struct {
	ProductID    primitive.ObjectID `bson:"product_id" json:"product_id"`
	ProductName  string             `bson:"product_name" json:"product_name"`
	QuantitySold int                `bson:"quantity_sold" json:"quantity_sold"`
	TotalRevenue float64            `bson:"total_revenue" json:"total_revenue"`
}
