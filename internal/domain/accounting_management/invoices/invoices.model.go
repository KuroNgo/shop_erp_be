package invoices_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionInvoices = "invoice"
)

type Invoices struct {
	ID            primitive.ObjectID `bson:"_id" json:"id"`
	InvoiceNumber string             `bson:"invoice_number" json:"invoiceNumber"`
	InvoiceDate   time.Time          `bson:"invoice_date" json:"invoiceDate"`
	CustomerID    primitive.ObjectID `bson:"customer_id" json:"customerID"`
	TotalAmount   int32              `bson:"total_amount" json:"totalAmount"`
	Status        int                `bson:"status" json:"status"`
	DueDate       time.Time          `bson:"due_date" json:"dueDate"`
	CreatedAt     time.Time          `bson:"created_at" json:"createdAt"`
	UpdatedAt     time.Time          `bson:"updated_at" json:"updatedAt"`
}

type Input struct {
	InvoiceNumber string    `bson:"invoice_number" json:"invoiceNumber"`
	InvoiceDate   time.Time `bson:"invoice_date" json:"invoiceDate"`
	CustomerEmail string    `bson:"customer_email" json:"customer_email"`
	Status        int       `bson:"status" json:"status"`
	DueDate       time.Time `bson:"due_date" json:"dueDate"`
}

type InvoicesResponse struct {
	ID            primitive.ObjectID `bson:"_id" json:"id"`
	InvoiceNumber string             `bson:"invoice_number" json:"invoiceNumber"`
	InvoiceDate   time.Time          `bson:"invoice_date" json:"invoiceDate"`
	CustomerID    primitive.ObjectID `bson:"customer_id" json:"customerID"`
	TotalAmount   int32              `bson:"total_amount" json:"totalAmount"`
	Status        int                `bson:"status" json:"status"`
	DueDate       time.Time          `bson:"due_date" json:"dueDate"`
	CreatedAt     time.Time          `bson:"created_at" json:"createdAt"`
	UpdatedAt     time.Time          `bson:"updated_at" json:"updatedAt"`
}

type InvoiceReport struct {
	TotalInvoices      int       `json:"total_invoices"`       // Tổng số hóa đơn
	TotalRevenue       float64   `json:"total_revenue"`        // Tổng doanh thu từ các hóa đơn
	PaidInvoices       int       `json:"paid_invoices"`        // Số hóa đơn đã thanh toán
	UnpaidInvoices     int       `json:"unpaid_invoices"`      // Số hóa đơn chưa thanh toán
	OverdueInvoices    int       `json:"overdue_invoices"`     // Số hóa đơn quá hạn
	TotalOverdueAmount float64   `json:"total_overdue_amount"` // Tổng số tiền từ các hóa đơn quá hạn
	StartDate          time.Time `json:"start_date,omitempty"` // Ngày bắt đầu báo cáo
	EndDate            time.Time `json:"end_date,omitempty"`   // Ngày kết thúc báo cáo
}
