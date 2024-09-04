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
	TotalAmount   int32     `bson:"total_amount" json:"totalAmount"`
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
