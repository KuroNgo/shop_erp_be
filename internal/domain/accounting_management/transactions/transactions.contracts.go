package transactions_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ITransactionsRepository interface {
	CreateOne(ctx context.Context, transaction *Transactions) error
	GetByID(ctx context.Context, id primitive.ObjectID) (Transactions, error)
	GetByAccountID(ctx context.Context, accountID primitive.ObjectID) ([]Transactions, error)
	UpdateOne(ctx context.Context, transaction *Transactions) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	GetAll(ctx context.Context) ([]Transactions, error)
}

type ITransactionsUseCase interface {
	CreateOne(ctx context.Context, input *Input) error
	GetByID(ctx context.Context, id string) (TransactionsResponse, error)
	GetByAccountID(ctx context.Context, accountID string) ([]TransactionsResponse, error)
	UpdateOne(ctx context.Context, id string, input *Input) error
	DeleteOne(ctx context.Context, id string) error
	GetAll(ctx context.Context) ([]TransactionsResponse, error)
}
