package contract_usecase

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/allegro/bigcache/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	contractsdomain "shop_erp_mono/internal/domain/human_resource_management/contracts"
	employeesdomain "shop_erp_mono/internal/domain/human_resource_management/employees"
	"shop_erp_mono/internal/usecase/human_resource_management/contract/validate"
	"time"
)

type contractUseCase struct {
	contextTimeout     time.Duration
	contractRepository contractsdomain.IContractsRepository
	employeeRepository employeesdomain.IEmployeeRepository
	cache              *bigcache.BigCache
}

func NewContractUseCase(contextTimeout time.Duration, contractRepository contractsdomain.IContractsRepository,
	employeeRepository employeesdomain.IEmployeeRepository, cacheTTL time.Duration) contractsdomain.IContractsUseCase {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(cacheTTL))
	if err != nil {
		return nil
	}
	return &contractUseCase{contextTimeout: contextTimeout, cache: cache, contractRepository: contractRepository, employeeRepository: employeeRepository}
}

func (c *contractUseCase) CreateOne(ctx context.Context, input *contractsdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	if err := validate.Contract(input); err != nil {
		return err
	}

	employeeData, err := c.employeeRepository.GetByEmail(ctx, input.EmployeeEmail)
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

	err = c.contractRepository.CreateOne(ctx, &contract)
	if err != nil {
		return err
	}

	if err = c.cache.Delete("contracts"); err != nil {
		log.Printf("failed to delete contracts cache: %v", err)
	}

	return nil
}

func (c *contractUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	contractID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	err = c.contractRepository.DeleteOne(ctx, contractID)
	if err != nil {
		return err
	}

	if err = c.cache.Delete(id); err != nil {
		log.Printf("failed to delete contract's id cache: %v", err)
	}
	if err = c.cache.Delete("contracts"); err != nil {
		log.Printf("failed to delete contracts cache: %v", err)
	}

	return nil
}

func (c *contractUseCase) UpdateOne(ctx context.Context, id string, input *contractsdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	if err := validate.Contract(input); err != nil {
		return err
	}

	contractID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	employeeData, err := c.employeeRepository.GetByEmail(ctx, input.EmployeeEmail)
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

	err = c.contractRepository.UpdateOne(ctx, contractID, &contract)
	if err != nil {
		return err
	}

	if err = c.cache.Delete(id); err != nil {
		log.Printf("failed to delete contract's id cache: %v", err)
	}
	if err = c.cache.Delete("contracts"); err != nil {
		log.Printf("failed to delete contracts cache: %v", err)
	}

	return nil
}

func (c *contractUseCase) GetByID(ctx context.Context, id string) (contractsdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	data, err := c.cache.Get(id)
	if err != nil {
		log.Printf("failed to get contract's id cache: %v", err)
	}
	if data != nil {
		var response contractsdomain.Output
		err := json.Unmarshal(data, &response)
		if err != nil {
			return contractsdomain.Output{}, err
		}
		return response, nil
	}

	contractID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return contractsdomain.Output{}, err
	}

	if contractID == primitive.NilObjectID {
		return contractsdomain.Output{}, errors.New("id do not nil")
	}

	contractData, err := c.contractRepository.GetByID(ctx, contractID)
	if err != nil {
		return contractsdomain.Output{}, err
	}

	employeeData, err := c.employeeRepository.GetByID(ctx, contractData.EmployeeID)
	if err != nil {
		return contractsdomain.Output{}, err
	}

	output := contractsdomain.Output{
		Contract: contractData,
		Employee: employeeData,
	}

	data, err = json.Marshal(output)
	if err != nil {
		return contractsdomain.Output{}, err
	}

	err = c.cache.Set(id, data)
	if err != nil {
		log.Printf("failed to delete id cache: %v", err)
	}
	return output, nil
}

func (c *contractUseCase) GetByEmail(ctx context.Context, email string) (contractsdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	data, err := c.cache.Get(email)
	if err != nil {
		log.Printf("failed to delete contract's email cache: %v", err)
	}
	if data != nil {
		var response contractsdomain.Output
		err := json.Unmarshal(data, &response)
		if err != nil {
			return contractsdomain.Output{}, err
		}
		return response, nil
	}

	if err := validate.IsNilEmail(email); err != nil {
		return contractsdomain.Output{}, err
	}

	employeeData, err := c.employeeRepository.GetByEmail(ctx, email)
	if err != nil {
		return contractsdomain.Output{}, err
	}

	contractData, err := c.contractRepository.GetByEmployeeID(ctx, employeeData.ID)
	if err != nil {
		return contractsdomain.Output{}, err
	}

	output := contractsdomain.Output{
		Contract: contractData,
		Employee: employeeData,
	}

	data, err = json.Marshal(output)
	if err != nil {
		return contractsdomain.Output{}, err
	}

	err = c.cache.Set(email, data)
	if err != nil {
		log.Printf("failed to delete contract's email cache: %v", err)
	}

	return output, nil
}

func (c *contractUseCase) GetAll(ctx context.Context) ([]contractsdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	data, err := c.cache.Get("contracts")
	if err != nil {
		log.Printf("failed to delete contracts cache: %v", err)
	}
	if data != nil {
		var response []contractsdomain.Output
		err := json.Unmarshal(data, &response)
		if err != nil {
			return nil, err
		}
		return response, nil
	}

	contractsData, err := c.contractRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var outputs []contractsdomain.Output
	outputs = make([]contractsdomain.Output, 0, len(contractsData))
	for _, contractData := range contractsData {
		employeeData, err := c.employeeRepository.GetByID(ctx, contractData.EmployeeID)
		if err != nil {
			return nil, err
		}

		output := contractsdomain.Output{
			Contract: contractData,
			Employee: employeeData,
		}

		outputs = append(outputs, output)
	}

	data, err = json.Marshal(outputs)
	if err != nil {
		return nil, err
	}

	err = c.cache.Set("contracts", data)
	if err != nil {
		log.Printf("failed to delete contracts cache: %v", err)
	}

	return outputs, nil
}
