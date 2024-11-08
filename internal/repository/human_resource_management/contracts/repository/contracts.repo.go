package contract_repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	contractsdomain "shop_erp_mono/internal/domain/human_resource_management/contracts"
)

type contractRepository struct {
	database           *mongo.Database
	collectionContract string
}

func NewContractRepository(db *mongo.Database, collectionContract string) contractsdomain.IContractsRepository {
	return &contractRepository{database: db, collectionContract: collectionContract}
}

func (c *contractRepository) CreateOne(ctx context.Context, contract *contractsdomain.Contract) error {
	collectionContract := c.database.Collection(c.collectionContract)

	_, err := collectionContract.InsertOne(ctx, contract)
	if err != nil {
		return err
	}

	return nil
}

func (c *contractRepository) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	collectionContract := c.database.Collection(c.collectionContract)

	if id == primitive.NilObjectID {
		return errors.New("the contractID do not nil")
	}

	filter := bson.M{"_id": id}
	_, err := collectionContract.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (c *contractRepository) UpdateOne(ctx context.Context, id primitive.ObjectID, contract *contractsdomain.Contract) error {
	collectionContract := c.database.Collection(c.collectionContract)

	if id == primitive.NilObjectID {
		return errors.New("id do not nil")
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"employee_id":   contract.EmployeeID,
		"contract_type": contract.ContractType,
		"start_date":    contract.StartDate,
		"end_date":      contract.EndDate,
		"salary":        contract.Salary,
		"status":        contract.Status,
	}}

	_, err := collectionContract.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (c *contractRepository) GetByID(ctx context.Context, id primitive.ObjectID) (contractsdomain.Contract, error) {
	collectionContract := c.database.Collection(c.collectionContract)

	var contract contractsdomain.Contract
	filter := bson.M{"_id": id}
	if err := collectionContract.FindOne(ctx, filter).Decode(&contract); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return contractsdomain.Contract{}, nil
		}

		return contractsdomain.Contract{}, err
	}

	return contract, nil
}

func (c *contractRepository) GetByEmployeeID(ctx context.Context, employeeID primitive.ObjectID) (contractsdomain.Contract, error) {
	collectionContract := c.database.Collection(c.collectionContract)

	var contract contractsdomain.Contract
	filter := bson.M{"employee_id": employeeID}
	if err := collectionContract.FindOne(ctx, filter).Decode(&contract); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return contractsdomain.Contract{}, nil
		}
		return contractsdomain.Contract{}, err
	}

	return contract, nil
}

func (c *contractRepository) GetAll(ctx context.Context) ([]contractsdomain.Contract, error) {
	collectionContract := c.database.Collection(c.collectionContract)

	filter := bson.M{}
	cursor, err := collectionContract.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	var contracts []contractsdomain.Contract
	contracts = make([]contractsdomain.Contract, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var contract contractsdomain.Contract
		if err = cursor.Decode(contract); err != nil {
			return nil, err
		}

		contracts = append(contracts, contract)
	}

	// Check for any errors encountered during iteration
	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return contracts, nil
}
