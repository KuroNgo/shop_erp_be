package sales_order_usecase

import (
	"context"
	"github.com/allegro/bigcache/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	customerdomain "shop_erp_mono/internal/domain/sales_and_distribution_management/customer"
	saleordersdomain "shop_erp_mono/internal/domain/sales_and_distribution_management/sale_orders"
	"shop_erp_mono/internal/usecase/sales_and_distribution_management/sale_order/validate"
	"time"
)

type saleOrderUseCase struct {
	contextTimeout      time.Duration
	saleOrderRepository saleordersdomain.ISalesOrderRepository
	customerRepository  customerdomain.ICustomerRepository
	cache               *bigcache.BigCache
}

func NewSaleOrderUseCase(contextTimeout time.Duration, saleOrderRepository saleordersdomain.ISalesOrderRepository,
	customerRepository customerdomain.ICustomerRepository, cacheTTL time.Duration) saleordersdomain.ISalesOrderUseCase {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(cacheTTL))
	if err != nil {
		return nil
	}
	return &saleOrderUseCase{contextTimeout: contextTimeout, cache: cache, saleOrderRepository: saleOrderRepository, customerRepository: customerRepository}
}

func (s *saleOrderUseCase) GetByID(ctx context.Context, id string) (*saleordersdomain.SalesOrderResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	saleOrderID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	saleOrderData, err := s.saleOrderRepository.GetByID(ctx, saleOrderID)
	if err != nil {
		return nil, err
	}

	customerData, err := s.customerRepository.GetOneByID(ctx, saleOrderData.ID)
	if err != nil {
		return nil, err
	}

	response := &saleordersdomain.SalesOrderResponse{
		SalesOrder: *saleOrderData,
		Customer:   *customerData,
	}

	return response, nil
}

func (s *saleOrderUseCase) GetByCustomerID(ctx context.Context, customerID string) ([]saleordersdomain.SalesOrderResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	idCustomer, err := primitive.ObjectIDFromHex(customerID)
	if err != nil {
		return nil, err
	}

	salesOrderData, err := s.saleOrderRepository.GetByCustomerID(ctx, idCustomer)
	if err != nil {
		return nil, err
	}

	var responses []saleordersdomain.SalesOrderResponse
	responses = make([]saleordersdomain.SalesOrderResponse, 0, len(salesOrderData))
	for _, salesOrder := range salesOrderData {
		customerData, err := s.customerRepository.GetOneByID(ctx, salesOrder.CustomerID)
		if err != nil {
			return nil, err
		}

		response := saleordersdomain.SalesOrderResponse{
			SalesOrder: salesOrder,
			Customer:   *customerData,
		}

		responses = append(responses, response)
	}

	return responses, nil

}

func (s *saleOrderUseCase) GetByStatus(ctx context.Context, status string) ([]saleordersdomain.SalesOrderResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	salesOrderData, err := s.saleOrderRepository.GetByStatus(ctx, status)
	if err != nil {
		return nil, err
	}

	var responses []saleordersdomain.SalesOrderResponse
	responses = make([]saleordersdomain.SalesOrderResponse, 0, len(salesOrderData))
	for _, salesOrder := range salesOrderData {
		customerData, err := s.customerRepository.GetOneByID(ctx, salesOrder.CustomerID)
		if err != nil {
			return nil, err
		}

		response := saleordersdomain.SalesOrderResponse{
			SalesOrder: salesOrder,
			Customer:   *customerData,
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (s *saleOrderUseCase) GetAll(ctx context.Context) ([]saleordersdomain.SalesOrderResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	salesOrderData, err := s.saleOrderRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []saleordersdomain.SalesOrderResponse
	responses = make([]saleordersdomain.SalesOrderResponse, 0, len(salesOrderData))
	for _, salesOrder := range salesOrderData {
		customerData, err := s.customerRepository.GetOneByID(ctx, salesOrder.CustomerID)
		if err != nil {
			return nil, err
		}

		response := saleordersdomain.SalesOrderResponse{
			SalesOrder: salesOrder,
			Customer:   *customerData,
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (s *saleOrderUseCase) CreateOne(ctx context.Context, input *saleordersdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	if err := validate.SaleOrder(input); err != nil {
		return err
	}

	customerID, err := primitive.ObjectIDFromHex(input.CustomerID)
	if err != nil {
		return err
	}

	saleOrder := saleordersdomain.SalesOrder{
		ID:              primitive.NewObjectID(),
		OrderDate:       input.OrderDate,
		CustomerID:      customerID,
		OrderNumber:     input.OrderNumber,
		ShippingAddress: input.ShippingAddress,
		TotalAmount:     input.TotalAmount,
		Status:          input.Status,
		UpdatedAt:       time.Now(),
	}

	return s.saleOrderRepository.UpdateOne(ctx, saleOrder)
}

func (s *saleOrderUseCase) UpdateOne(ctx context.Context, id string, input *saleordersdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	if err := validate.SaleOrder(input); err != nil {
		return err
	}

	saleOrderID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	customerID, err := primitive.ObjectIDFromHex(input.CustomerID)
	if err != nil {
		return err
	}

	saleOrder := saleordersdomain.SalesOrder{
		ID:              saleOrderID,
		OrderDate:       input.OrderDate,
		CustomerID:      customerID,
		OrderNumber:     input.OrderNumber,
		ShippingAddress: input.ShippingAddress,
		TotalAmount:     input.TotalAmount,
		Status:          input.Status,
		UpdatedAt:       time.Now(),
	}

	return s.saleOrderRepository.UpdateOne(ctx, saleOrder)
}

func (s *saleOrderUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	saleOrderID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return s.saleOrderRepository.DeleteOne(ctx, saleOrderID)
}
