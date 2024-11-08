package customer_usecase

import (
	"context"
	"encoding/json"
	"github.com/allegro/bigcache/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	customerdomain "shop_erp_mono/internal/domain/sales_and_distribution_management/customer"
	"shop_erp_mono/internal/usecase/sales_and_distribution_management/customer/validate"
	"time"
)

type customerUseCase struct {
	contextTimeout     time.Duration
	customerRepository customerdomain.ICustomerRepository
	cache              *bigcache.BigCache
}

func NewCustomerUseCase(contextTimeout time.Duration, customerRepository customerdomain.ICustomerRepository, cacheTTL time.Duration) customerdomain.ICustomerUseCase {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(cacheTTL))
	if err != nil {
		return nil
	}
	return &customerUseCase{contextTimeout: contextTimeout, cache: cache, customerRepository: customerRepository}
}

func (c *customerUseCase) CreateOne(ctx context.Context, input *customerdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	if err := validate.Customer(input); err != nil {
		return err
	}

	customer := &customerdomain.Customer{
		ID:          primitive.NewObjectID(),
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		Address:     input.Address,
		City:        input.City,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	_ = c.cache.Delete("customers")

	return c.customerRepository.CreateOne(ctx, customer)
}

func (c *customerUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	customerID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_ = c.cache.Delete(id)
	_ = c.cache.Delete("customers")

	return c.customerRepository.DeleteOne(ctx, customerID)
}

func (c *customerUseCase) UpdateOne(ctx context.Context, id string, input *customerdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	customerID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	if err = validate.Customer(input); err != nil {
		return err
	}

	customer := &customerdomain.Customer{
		ID:          customerID,
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		Address:     input.Address,
		City:        input.City,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	_ = c.cache.Delete(id)
	_ = c.cache.Delete("customers")

	return c.customerRepository.UpdateOne(ctx, customer)
}

func (c *customerUseCase) GetOneByID(ctx context.Context, id string) (*customerdomain.CustomerResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	data, err := c.cache.Get(id)
	if err != nil {
		return nil, err
	}
	if data != nil {
		var responseData *customerdomain.CustomerResponse
		err = json.Unmarshal(data, responseData)
		if err != nil {
			return nil, err
		}
		return responseData, nil
	}

	customerID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	customerData, err := c.customerRepository.GetOneByID(ctx, customerID)
	if err != nil {
		return nil, err
	}

	response := &customerdomain.CustomerResponse{
		Customer: *customerData,
	}

	responseData, err := json.Marshal(response)
	err = c.cache.Set(id, responseData)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *customerUseCase) GetOneByName(ctx context.Context, name string) (*customerdomain.CustomerResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	data, err := c.cache.Get(name)
	if err != nil {
		return nil, err
	}

	if data != nil {
		var response *customerdomain.CustomerResponse
		err = json.Unmarshal(data, response)
		return response, nil
	}

	customerData, err := c.customerRepository.GetOneByName(ctx, name)
	if err != nil {
		return nil, err
	}

	response := &customerdomain.CustomerResponse{
		Customer: *customerData,
	}

	responseData, err := json.Marshal(response)
	err = c.cache.Set(name, responseData)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *customerUseCase) GetAll(ctx context.Context) ([]customerdomain.CustomerResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	data, err := c.cache.Get("customers")
	if err != nil {
		return nil, err
	}

	if data != nil {
		var response []customerdomain.CustomerResponse
		err = json.Unmarshal(data, &response)
		return response, nil
	}

	customerData, err := c.customerRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []customerdomain.CustomerResponse
	responses = make([]customerdomain.CustomerResponse, 0, len(customerData))
	for _, customer := range customerData {
		response := customerdomain.CustomerResponse{
			Customer: customer,
		}

		responses = append(responses, response)
	}

	responsesData, err := json.Marshal(responses)
	err = c.cache.Set("customers", responsesData)
	if err != nil {
		return nil, err
	}
	return responses, nil
}
