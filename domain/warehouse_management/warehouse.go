package warehouse_management

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Warehouse represents a storage warehouse.
type Warehouse struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	WarehouseName string             `bson:"warehouse_name" json:"warehouse_name"`
	Location      string             `bson:"location" json:"location"`
	Capacity      int                `bson:"capacity" json:"capacity"`
	CreatedAt     time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at" json:"updated_at"`
}
