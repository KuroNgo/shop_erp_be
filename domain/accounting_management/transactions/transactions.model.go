package transactions_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionTransaction = "transaction"
)

type Transactions struct {
	TransactionsID  primitive.ObjectID `bson:"transactions_id" json:"transactions_id"`
	AccountID       primitive.ObjectID `bson:"account_id" json:"account_id"`
	Amount          float32            `bson:"amount" json:"amount"`
	TransactionType string             `bson:"transaction_type" json:"transaction_type"`
	Description     string             `bson:"description" json:"description"`
	CategoryID      primitive.ObjectID `bson:"category_id" json:"category_id"`
	CreatedAt       time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at" json:"updated_at"`
}
