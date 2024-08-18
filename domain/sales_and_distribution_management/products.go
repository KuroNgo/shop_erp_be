package sales_and_distribution_management

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Product represents a product in the inventory.
type Product struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ProductName     string             `bson:"product_name" json:"product_name"`
	Description     string             `bson:"description" json:"description"`
	Price           float64            `bson:"price" json:"price"`
	QuantityInStock int                `bson:"quantity_in_stock" json:"quantity_in_stock"`
	CategoryID      primitive.ObjectID `bson:"category_id" json:"category_id"`
	CreatedAt       time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at" json:"updated_at"`
}
