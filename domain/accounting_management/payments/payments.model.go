package payment_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	accountdomain "shop_erp_mono/domain/accounting_management/account"
	invoices_domain "shop_erp_mono/domain/sales_and_distribution_management/invoices"
	"time"
)

const (
	CollectionPayments = "payment"
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

type Input struct {
	PaymentDate   time.Time          `bson:"payment_date" json:"paymentDate"`
	Amount        int32              `bson:"amount" json:"amount"`
	PaymentMethod string             `bson:"payment_method" json:"paymentMethod"`
	InvoiceID     primitive.ObjectID `bson:"invoice_id" json:"invoiceID"`
	AccountID     primitive.ObjectID `bson:"account_id" json:"accountID"`
}

type PaymentsResponse struct {
	ID            primitive.ObjectID      `bson:"_id" json:"id"`
	PaymentDate   time.Time               `bson:"payment_date" json:"paymentDate"`
	Amount        int32                   `bson:"amount" json:"amount"`
	PaymentMethod string                  `bson:"payment_method" json:"paymentMethod"`
	Invoice       invoices_domain.Invoice `bson:"invoice" json:"invoice"`
	AccountID     accountdomain.Accounts  `bson:"account" json:"accountID"`
	CreatedAt     time.Time               `bson:"created_at" json:"createdAt"`
	UpdatedAt     time.Time               `bson:"updated_at" json:"updatedAt"`
}
