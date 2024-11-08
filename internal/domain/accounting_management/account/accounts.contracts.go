package account_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type IAccountRepository interface {
	CreateOne(ctx context.Context, account *Accounts) error
	GetByID(ctx context.Context, id primitive.ObjectID) (Accounts, error)
	GetByName(ctx context.Context, name string) (Accounts, error)
	UpdateOne(ctx context.Context, account *Accounts) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	GetAll(ctx context.Context) ([]Accounts, error)
	GetByDateRange(ctx context.Context, startDate, endDate time.Time) ([]Accounts, error)
	GetTotalAccountBalance(ctx context.Context) (float64, error)
	DeactivateAccount(ctx context.Context, id primitive.ObjectID) error
	ReactivateAccount(ctx context.Context, id primitive.ObjectID) error
}

type IAccountUseCase interface {
	CreateOne(ctx context.Context, input *Input) error
	GetByID(ctx context.Context, id string) (AccountResponse, error)
	GetByName(ctx context.Context, name string) (AccountResponse, error)
	UpdateOne(ctx context.Context, id string, input *Input) error
	DeleteOne(ctx context.Context, id string) error
	GetAll(ctx context.Context) ([]AccountResponse, error)
	GetByDateRange(ctx context.Context, startDate, endDate string) ([]AccountResponse, error)
	GetTotalBalance(ctx context.Context) (float64, error)
	DeactivateAccount(ctx context.Context, id string) error
	ReactivateAccount(ctx context.Context, id string) error
}
