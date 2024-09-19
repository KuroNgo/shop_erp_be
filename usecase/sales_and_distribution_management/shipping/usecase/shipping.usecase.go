package shipping_usecase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	saleordersdomain "shop_erp_mono/domain/sales_and_distribution_management/sale_orders"
	shippingdomain "shop_erp_mono/domain/sales_and_distribution_management/shipping"
	"time"
)

type shippingUseCase struct {
	contextTimeout      time.Duration
	shippingRepository  shippingdomain.IShippingRepository
	saleOrderRepository saleordersdomain.ISalesOrderRepository
}

func NewShippingUseCase(contextTimeout time.Duration, shippingRepository shippingdomain.IShippingRepository, saleOrderRepository saleordersdomain.ISalesOrderRepository) shippingdomain.IShippingUseCase {
	return &shippingUseCase{contextTimeout: contextTimeout, shippingRepository: shippingRepository, saleOrderRepository: saleOrderRepository}
}

func (s *shippingUseCase) CreateOne(ctx context.Context, input *shippingdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	orderID, err := primitive.ObjectIDFromHex(input.OrderID)
	if err != nil {
		return err
	}

	shipping := shippingdomain.Shipping{
		ID:                primitive.NewObjectID(),
		OrderID:           orderID,
		ShippingMethod:    input.ShippingMethod,
		ShippingDate:      input.ShippingDate,
		EstimatedDelivery: input.EstimatedDelivery,
		ActualDelivery:    input.ActualDelivery,
		TrackingNumber:    input.TrackingNumber,
		Status:            input.Status,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	return s.shippingRepository.CreateOne(ctx, shipping)
}

func (s *shippingUseCase) GetByID(ctx context.Context, id string) (*shippingdomain.ShippingResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	shippingID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	shippingData, err := s.shippingRepository.GetByID(ctx, shippingID)
	if err != nil {
		return nil, err
	}

	orderData, err := s.saleOrderRepository.GetByID(ctx, shippingData.OrderID)
	if err != nil {
		return nil, err
	}
	response := &shippingdomain.ShippingResponse{
		Shipping: *shippingData,
		Order:    *orderData,
	}

	return response, nil
}

func (s *shippingUseCase) GetByOrderID(ctx context.Context, orderID string) (*shippingdomain.ShippingResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	idOrder, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		return nil, err
	}

	shippingData, err := s.shippingRepository.GetByOrderID(ctx, idOrder)
	if err != nil {
		return nil, err
	}

	orderData, err := s.saleOrderRepository.GetByID(ctx, shippingData.OrderID)
	if err != nil {
		return nil, err
	}
	response := &shippingdomain.ShippingResponse{
		Shipping: *shippingData,
		Order:    *orderData,
	}

	return response, nil
}

func (s *shippingUseCase) UpdateOne(ctx context.Context, id string, updatedShipping *shippingdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	shippingID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	idOrder, err := primitive.ObjectIDFromHex(updatedShipping.OrderID)
	if err != nil {
		return err
	}

	shipping := shippingdomain.Shipping{
		ID:                shippingID,
		OrderID:           idOrder,
		ShippingMethod:    updatedShipping.ShippingMethod,
		ShippingDate:      updatedShipping.ShippingDate,
		EstimatedDelivery: updatedShipping.EstimatedDelivery,
		ActualDelivery:    updatedShipping.ActualDelivery,
		TrackingNumber:    updatedShipping.TrackingNumber,
		Status:            updatedShipping.Status,
		UpdatedAt:         time.Now(),
	}

	return s.shippingRepository.UpdateOne(ctx, shipping)
}

func (s *shippingUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	shippingID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return s.shippingRepository.DeleteOne(ctx, shippingID)
}

func (s *shippingUseCase) List(ctx context.Context) ([]shippingdomain.ShippingResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	shippingData, err := s.shippingRepository.List(ctx)
	if err != nil {
		return nil, err
	}

	var responses []shippingdomain.ShippingResponse
	responses = make([]shippingdomain.ShippingResponse, 0, len(shippingData))
	for _, shipping := range shippingData {
		orderData, err := s.saleOrderRepository.GetByID(ctx, shipping.OrderID)
		if err != nil {
			return nil, err
		}

		response := shippingdomain.ShippingResponse{
			Shipping: shipping,
			Order:    *orderData,
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (s *shippingUseCase) GetByStatus(ctx context.Context, status string) ([]shippingdomain.ShippingResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	shippingData, err := s.shippingRepository.GetByStatus(ctx, status)
	if err != nil {
		return nil, err
	}

	var responses []shippingdomain.ShippingResponse
	responses = make([]shippingdomain.ShippingResponse, 0, len(shippingData))
	for _, shipping := range shippingData {
		orderData, err := s.saleOrderRepository.GetByID(ctx, shipping.OrderID)
		if err != nil {
			return nil, err
		}

		response := shippingdomain.ShippingResponse{
			Shipping: shipping,
			Order:    *orderData,
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (s *shippingUseCase) UpdateDeliveryStatus(ctx context.Context, id string, status string, actualDelivery *time.Time) error {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	shippingID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	shipping := shippingdomain.Shipping{
		ID:             shippingID,
		ActualDelivery: actualDelivery,
		Status:         status,
		UpdatedAt:      time.Now(),
	}

	return s.shippingRepository.UpdateOne(ctx, shipping)
}
