package contract_repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	contractsdomain "shop_erp_mono/domain/human_resource_management/contracts"
	employeesdomain "shop_erp_mono/domain/human_resource_management/employees"
	"shop_erp_mono/repository/human_resource_management/contracts/validate"
	"time"
)

type contractRepository struct {
	database           *mongo.Database
	collectionContract string
	collectionEmployee string
}

func NewContractRepository(db *mongo.Database, collectionContract string, collectionEmployee string) contractsdomain.IContractsRepository {
	return &contractRepository{database: db, collectionContract: collectionContract, collectionEmployee: collectionEmployee}
}

func (c contractRepository) CreateOne(ctx context.Context, input *contractsdomain.Input) error {
	collectionContract := c.database.Collection(c.collectionContract)
	collectionEmployee := c.database.Collection(c.collectionEmployee)

	if err := validate.IsNilContract(input); err != nil {
		return err
	}

	filterEmployee := bson.M{"email": input.EmployeeEmail}
	var employee employeesdomain.Employee
	if err := collectionEmployee.FindOne(ctx, filterEmployee).Decode(&employee); err != nil {
		return err
	}

	if input.Status == "" {
		input.Status = "Active"
	}

	contract := contractsdomain.Contract{
		ID:           primitive.NewObjectID(),
		EmployeeID:   employee.ID,
		ContractType: input.ContractType,
		StartDate:    input.StartDate,
		EndDate:      input.EndDate,
		Salary:       input.Salary,
		Status:       input.Status,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	_, err := collectionContract.InsertOne(ctx, contract)
	if err != nil {
		return err
	}

	return nil
}

func (c contractRepository) DeleteOne(ctx context.Context, id string) error {
	collectionContract := c.database.Collection(c.collectionContract)

	contractID, _ := primitive.ObjectIDFromHex(id)
	if contractID == primitive.NilObjectID {
		return errors.New("the contractID do not nil")
	}

	filter := bson.M{"_id": contractID}
	_, err := collectionContract.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (c contractRepository) UpdateOne(ctx context.Context, id string, input *contractsdomain.Input) error {
	collectionContract := c.database.Collection(c.collectionContract)
	collectionEmployee := c.database.Collection(c.collectionEmployee)

	contractID, _ := primitive.ObjectIDFromHex(id)
	if contractID == primitive.NilObjectID {
		return errors.New("id do not nil")
	}

	if err := validate.IsNilContract(input); err != nil {
		return err
	}

	filterEmployee := bson.M{"email": input.EmployeeEmail}
	var employee employeesdomain.Employee
	if err := collectionEmployee.FindOne(ctx, filterEmployee).Decode(&employee); err != nil {
		return err
	}

	if input.Status == "" {
		input.Status = "Active"
	}

	filter := bson.M{"_id": contractID}
	update := bson.M{"$set": bson.M{
		"employee_id":   employee.ID,
		"contract_type": input.ContractType,
		"start_date":    input.StartDate,
		"end_date":      input.EndDate,
		"salary":        input.Salary,
		"status":        input.Status,
	}}

	_, err := collectionContract.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (c contractRepository) GetOneByID(ctx context.Context, id string) (contractsdomain.Output, error) {
	collectionContract := c.database.Collection(c.collectionContract)
	collectionEmployee := c.database.Collection(c.collectionEmployee)

	contractID, _ := primitive.ObjectIDFromHex(id)
	if contractID == primitive.NilObjectID {
		return contractsdomain.Output{}, errors.New("id do not nil")
	}

	var contract contractsdomain.Contract
	filter := bson.M{"_id": contractID}
	if err := collectionContract.FindOne(ctx, filter).Decode(&contract); err != nil {
		return contractsdomain.Output{}, err
	}

	var employee employeesdomain.Employee
	filterEmployee := bson.M{"_id": contract.EmployeeID}
	if err := collectionEmployee.FindOne(ctx, filterEmployee).Decode(&employee); err != nil {
		return contractsdomain.Output{}, err
	}

	output := contractsdomain.Output{
		Contract: contract,
		Employee: employee.LastName,
	}
	return output, nil
}

func (c contractRepository) GetOneByEmail(ctx context.Context, email string) (contractsdomain.Output, error) {
	collectionContract := c.database.Collection(c.collectionContract)
	collectionEmployee := c.database.Collection(c.collectionEmployee)

	if err := validate.IsNilEmail(email); err != nil {
		return contractsdomain.Output{}, err
	}

	var employee employeesdomain.Employee
	filterEmployee := bson.M{"email": email}
	if err := collectionEmployee.FindOne(ctx, filterEmployee).Decode(&employee); err != nil {
		return contractsdomain.Output{}, err
	}

	var contract contractsdomain.Contract
	filter := bson.M{"employee_id": employee}
	if err := collectionContract.FindOne(ctx, filter).Decode(&contract); err != nil {
		return contractsdomain.Output{}, err
	}

	output := contractsdomain.Output{
		Contract: contract,
		Employee: employee.LastName,
	}
	return output, nil
}

func (c contractRepository) GetAll(ctx context.Context) ([]contractsdomain.Output, error) {
	collectionContract := c.database.Collection(c.collectionContract)
	collectionEmployee := c.database.Collection(c.collectionEmployee)

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

	var contracts []contractsdomain.Output
	contracts = make([]contractsdomain.Output, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var contract contractsdomain.Contract
		if err = cursor.Decode(contract); err != nil {
			return nil, err
		}

		var employee employeesdomain.Employee
		filterEmployee := bson.M{"_id": contract.EmployeeID}
		err = collectionEmployee.FindOne(ctx, filterEmployee).Decode(&employee)
		if err != nil {
			return nil, err
		}

		output := contractsdomain.Output{
			Contract: contract,
			Employee: employee.LastName,
		}

		contracts = append(contracts, output)
	}

	return contracts, nil
}
