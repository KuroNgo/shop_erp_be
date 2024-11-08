package sale_orders_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	customerdomain "shop_erp_mono/internal/domain/sales_and_distribution_management/customer"
	"time"
)

const (
	CollectionSalesOrder = "sales_order"
)

type SalesOrder struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	OrderNumber     string             `bson:"order_number"`
	CustomerID      primitive.ObjectID `bson:"customer_id"`
	OrderDate       time.Time          `bson:"order_date"`
	ShippingAddress string             `bson:"shipping_address"`
	TotalAmount     float64            `bson:"total_amount"`
	Status          string             `bson:"status"`
	CreatedAt       time.Time          `bson:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at"`
}

type Input struct {
	OrderNumber     string    `bson:"order_number"`
	CustomerID      string    `bson:"customer_email"`
	OrderDate       time.Time `bson:"order_date"`
	ShippingAddress string    `bson:"shipping_address"`
	TotalAmount     float64   `bson:"total_amount"`
	Status          string    `bson:"status"`
}

type SalesOrderResponse struct {
	Customer   customerdomain.Customer `bson:"customer"`
	SalesOrder SalesOrder              `bson:"salesOrder"`
}
