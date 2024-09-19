package payments_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	payment_domain "shop_erp_mono/domain/accounting_management/payments"
	"time"
)

type paymentsRepository struct {
	database           *mongo.Database
	paymentsCollection string
}

func NewPaymentsRepository(database *mongo.Database, paymentsCollection string) payment_domain.IPaymentsRepository {
	return &paymentsRepository{database: database, paymentsCollection: paymentsCollection}
}

func (p *paymentsRepository) CreatePayment(ctx context.Context, payment *payment_domain.Payments) error {
	//TODO implement me
	panic("implement me")
}

func (p *paymentsRepository) GetPaymentByID(ctx context.Context, id primitive.ObjectID) (payment_domain.Payments, error) {
	//TODO implement me
	panic("implement me")
}

func (p *paymentsRepository) GetPaymentsByInvoiceID(ctx context.Context, invoiceID primitive.ObjectID) ([]payment_domain.Payments, error) {
	//TODO implement me
	panic("implement me")
}

func (p *paymentsRepository) GetPaymentsByAccountID(ctx context.Context, accountID primitive.ObjectID) ([]payment_domain.Payments, error) {
	//TODO implement me
	panic("implement me")
}

func (p *paymentsRepository) UpdatePayment(ctx context.Context, payment *payment_domain.Payments) error {
	//TODO implement me
	panic("implement me")
}

func (p *paymentsRepository) DeletePayment(ctx context.Context, id primitive.ObjectID) error {
	//TODO implement me
	panic("implement me")
}

func (p *paymentsRepository) ListPayments(ctx context.Context) ([]payment_domain.Payments, error) {
	//TODO implement me
	panic("implement me")
}

func (p *paymentsRepository) GetPaymentsByDateRange(ctx context.Context, startDate, endDate time.Time) ([]payment_domain.Payments, error) {
	//TODO implement me
	panic("implement me")
}

func (p *paymentsRepository) GetTotalPaymentsAmount(ctx context.Context) (int32, error) {
	//TODO implement me
	panic("implement me")
}
