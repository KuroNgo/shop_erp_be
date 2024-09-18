package payment_usecase

import (
	"context"
	payments_domain "shop_erp_mono/domain/sales_and_distribution_management/payments"
	"time"
)

type paymentUseCase struct {
	contextTimeout    time.Duration
	paymentRepository payments_domain.IPaymentRepository
}

func NewPaymentUseCase(contextTimeout time.Duration, paymentRepository payments_domain.IPaymentRepository) payments_domain.IPaymentUseCase {
	return &paymentUseCase{contextTimeout: contextTimeout, paymentRepository: paymentRepository}
}

func (p *paymentUseCase) CreateOne(ctx context.Context, input *payments_domain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (p *paymentUseCase) GetByID(ctx context.Context, id string) (*payments_domain.PaymentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p *paymentUseCase) GetByOrderID(ctx context.Context, orderID string) ([]payments_domain.PaymentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p *paymentUseCase) GetByStatus(ctx context.Context, status string) ([]payments_domain.PaymentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p *paymentUseCase) UpdateOne(ctx context.Context, id string, input *payments_domain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (p *paymentUseCase) DeleteOne(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (p *paymentUseCase) List(ctx context.Context) ([]payments_domain.PaymentResponse, error) {
	//TODO implement me
	panic("implement me")
}
