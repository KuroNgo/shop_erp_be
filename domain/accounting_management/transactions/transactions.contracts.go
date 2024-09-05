package transactions_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ITransactionsRepository interface {
	Create(ctx context.Context, transaction *Transactions) error
	GetByID(ctx context.Context, id primitive.ObjectID) (Transactions, error)
	GetByAccountID(ctx context.Context, accountID primitive.ObjectID) ([]Transactions, error)
	Update(ctx context.Context, transaction *Transactions) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	List(ctx context.Context) ([]Transactions, error)
}

type ITransactionsUseCase interface {
	CreateTransaction(ctx context.Context, input *Input) error
	GetTransactionByID(ctx context.Context, id string) (TransactionsResponse, error)
	GetTransactionByAccountID(ctx context.Context, accountID string) ([]TransactionsResponse, error)
	UpdateTransaction(ctx context.Context, id string, input *Input) error
	DeleteTransaction(ctx context.Context, id string) error
	ListTransactions(ctx context.Context) ([]TransactionsResponse, error)
}
