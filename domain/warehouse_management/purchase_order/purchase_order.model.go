package purchase_order_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionPurchaseOrder = "purchase_order"
)

// PurchaseOrder represents a purchase order from a supplier.
type PurchaseOrder struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	OrderNumber string             `bson:"order_number" json:"order_number"`
	SupplierID  primitive.ObjectID `bson:"supplier_id" json:"supplier_id"`
	OrderDate   time.Time          `bson:"order_date" json:"order_date"`
	TotalAmount float64            `bson:"total_amount" json:"total_amount"`
	Status      string             `bson:"status" json:"status"` // Example: "Processing", "Received", "Cancelled"
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

type Input struct {
	OrderNumber string    `bson:"order_number" json:"order_number"`
	SupplierID  string    `bson:"supplier" json:"supplier"`
	OrderDate   time.Time `bson:"order_date" json:"order_date"`
	TotalAmount float64   `bson:"total_amount" json:"total_amount"`
	Status      string    `bson:"status" json:"status"` // Example: "Processing", "Received", "Cancelled"
}

type PurchaseOrderResponse struct {
	PurchaseOrder PurchaseOrder `bson:"purchase_order" json:"purchase_order"`
}
