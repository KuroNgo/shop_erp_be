package payment_usecase

import (
	"context"
	"github.com/allegro/bigcache/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	paymentsdomain "shop_erp_mono/internal/domain/sales_and_distribution_management/payments"
	saleordersdomain "shop_erp_mono/internal/domain/sales_and_distribution_management/sale_orders"
	"shop_erp_mono/internal/usecase/sales_and_distribution_management/payment/validate"
	"time"
)

type paymentUseCase struct {
	contextTimeout       time.Duration
	paymentRepository    paymentsdomain.IPaymentRepository
	salesOrderRepository saleordersdomain.ISalesOrderRepository
	cache                *bigcache.BigCache
}

func NewPaymentUseCase(contextTimeout time.Duration, paymentRepository paymentsdomain.IPaymentRepository,
	salesOrderRepository saleordersdomain.ISalesOrderRepository, cacheTTL time.Duration) paymentsdomain.IPaymentUseCase {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(cacheTTL))
	if err != nil {
		return nil
	}
	return &paymentUseCase{contextTimeout: contextTimeout, cache: cache, paymentRepository: paymentRepository, salesOrderRepository: salesOrderRepository}
}

func (p *paymentUseCase) CreateOne(ctx context.Context, input *paymentsdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	if err := validate.Payment(input); err != nil {
		return err
	}

	orderID, err := primitive.ObjectIDFromHex(input.OrderID)
	if err != nil {
		return err
	}

	payment := paymentsdomain.Payment{
		ID:            primitive.NewObjectID(),
		OrderID:       orderID,
		PaymentDate:   input.PaymentDate,
		PaymentMethod: input.PaymentMethod,
		AmountPaid:    input.AmountPaid,
		Status:        input.Status,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	return p.paymentRepository.CreateOne(ctx, payment)
}

func (p *paymentUseCase) GetByID(ctx context.Context, id string) (*paymentsdomain.PaymentResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	orderID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	paymentData, err := p.paymentRepository.GetByID(ctx, orderID)
	if err != nil {
		return nil, err
	}

	orderData, err := p.salesOrderRepository.GetByID(ctx, paymentData.OrderID)
	if err != nil {
		return nil, err
	}

	response := &paymentsdomain.PaymentResponse{
		Payment: *paymentData,
		Order:   *orderData,
	}

	return response, nil
}

func (p *paymentUseCase) GetByOrderID(ctx context.Context, orderID string) ([]paymentsdomain.PaymentResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	idOrder, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		return nil, err
	}

	paymentData, err := p.paymentRepository.GetByOrderID(ctx, idOrder)
	if err != nil {
		return nil, err
	}

	var responses []paymentsdomain.PaymentResponse
	responses = make([]paymentsdomain.PaymentResponse, 0, len(responses))
	for _, payment := range paymentData {
		orderData, err := p.salesOrderRepository.GetByID(ctx, payment.OrderID)
		if err != nil {
			return nil, err
		}

		response := paymentsdomain.PaymentResponse{
			Payment: payment,
			Order:   *orderData,
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (p *paymentUseCase) GetByStatus(ctx context.Context, status string) ([]paymentsdomain.PaymentResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	paymentData, err := p.paymentRepository.GetByStatus(ctx, status)
	if err != nil {
		return nil, err
	}

	var responses []paymentsdomain.PaymentResponse
	responses = make([]paymentsdomain.PaymentResponse, 0, len(responses))
	for _, payment := range paymentData {
		orderData, err := p.salesOrderRepository.GetByID(ctx, payment.OrderID)
		if err != nil {
			return nil, err
		}

		response := paymentsdomain.PaymentResponse{
			Payment: payment,
			Order:   *orderData,
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (p *paymentUseCase) UpdateOne(ctx context.Context, id string, input *paymentsdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	if err := validate.Payment(input); err != nil {
		return err
	}

	idPayment, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	idOrder, err := primitive.ObjectIDFromHex(input.OrderID)
	if err != nil {
		return err
	}

	payment := paymentsdomain.Payment{
		ID:            idPayment,
		OrderID:       idOrder,
		PaymentDate:   input.PaymentDate,
		PaymentMethod: input.PaymentMethod,
		AmountPaid:    input.AmountPaid,
		Status:        input.Status,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	return p.paymentRepository.UpdateOne(ctx, payment)
}

func (p *paymentUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	idOrder, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return p.paymentRepository.DeleteOne(ctx, idOrder)
}

func (p *paymentUseCase) GetAll(ctx context.Context) ([]paymentsdomain.PaymentResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	paymentData, err := p.paymentRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []paymentsdomain.PaymentResponse
	responses = make([]paymentsdomain.PaymentResponse, 0, len(responses))
	for _, payment := range paymentData {
		orderData, err := p.salesOrderRepository.GetByID(ctx, payment.OrderID)
		if err != nil {
			return nil, err
		}

		response := paymentsdomain.PaymentResponse{
			Payment: payment,
			Order:   *orderData,
		}

		responses = append(responses, response)
	}

	return responses, nil
}
