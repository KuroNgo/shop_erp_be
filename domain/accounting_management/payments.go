package accounting_management_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Payments struct {
	ID            primitive.ObjectID `bson:"_id" json:"id"`
	PaymentDate   time.Time          `bson:"payment_date" json:"paymentDate"`
	Amount        int32              `bson:"amount" json:"amount"`
	PaymentMethod string             `bson:"payment_method" json:"paymentMethod"`
	InvoiceID     primitive.ObjectID `bson:"invoice_id" json:"invoiceID"`
	AccountID     primitive.ObjectID `bson:"account_id" json:"accountID"`
	CreatedAt     time.Time          `bson:"created_at" json:"createdAt"`
	UpdatedAt     time.Time          `bson:"updated_at" json:"updatedAt"`
}
