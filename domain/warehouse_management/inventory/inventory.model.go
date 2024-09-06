package inventory_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	products_domain "shop_erp_mono/domain/sales_and_distribution_management/products"
	"shop_erp_mono/domain/warehouse_management/warehouse"
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
	ID          primitive.ObjectID      `bson:"_id,omitempty" json:"id,omitempty"`
	Product     products_domain.Product `bson:"product_id" json:"product_id"`
	WarehouseID warehouse.Warehouse     `bson:"warehouse_id" json:"warehouse_id"`
	Quantity    int                     `bson:"quantity" json:"quantity"`
	CreatedAt   time.Time               `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time               `bson:"updated_at" json:"updated_at"`
}
