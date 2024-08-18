package accounting_management_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Taxes struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	TaxName      string             `bson:"tax_name" json:"taxName"`
	Rate         float32            `bson:"rate" json:"rate"`
	ApplicableTo string             `bson:"applicable_to" json:"applicableTo"`
	CreatedAt    time.Time          `bson:"created_at" json:"createdAt"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updatedAt"`
}
