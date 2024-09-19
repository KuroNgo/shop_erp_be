package payments_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	sale_orders_domain "shop_erp_mono/domain/sales_and_distribution_management/sale_orders"
	"time"
)

const (
	CollectionPayment = "payments"
)

// Payment represents payment details for an order.
type Payment struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	OrderID       primitive.ObjectID `bson:"order_id" json:"order_id"`
	PaymentDate   time.Time          `bson:"payment_date" json:"payment_date"`
	PaymentMethod string             `bson:"payment_method" json:"payment_method"` // Example: "Credit Card", "Bank Transfer", "Cash on Delivery"
	AmountPaid    float64            `bson:"amount_paid" json:"amount_paid"`
	Status        string             `bson:"status" json:"status"` // Example: "Paid", "Unpaid"
	CreatedAt     time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at" json:"updated_at"`
}

type Input struct {
	OrderID       string    `bson:"order_id" json:"order_id"`
	PaymentDate   time.Time `bson:"payment_date" json:"payment_date"`
	PaymentMethod string    `bson:"payment_method" json:"payment_method"` // Example: "Credit Card", "Bank Transfer", "Cash on Delivery"
	AmountPaid    float64   `bson:"amount_paid" json:"amount_paid"`
	Status        string    `bson:"status" json:"status"` // Example: "Paid", "Unpaid"
}

type PaymentResponse struct {
	Payment Payment                       `bson:"payment" json:"payment"`
	Order   sale_orders_domain.SalesOrder `bson:"order" json:"order"`
}
