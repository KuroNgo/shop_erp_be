package transactions_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	accountdomain "shop_erp_mono/domain/accounting_management/account"
	transactioncategoriesdomain "shop_erp_mono/domain/accounting_management/transaction_categories"
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

type Input struct {
	Account         string  `bson:"account" json:"account"`
	Amount          float32 `bson:"amount" json:"amount"`
	TransactionType string  `bson:"transaction_type" json:"transaction_type"`
	Description     string  `bson:"description" json:"description"`
	CategoryID      string  `bson:"category_id" json:"category_id"`
}

type TransactionsResponse struct {
	TransactionsID  primitive.ObjectID                                `bson:"transactions_id" json:"transactions_id"`
	AccountID       accountdomain.Accounts                            `bson:"account_id" json:"account_id"`
	CategoryID      transactioncategoriesdomain.TransactionCategories `bson:"category_id" json:"category_id"`
	Amount          float32                                           `bson:"amount" json:"amount"`
	TransactionType string                                            `bson:"transaction_type" json:"transaction_type"`
	Description     string                                            `bson:"description" json:"description"`
	CreatedAt       time.Time                                         `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time                                         `bson:"updated_at" json:"updated_at"`
}
