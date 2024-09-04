package budgets_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IBudgetRepository interface {
	CreateBudget(ctx context.Context, budget *Budget) (Budget, error)
	GetBudgetByID(ctx context.Context, id primitive.ObjectID) (Budget, error)
	GetBudgetByName(ctx context.Context, name string) (Budget, error)
	UpdateBudget(ctx context.Context, budget *Budget) (Budget, error)
	DeleteBudget(ctx context.Context, id primitive.ObjectID) error
	ListBudgets(ctx context.Context) ([]Budget, error)
}

type IBudgetUseCase interface {
	CreateBudget(ctx context.Context, input *Input) (BudgetResponse, error)
	GetBudget(ctx context.Context, id string) (BudgetResponse, error)
	GetBudgetByName(ctx context.Context, name string) (Budget, error)
	UpdateBudget(ctx context.Context, id string, input *Input) (BudgetResponse, error)
	DeleteBudget(ctx context.Context, id string) error
	ListBudgets(ctx context.Context) ([]BudgetResponse, error)
}
