package warehouse_management

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
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
