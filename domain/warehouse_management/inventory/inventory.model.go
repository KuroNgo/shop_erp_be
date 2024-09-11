package inventory_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	productdomain "shop_erp_mono/domain/warehouse_management/product"
	warehousedomain "shop_erp_mono/domain/warehouse_management/warehouse"
	"time"
)

const (
	CollectionInventory = "inventory"
)

// Inventory represents the stock of a product in a warehouse.
type Inventory struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ProductID   primitive.ObjectID `bson:"product_id" json:"product_id"`
	WarehouseID primitive.ObjectID `bson:"warehouse_id" json:"warehouse_id"`
	Quantity    int                `bson:"quantity" json:"quantity"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

type Input struct {
	ProductName   string `bson:"product_id" json:"product_id"`
	WarehouseName string `bson:"warehouse_id" json:"warehouse_id"`
	Quantity      int    `bson:"quantity" json:"quantity"`
}

type InventoryResponse struct {
	ID        primitive.ObjectID        `bson:"_id,omitempty" json:"id,omitempty"`
	Product   productdomain.Product     `bson:"product" json:"product"`
	Warehouse warehousedomain.Warehouse `bson:"warehouse" json:"warehouse"`
	Quantity  int                       `bson:"quantity" json:"quantity"`
	CreatedAt time.Time                 `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time                 `bson:"updated_at" json:"updated_at"`
}
