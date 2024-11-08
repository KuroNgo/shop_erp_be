package category_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionCategory = "product_category"
)

// Category represents a wm_product product_category.
type Category struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CategoryName string             `bson:"category_name" json:"category_name"`
	AssetURL     string             `bson:"asset_url"  json:"asset_url"`
	AvatarURL    string             `bson:"avatar_url"  json:"avatar_url"`
	Description  string             `bson:"description" json:"description"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
}

type Input struct {
	CategoryName string `bson:"category_name" json:"category_name"`
	Description  string `bson:"description" json:"description"`
}

type CategoryResponse struct {
	Category Category `bson:"product_category" json:"product_category"`
}
