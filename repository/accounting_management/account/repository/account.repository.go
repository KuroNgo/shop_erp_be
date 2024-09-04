package account_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	accountdomain "shop_erp_mono/domain/accounting_management/account"
	"time"
)

type accountRepository struct {
	database          *mongo.Database
	collectionAccount string
}

func NewAccountRepository(database *mongo.Database, collectionAccount string) accountdomain.IAccountRepository {
	return &accountRepository{database: database, collectionAccount: collectionAccount}
}

func (a accountRepository) CreateAccount(ctx context.Context, account *accountdomain.Accounts) error {
	collectionAccount := a.database.Collection(a.collectionAccount)

	_, err := collectionAccount.InsertOne(ctx, account)
	if err != nil {
		return err
	}

	return nil
}

func (a accountRepository) GetAccountByID(ctx context.Context, id primitive.ObjectID) (accountdomain.Accounts, error) {
	collectionAccount := a.database.Collection(a.collectionAccount)

	var account accountdomain.Accounts
	filter := bson.M{"_id": id}
	if err := collectionAccount.FindOne(ctx, filter).Decode(&account); err != nil {
		return accountdomain.Accounts{}, err
	}

	return account, nil
}

func (a accountRepository) GetAccountByName(ctx context.Context, name string) (accountdomain.Accounts, error) {
	collectionAccount := a.database.Collection(a.collectionAccount)

	var account accountdomain.Accounts
	filter := bson.M{"name": name}
	if err := collectionAccount.FindOne(ctx, filter).Decode(&account); err != nil {
		return accountdomain.Accounts{}, err
	}

	return account, nil
}

func (a accountRepository) UpdateAccount(ctx context.Context, budget *accountdomain.Accounts) error {
	collectionAccount := a.database.Collection(a.collectionAccount)

	filter := bson.M{"_id": budget.AccountID}
	update := bson.M{"$set": bson.M{
		"account_name":   budget.AccountName,
		"account_number": budget.AccountNumber,
		"balance":        budget.Balance,
		"account_type":   budget.AccountType,
		"updated_at":     time.Now(),
	}}

	_, err := collectionAccount.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (a accountRepository) DeleteAccount(ctx context.Context, id primitive.ObjectID) error {
	collectionAccount := a.database.Collection(a.collectionAccount)

	filter := bson.M{"_id": id}
	_, err := collectionAccount.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (a accountRepository) ListAccounts(ctx context.Context) ([]accountdomain.Accounts, error) {
	collectionAccount := a.database.Collection(a.collectionAccount)

	filter := bson.M{}
	cursor, err := collectionAccount.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	var accounts []accountdomain.Accounts
	accounts = make([]accountdomain.Accounts, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var account accountdomain.Accounts
		if err = cursor.Decode(&account); err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}
