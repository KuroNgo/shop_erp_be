package order_details_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	sale_orders_domain "shop_erp_mono/domain/sales_and_distribution_management/sale_orders"
	productdomain "shop_erp_mono/domain/warehouse_management/product"
	"time"
)

const (
	CollectionOrderDetail = "order_detail"
)

// OrderDetail represents details of a single product within an order.
type OrderDetail struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	OrderID    primitive.ObjectID `bson:"order_id" json:"order_id"`
	ProductID  primitive.ObjectID `bson:"product_id" json:"product_id"`
	Quantity   int                `bson:"quantity" json:"quantity"`
	UnitPrice  float64            `bson:"unit_price" json:"unit_price"`
	TotalPrice float64            `bson:"total_price" json:"total_price"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
}

type Input struct {
	OrderID    string  `bson:"order_id" json:"order_id"`
	ProductID  string  `bson:"product_id" json:"product_id"`
	Quantity   int     `bson:"quantity" json:"quantity"`
	UnitPrice  float64 `bson:"unit_price" json:"unit_price"`
	TotalPrice float64 `bson:"total_price" json:"total_price"`
}

type OrderDetailResponse struct {
	OrderDetail OrderDetail                   `bson:"order_detail" json:"order_detail"`
	Order       sale_orders_domain.SalesOrder `bson:"order" json:"order"`
	Product     productdomain.Product         `bson:"product" json:"product"`
}
