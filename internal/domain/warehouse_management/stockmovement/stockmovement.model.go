package stockmovement_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	userdomain "shop_erp_mono/internal/domain/human_resource_management/user"
	productdomain "shop_erp_mono/internal/domain/warehouse_management/product"
	warehousedomain "shop_erp_mono/internal/domain/warehouse_management/warehouse"
	"time"
)

const (
	CollectionStockMovement = "stock_movement"
)

// StockMovement represents movements of stock in and out of a warehouse.
type StockMovement struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ProductID    primitive.ObjectID `bson:"product_id" json:"product_id"`
	WarehouseID  primitive.ObjectID `bson:"warehouse_id" json:"warehouse_id"`
	MovementType string             `bson:"movement_type" json:"movement_type"`
	Quantity     int                `bson:"quantity" json:"quantity"`
	MovementDate time.Time          `bson:"movement_date" json:"movement_date"`
	UserID       primitive.ObjectID `bson:"user_id" json:"user_id"`
	Reference    string             `bson:"reference" json:"reference"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
}

type Input struct {
	Product      string    `bson:"wm_product" json:"wm_product"`
	Warehouse    string    `bson:"warehouse" json:"warehouse"`
	MovementType string    `bson:"movement_type" json:"movement_type"`
	Quantity     int       `bson:"quantity" json:"quantity"`
	MovementDate time.Time `bson:"movement_date" json:"movement_date"`
	User         string    `bson:"user" json:"user"`
	Reference    string    `bson:"reference" json:"reference"`
}

type StockMovementResponse struct {
	StockMovement StockMovement             `bson:"stock_movement" json:"stock_movement"`
	Product       productdomain.Product     `bson:"wm_product" json:"wm_product"`
	Warehouse     warehousedomain.Warehouse `bson:"warehouse" json:"warehouse"`
	User          userdomain.User           `bson:"user" json:"user"`
}
