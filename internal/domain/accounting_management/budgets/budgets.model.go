package budgets_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	transaction_categories_domain "shop_erp_mono/internal/domain/accounting_management/transaction_categories"
	"time"
)

const (
	CollectionBudgets = "budget"
)

type Budget struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	BudgetName string             `bson:"budget_name" json:"budgetName"`
	Amount     int32              `bson:"amount" json:"amount"`
	StartDate  time.Time          `bson:"start_date" json:"startDate"`
	EndDate    time.Time          `bson:"end_date" json:"endDate"`
	CategoryID primitive.ObjectID `bson:"category_id" json:"categoryID"`
	CreatedAt  time.Time          `bson:"created_at" json:"createdAt"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updatedAt"`
}

type Input struct {
	BudgetName string    `bson:"budget_name" json:"budget_name"`
	Amount     int32     `bson:"amount" json:"amount"`
	StartDate  time.Time `bson:"start_date" json:"start_date"`
	EndDate    time.Time `bson:"end_date" json:"end_date"`
	Category   string    `bson:"transaction_category" json:"transaction_category"`
}

type BudgetResponse struct {
	Budget              Budget
	TransactionCategory transaction_categories_domain.TransactionCategories
}
