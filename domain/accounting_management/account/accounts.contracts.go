package account_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IAccountRepository interface {
	CreateAccount(ctx context.Context, account *Accounts) error
	GetAccountByID(ctx context.Context, id primitive.ObjectID) (Accounts, error)
	GetAccountByName(ctx context.Context, name string) (Accounts, error)
	UpdateAccount(ctx context.Context, account *Accounts) error
	DeleteAccount(ctx context.Context, id primitive.ObjectID) error
	ListAccounts(ctx context.Context) ([]Accounts, error)
}

type IAccountUseCase interface {
	CreateAccount(ctx context.Context, input *Input) error
	GetAccountByID(ctx context.Context, id string) (AccountResponse, error)
	GetAccountByName(ctx context.Context, name string) (AccountResponse, error)
	UpdateAccount(ctx context.Context, id string, input *Input) error
	DeleteAccount(ctx context.Context, id string) error
	ListAccounts(ctx context.Context) ([]AccountResponse, error)
}
