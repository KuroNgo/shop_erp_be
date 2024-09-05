package transaction_categories_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ITransactionCategoriesRepository interface {
	Create(ctx context.Context, category *TransactionCategories) error
	GetByID(ctx context.Context, id primitive.ObjectID) (TransactionCategories, error)
	GetByName(ctx context.Context, name string) (TransactionCategories, error)
	Update(ctx context.Context, category *TransactionCategories) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	List(ctx context.Context) ([]TransactionCategories, error)
}

type ITransactionCategoriesUseCase interface {
	CreateOne(ctx context.Context, input *Input) error
	GetTransactionCategoryByID(ctx context.Context, id string) (TransactionCategories, error)
	GetTransactionCategoryByName(ctx context.Context, name string) (TransactionCategories, error)
	UpdateOne(ctx context.Context, id string, input *Input) error
	DeleteOne(ctx context.Context, id string) error
	ListTransactionCategories(ctx context.Context) ([]TransactionCategories, error)
}
