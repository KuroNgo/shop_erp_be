package stock_adjustment_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	product_domain "shop_erp_mono/internal/domain/warehouse_management/product"
	warehouse_domain "shop_erp_mono/internal/domain/warehouse_management/warehouse"
	"time"
)

const (
	CollectionStockAdjustment = "stock_adjustment"
)

// StockAdjustment represents adjustments made to inventory levels.
type StockAdjustment struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ProductID      primitive.ObjectID `bson:"product_id" json:"product_id"`
	WarehouseID    primitive.ObjectID `bson:"warehouse_id" json:"warehouse_id"`
	AdjustmentType string             `bson:"adjustment_type" json:"adjustment_type"` // Example: "Increase", "Decrease"
	Quantity       int                `bson:"quantity" json:"quantity"`
	Reason         string             `bson:"reason" json:"reason"`
	AdjustmentDate time.Time          `bson:"adjustment_date" json:"adjustment_date"`
	CreatedAt      time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at" json:"updated_at"`
}

type Input struct {
	Product        string    `bson:"wm_product" json:"wm_product"`
	Warehouse      string    `bson:"warehouse" json:"warehouse"`
	AdjustmentType string    `bson:"adjustment_type" json:"adjustment_type"` // Example: "Increase", "Decrease"
	Quantity       int       `bson:"quantity" json:"quantity"`
	Reason         string    `bson:"reason" json:"reason"`
	AdjustmentDate time.Time `bson:"adjustment_date" json:"adjustment_date"`
}

type StockAdjustmentResponse struct {
	StockAdjustment StockAdjustment            `bson:"stock_adjustment" json:"stock_adjustment"`
	Product         product_domain.Product     `bson:"wm_product" json:"wm_product"`
	Warehouse       warehouse_domain.Warehouse `bson:"warehouse_id" json:"warehouse_id"`
}
