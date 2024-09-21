package purchase_order_detail_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	productdomain "shop_erp_mono/domain/warehouse_management/product"
	purchaseorderdomain "shop_erp_mono/domain/warehouse_management/purchase_order"
	"time"
)

const (
	CollectionPurchaseOrderDetail = "purchase_order_detail"
)

// PurchaseOrderDetail represents details of a wm_product in a purchase order.
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
	PurchaseOrderID primitive.ObjectID `bson:"purchase_order_id" json:"purchase_order_id"`
	Product         string             `bson:"wm_product" json:"wm_product"`
	Quantity        int                `bson:"quantity" json:"quantity"`
	UnitPrice       float64            `bson:"unit_price" json:"unit_price"`
	TotalPrice      float64            `bson:"total_price" json:"total_price"`
}

type PurchaseOrderDetailResponse struct {
	PurchaseOrderDetail PurchaseOrderDetail               `bson:"purchase_order_detail" json:"purchase_order_detail"`
	PurchaseOrder       purchaseorderdomain.PurchaseOrder `bson:"purchase_order" json:"purchase_order"`
	Product             productdomain.Product             `bson:"wm_product" json:"wm_product"`
}
