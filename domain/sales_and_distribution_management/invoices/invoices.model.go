package invoices_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionInvoice = "invoice"
)

// Invoice represents the invoice information for an order.
type Invoice struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	OrderID     primitive.ObjectID `bson:"order_id" json:"order_id"`
	InvoiceDate time.Time          `bson:"invoice_date" json:"invoice_date"`
	DueDate     time.Time          `bson:"due_date" json:"due_date"`
	AmountDue   float64            `bson:"amount_due" json:"amount_due"`
	AmountPaid  float64            `bson:"amount_paid" json:"amount_paid"`
	Status      string             `bson:"status" json:"status"` // Example: "Paid", "Unpaid", "Overdue"
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}
