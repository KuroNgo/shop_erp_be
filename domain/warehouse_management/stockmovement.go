package warehouse_management

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// StockMovement represents movements of stock in and out of a warehouse.
type StockMovement struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ProductID    primitive.ObjectID `bson:"product_id" json:"product_id"`
	WarehouseID  primitive.ObjectID `bson:"warehouse_id" json:"warehouse_id"`
	MovementType string             `bson:"movement_type" json:"movement_type"` // Example: "In", "Out"
	Quantity     int                `bson:"quantity" json:"quantity"`
	MovementDate time.Time          `bson:"movement_date" json:"movement_date"`
	Reference    string             `bson:"reference" json:"reference"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
}
