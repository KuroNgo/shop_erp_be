package transaction_category_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	transaction_categories_domain "shop_erp_mono/domain/accounting_management/transaction_categories"
)

type transactionCategoryRepository struct {
	database                      *mongo.Database
	transactionCategoryCollection string
}

func NewTransactionCategoryRepository(database *mongo.Database, transactionCategoryCollection string) transaction_categories_domain.ITransactionCategoriesRepository {
	return &transactionCategoryRepository{database: database, transactionCategoryCollection: transactionCategoryCollection}
}

func (t *transactionCategoryRepository) CreateOne(ctx context.Context, category *transaction_categories_domain.TransactionCategories) error {
	//TODO implement me
	panic("implement me")
}

func (t *transactionCategoryRepository) GetByID(ctx context.Context, id primitive.ObjectID) (transaction_categories_domain.TransactionCategories, error) {
	//TODO implement me
	panic("implement me")
}

func (t *transactionCategoryRepository) GetByName(ctx context.Context, name string) (transaction_categories_domain.TransactionCategories, error) {
	//TODO implement me
	panic("implement me")
}

func (t *transactionCategoryRepository) UpdateOne(ctx context.Context, category *transaction_categories_domain.TransactionCategories) error {
	//TODO implement me
	panic("implement me")
}

func (t *transactionCategoryRepository) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	//TODO implement me
	panic("implement me")
}

func (t *transactionCategoryRepository) GetAll(ctx context.Context) ([]transaction_categories_domain.TransactionCategories, error) {
	//TODO implement me
	panic("implement me")
}
