package shipping_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	sale_orders_domain "shop_erp_mono/domain/sales_and_distribution_management/sale_orders"
	"time"
)

const (
	CollectionShipping = "shipping"
)

// Shipping represents the shipping details for an order.
type Shipping struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	OrderID           primitive.ObjectID `bson:"order_id" json:"order_id"`
	ShippingMethod    string             `bson:"shipping_method" json:"shipping_method"` // Example: "Standard", "Express", "International"
	ShippingDate      time.Time          `bson:"shipping_date" json:"shipping_date"`
	EstimatedDelivery time.Time          `bson:"estimated_delivery" json:"estimated_delivery"`
	ActualDelivery    *time.Time         `bson:"actual_delivery,omitempty" json:"actual_delivery,omitempty"`
	TrackingNumber    string             `bson:"tracking_number" json:"tracking_number"`
	Status            string             `bson:"status" json:"status"` // Example: "In Transit", "Delivered", "Returned"
	CreatedAt         time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt         time.Time          `bson:"updated_at" json:"updated_at"`
}

type Input struct {
	OrderID           primitive.ObjectID `bson:"order_id" json:"order_id"`
	ShippingMethod    string             `bson:"shipping_method" json:"shipping_method"` // Example: "Standard", "Express", "International"
	ShippingDate      time.Time          `bson:"shipping_date" json:"shipping_date"`
	EstimatedDelivery time.Time          `bson:"estimated_delivery" json:"estimated_delivery"`
	ActualDelivery    *time.Time         `bson:"actual_delivery,omitempty" json:"actual_delivery,omitempty"`
	TrackingNumber    string             `bson:"tracking_number" json:"tracking_number"`
	Status            string             `bson:"status" json:"status"` // Example: "In Transit", "Delivered", "Returned"
}

type ShippingResponse struct {
	Shipping Shipping                      `bson:"shipping" json:"shipping"`
	Order    sale_orders_domain.SalesOrder `bson:"order" json:"order"`
}
