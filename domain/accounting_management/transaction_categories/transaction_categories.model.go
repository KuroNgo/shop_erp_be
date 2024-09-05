package transaction_categories_domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionTransactionCategory = "transaction_category"
)

type TransactionCategories struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	CategoryName string             `bson:"category_name" json:"categoryName"`
	CategoryType string             `bson:"category_type" json:"categoryType"`
}

type Input struct {
	CategoryName string `bson:"category_name" json:"categoryName"`
	CategoryType string `bson:"category_type" json:"categoryType"`
}
