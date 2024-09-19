package budgets_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type IBudgetRepository interface {
	CreateOne(ctx context.Context, budget *Budget) error
	GetByID(ctx context.Context, id primitive.ObjectID) (Budget, error)
	GetByName(ctx context.Context, name string) (Budget, error)
	UpdateOne(ctx context.Context, budget *Budget) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	GetAll(ctx context.Context) ([]Budget, error)
	GetByDateRange(ctx context.Context, startDate, endDate time.Time) ([]Budget, error)
	GetTotalAmount(ctx context.Context) (float64, error)
}

type IBudgetUseCase interface {
	CreateOne(ctx context.Context, input *Input) error
	GetByID(ctx context.Context, id string) (BudgetResponse, error)
	GetByName(ctx context.Context, name string) (BudgetResponse, error)
	UpdateOne(ctx context.Context, id string, input *Input) error
	DeleteOne(ctx context.Context, id string) error
	GetAll(ctx context.Context) ([]BudgetResponse, error)
	GetByDateRange(ctx context.Context, startDate, endDate string) ([]BudgetResponse, error)
	GetTotalAmount(ctx context.Context) (float64, error)
}
