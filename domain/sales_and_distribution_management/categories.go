package sales_and_distribution_management

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Category represents a product category.
type Category struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CategoryName string             `bson:"category_name" json:"category_name"`
	Description  string             `bson:"description" json:"description"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
}
