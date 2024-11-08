package product_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionProduct = "wm_product"
)

// Product represents a wm_product in the inventory.
type Product struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ProductName     string             `bson:"product_name" json:"product_name"`
	AssetURL        string             `bson:"asset_url"  json:"asset_url"`
	AvatarURL       string             `bson:"avatar_url"  json:"avatar_url"`
	Description     string             `bson:"description" json:"description"`
	Price           float64            `bson:"price" json:"price"`
	ExpiryDate      *time.Time         `bson:"expiry_date" json:"expiry_date"`
	QuantityInStock int                `bson:"quantity_in_stock" json:"quantity_in_stock"`
	CategoryID      primitive.ObjectID `bson:"category_id" json:"category_id"`
	CreatedAt       time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at" json:"updated_at"`
}

type Input struct {
	ProductName     string     `bson:"product_name" json:"product_name"`
	AssetURL        string     `bson:"asset_url"  json:"asset_url"`
	AvatarURL       string     `bson:"avatar_url"  json:"avatar_url"`
	Description     string     `bson:"description" json:"description"`
	ExpiryDate      *time.Time `bson:"expiry_date" json:"expiry_date"`
	QuantityInStock int        `bson:"quantity_in_stock" json:"quantity_in_stock"`
	Price           float64    `bson:"price" json:"price"`
	Category        string     `bson:"product_category" json:"product_category"`
}

type ProductResponse struct {
	Product Product `json:"wm_product" bson:"wm_product"`
}
