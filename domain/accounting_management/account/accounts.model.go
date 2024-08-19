package account_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionAccount = "account"
)

type Accounts struct {
	AccountID     primitive.ObjectID `bson:"_id" json:"account_id"`
	AccountName   string             `bson:"account_name" json:"account_name"`
	AccountNumber string             `bson:"account_number" json:"account_number"`
	Balance       float64            `bson:"balance" json:"balance"`
	AccountType   string             `bson:"account_type" json:"account_type"`
	CreatedAt     time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at" json:"updated_at"`
}
