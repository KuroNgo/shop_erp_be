package customer_usecase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	customerdomain "shop_erp_mono/domain/sales_and_distribution_management/customer"
	"shop_erp_mono/usecase/sales_and_distribution_management/customer/validate"
	"time"
)

type customerUseCase struct {
	contextTimeout     time.Duration
	customerRepository customerdomain.ICustomerRepository
}

func NewCustomerUseCase(contextTimeout time.Duration, customerRepository customerdomain.ICustomerRepository) customerdomain.ICustomerUseCase {
	return &customerUseCase{contextTimeout: contextTimeout, customerRepository: customerRepository}
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

	return c.customerRepository.CreateOne(ctx, customer)
}

func (c *customerUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	customerID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

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

	return c.customerRepository.UpdateOne(ctx, customer)

}

func (c *customerUseCase) GetOneByID(ctx context.Context, id string) (*customerdomain.CustomerResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

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

	return response, nil
}

func (c *customerUseCase) GetOneByName(ctx context.Context, name string) (*customerdomain.CustomerResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	customerData, err := c.customerRepository.GetOneByName(ctx, name)
	if err != nil {
		return nil, err
	}

	response := &customerdomain.CustomerResponse{
		Customer: *customerData,
	}

	return response, nil
}

func (c *customerUseCase) GetAll(ctx context.Context) ([]customerdomain.CustomerResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

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

	return responses, nil
}
