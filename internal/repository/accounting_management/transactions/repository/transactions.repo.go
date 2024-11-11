package transactions_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	transactions_domain "shop_erp_mono/internal/domain/accounting_management/transactions"
)

type transactionRepository struct {
	database              *mongo.Database
	transactionCollection string
}

func (t *transactionRepository) CreateOne(ctx context.Context, transaction *transactions_domain.Transactions) error {
	//TODO implement me
	panic("implement me")
}

func (t *transactionRepository) GetByID(ctx context.Context, id primitive.ObjectID) (transactions_domain.Transactions, error) {
	//TODO implement me
	panic("implement me")
}

func (t *transactionRepository) GetByAccountID(ctx context.Context, accountID primitive.ObjectID) ([]transactions_domain.Transactions, error) {
	//TODO implement me
	panic("implement me")
}

func (t *transactionRepository) UpdateOne(ctx context.Context, transaction *transactions_domain.Transactions) error {
	//TODO implement me
	panic("implement me")
}

func (t *transactionRepository) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	//TODO implement me
	panic("implement me")
}

func (t *transactionRepository) GetAll(ctx context.Context) ([]transactions_domain.Transactions, error) {
	//TODO implement me
	panic("implement me")
}

func NewTransactionRepository(database *mongo.Database, transactionCollection string) transactions_domain.ITransactionsRepository {
	return &transactionRepository{database: database, transactionCollection: transactionCollection}
}
