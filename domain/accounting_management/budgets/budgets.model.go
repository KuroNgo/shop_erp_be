package budgets_domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	BudgetName string             `bson:"budget_name" json:"budgetName"`
	Amount     int32              `bson:"amount" json:"amount"`
	StartDate  time.Time          `bson:"start_date" json:"startDate"`
	EndDate    time.Time          `bson:"end_date" json:"endDate"`
	CategoryID primitive.ObjectID `bson:"category_id" json:"categoryID"`
}

// BudgetResponse represents the response structure for a budget
type BudgetResponse struct {
	ID         primitive.ObjectID `json:"id"`
	BudgetName string             `json:"budgetName"`
	Amount     int32              `json:"amount"`
	StartDate  time.Time          `json:"startDate"`
	EndDate    time.Time          `json:"endDate"`
	CategoryID primitive.ObjectID `json:"categoryID"`
	CreatedAt  time.Time          `json:"createdAt"`
	UpdatedAt  time.Time          `json:"updatedAt"`
}
