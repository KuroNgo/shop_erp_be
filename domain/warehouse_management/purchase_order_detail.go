package warehouse_management

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
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
