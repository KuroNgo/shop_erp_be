package customer_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionCustomer = "customer"
)

// Customer represents a customer in the sales system.
type Customer struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName   string             `bson:"first_name" json:"first_name"`
	LastName    string             `bson:"last_name" json:"last_name"`
	Email       string             `bson:"email" json:"email"`
	PhoneNumber string             `bson:"phone_number" json:"phone_number"`
	Address     string             `bson:"address" json:"address"`
	City        string             `bson:"city" json:"city"`
	Country     string             `bson:"country" json:"country"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

type Input struct {
	FirstName   string `bson:"first_name" json:"first_name"`
	LastName    string `bson:"last_name" json:"last_name"`
	Email       string `bson:"email" json:"email"`
	PhoneNumber string `bson:"phone_number" json:"phone_number"`
	Address     string `bson:"address" json:"address"`
	City        string `bson:"city" json:"city"`
	Country     string `bson:"country" json:"country"`
}

type CustomerResponse struct {
	Customer             Customer `bson:"customer" json:"customer"`
	CountPurchaseHistory int      `bson:"count_purchase_history" json:"count_purchase_history"`
}
