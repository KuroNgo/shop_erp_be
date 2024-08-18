package sales_and_distribution_management

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// SalesOrder represents a sales order from a customer.
type SalesOrder struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	OrderNumber     string             `bson:"order_number" json:"order_number"`
	CustomerID      primitive.ObjectID `bson:"customer_id" json:"customer_id"`
	OrderDate       time.Time          `bson:"order_date" json:"order_date"`
	ShippingAddress string             `bson:"shipping_address" json:"shipping_address"`
	TotalAmount     float64            `bson:"total_amount" json:"total_amount"`
	Status          string             `bson:"status" json:"status"` // Example: "Processing", "Shipped", "Delivered", "Cancelled"
	CreatedAt       time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at" json:"updated_at"`
}
