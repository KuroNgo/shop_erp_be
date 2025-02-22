package inventory_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	productdomain "shop_erp_mono/internal/domain/warehouse_management/product"
	warehousedomain "shop_erp_mono/internal/domain/warehouse_management/warehouse"
	"time"
)

const (
	CollectionInventory = "inventory"
)

// Inventory represents the stock of a wm_product in a warehouse.
type Inventory struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ProductID       primitive.ObjectID `bson:"product_id" json:"product_id"`
	WarehouseID     primitive.ObjectID `bson:"warehouse_id" json:"warehouse_id"`
	Quantity        int                `bson:"quantity" json:"quantity"`
	QuantityWarning int                `bson:"quantity_warning" json:"quantity_warning"`
	ExpiryDate      *time.Time         `bson:"expiry_date,omitempty" json:"expiry_date,omitempty"` // if applicable
	Status          string             `bson:"status,omitempty" json:"status,omitempty"`           // e.g., "in stock", "out of stock"
	CreatedAt       time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at" json:"updated_at"`
}

type Input struct {
	ProductName   string     `bson:"product_id" json:"product_id"`
	WarehouseName string     `bson:"warehouse_id" json:"warehouse_id"`
	Quantity      int        `bson:"quantity" json:"quantity"`
	ExpiryDate    *time.Time `bson:"expiry_date,omitempty" json:"expiry_date,omitempty"` // if applicable
	Status        string     `bson:"status,omitempty" json:"status,omitempty"`           // e.g., "in stock", "out of stock"
}

type InventoryResponse struct {
	Inventory Inventory                 `bson:"inventory" json:"inventory"`
	Product   productdomain.Product     `bson:"wm_product" json:"wm_product"`
	Warehouse warehousedomain.Warehouse `bson:"warehouse" json:"warehouse"`
}
