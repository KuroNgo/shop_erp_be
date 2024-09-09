package product_order_detail

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	productsdomain "shop_erp_mono/domain/sales_and_distribution_management/products"
	purchaseorderdomain "shop_erp_mono/domain/warehouse_management/purchase_order"
	"time"
)

const (
	CollectionPurchaseOrderDetail = "purchase_order_detail"
)

// PurchaseOrderDetail represents details of a product in a purchase order.
type PurchaseOrderDetail struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	PurchaseOrderID primitive.ObjectID `bson:"purchase_order_id" json:"purchase_order_id"`
	ProductID       primitive.ObjectID `bson:"product_id" json:"product_id"`
	Quantity        int                `bson:"quantity" json:"quantity"`
	UnitPrice       float64            `bson:"unit_price" json:"unit_price"`
	TotalPrice      float64            `bson:"total_price" json:"total_price"`
	CreatedAt       time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at" json:"updated_at"`
}

type Input struct {
	PurchaseOrder string  `bson:"purchase_order" json:"purchase_order"`
	Product       string  `bson:"product" json:"product"`
	Quantity      int     `bson:"quantity" json:"quantity"`
	UnitPrice     float64 `bson:"unit_price" json:"unit_price"`
	TotalPrice    float64 `bson:"total_price" json:"total_price"`
}

type PurchaseOrderDetailResponse struct {
	ID              primitive.ObjectID                `bson:"_id,omitempty" json:"id,omitempty"`
	PurchaseOrderID purchaseorderdomain.PurchaseOrder `bson:"purchase_order_id" json:"purchase_order_id"`
	ProductID       productsdomain.Product            `bson:"product_id" json:"product_id"`
	Quantity        int                               `bson:"quantity" json:"quantity"`
	UnitPrice       float64                           `bson:"unit_price" json:"unit_price"`
	TotalPrice      float64                           `bson:"total_price" json:"total_price"`
	CreatedAt       time.Time                         `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time                         `bson:"updated_at" json:"updated_at"`
}
