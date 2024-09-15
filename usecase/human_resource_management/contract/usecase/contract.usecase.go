package contract_usecase

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	contractsdomain "shop_erp_mono/domain/human_resource_management/contracts"
	employeesdomain "shop_erp_mono/domain/human_resource_management/employees"
	"shop_erp_mono/usecase/human_resource_management/contract/validate"
	"time"
)

type contractUseCase struct {
	contextTimeout     time.Duration
	contractRepository contractsdomain.IContractsRepository
	employeeRepository employeesdomain.IEmployeeRepository
}

func NewContractUseCase(contextTimeout time.Duration, contractRepository contractsdomain.IContractsRepository, employeeRepository employeesdomain.IEmployeeRepository) contractsdomain.IContractsUseCase {
	return &contractUseCase{contextTimeout: contextTimeout, contractRepository: contractRepository, employeeRepository: employeeRepository}
}

func (c *contractUseCase) CreateOne(ctx context.Context, input *contractsdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	if err := validate.ValidateContract(input); err != nil {
		return err
	}

	employeeData, err := c.employeeRepository.GetOneByEmail(ctx, input.EmployeeEmail)
	if err != nil {
		return err
	}

	if input.Status == "" {
		input.Status = "Active"
	}

	contract := contractsdomain.Contract{
		ID:           primitive.NewObjectID(),
		EmployeeID:   employeeData.ID,
		ContractType: input.ContractType,
		StartDate:    input.StartDate,
		EndDate:      input.EndDate,
		Salary:       input.Salary,
		Status:       input.Status,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	return c.contractRepository.CreateOne(ctx, &contract)
}

func (c *contractUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	contractID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return c.contractRepository.DeleteOne(ctx, contractID)
}

func (c *contractUseCase) UpdateOne(ctx context.Context, id string, input *contractsdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	if err := validate.ValidateContract(input); err != nil {
		return err
	}

	contractID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	employeeData, err := c.employeeRepository.GetOneByEmail(ctx, input.EmployeeEmail)
	if err != nil {
		return err
	}

	if input.Status == "" {
		input.Status = "Active"
	}

	contract := contractsdomain.Contract{
		EmployeeID:   employeeData.ID,
		ContractType: input.ContractType,
		StartDate:    input.StartDate,
		EndDate:      input.EndDate,
		Salary:       input.Salary,
		Status:       input.Status,
	}

	return c.contractRepository.UpdateOne(ctx, contractID, &contract)
}

func (c *contractUseCase) GetOneByID(ctx context.Context, id string) (contractsdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	contractID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return contractsdomain.Output{}, err
	}

	if contractID == primitive.NilObjectID {
		return contractsdomain.Output{}, errors.New("id do not nil")
	}

	contractData, err := c.contractRepository.GetOneByID(ctx, contractID)
	if err != nil {
		return contractsdomain.Output{}, err
	}

	employeeData, err := c.employeeRepository.GetOneByID(ctx, contractData.EmployeeID)
	if err != nil {
		return contractsdomain.Output{}, err
	}

	output := contractsdomain.Output{
		Contract: contractData,
		Employee: employeeData,
	}

	return output, nil
}

func (c *contractUseCase) GetOneByEmail(ctx context.Context, email string) (contractsdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	if err := validate.IsNilEmail(email); err != nil {
		return contractsdomain.Output{}, err
	}

	employeeData, err := c.employeeRepository.GetOneByEmail(ctx, email)
	if err != nil {
		return contractsdomain.Output{}, err
	}

	contractData, err := c.contractRepository.GetOneByEmployeeID(ctx, employeeData.ID)
	if err != nil {
		return contractsdomain.Output{}, err
	}

	output := contractsdomain.Output{
		Contract: contractData,
		Employee: employeeData,
	}

	return output, nil
}

func (c *contractUseCase) GetAll(ctx context.Context) ([]contractsdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	contractsData, err := c.contractRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var outputs []contractsdomain.Output
	outputs = make([]contractsdomain.Output, 0, len(contractsData))
	for _, contractData := range contractsData {
		employeeData, err := c.employeeRepository.GetOneByID(ctx, contractData.EmployeeID)
		if err != nil {
			return nil, err
		}

		output := contractsdomain.Output{
			Contract: contractData,
			Employee: employeeData,
		}

		outputs = append(outputs, output)
	}

	return outputs, nil
}
