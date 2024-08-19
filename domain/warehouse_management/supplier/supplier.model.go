package supplier

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionSupplier = "supplier"
)

// Supplier represents a supplier of products.
type Supplier struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	SupplierName  string             `bson:"supplier_name" json:"supplier_name"`
	ContactPerson string             `bson:"contact_person" json:"contact_person"`
	PhoneNumber   string             `bson:"phone_number" json:"phone_number"`
	Email         string             `bson:"email" json:"email"`
	Address       string             `bson:"address" json:"address"`
	CreatedAt     time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at" json:"updated_at"`
}
