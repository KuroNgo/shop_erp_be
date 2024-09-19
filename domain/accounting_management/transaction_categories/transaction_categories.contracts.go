package transaction_categories_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ITransactionCategoriesRepository interface {
	CreateOne(ctx context.Context, category *TransactionCategories) error
	GetByID(ctx context.Context, id primitive.ObjectID) (TransactionCategories, error)
	GetByName(ctx context.Context, name string) (TransactionCategories, error)
	UpdateOne(ctx context.Context, category *TransactionCategories) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	GetAll(ctx context.Context) ([]TransactionCategories, error)
}

type ITransactionCategoriesUseCase interface {
	CreateOne(ctx context.Context, input *Input) error
	GetByID(ctx context.Context, id string) (TransactionCategories, error)
	GetByName(ctx context.Context, name string) (TransactionCategories, error)
	UpdateOne(ctx context.Context, id string, input *Input) error
	DeleteOne(ctx context.Context, id string) error
	GetAll(ctx context.Context) ([]TransactionCategories, error)
}
