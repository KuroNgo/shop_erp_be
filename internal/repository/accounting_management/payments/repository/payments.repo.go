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

func (p *paymentsRepository) CreateOne(ctx context.Context, payment *payment_domain.Payments) error {
	//TODO implement me
	panic("implement me")
}

func (p *paymentsRepository) GetByID(ctx context.Context, id primitive.ObjectID) (payment_domain.Payments, error) {
	//TODO implement me
	panic("implement me")
}

func (p *paymentsRepository) GetByInvoiceID(ctx context.Context, invoiceID primitive.ObjectID) ([]payment_domain.Payments, error) {
	//TODO implement me
	panic("implement me")
}

func (p *paymentsRepository) GetByAccountID(ctx context.Context, accountID primitive.ObjectID) ([]payment_domain.Payments, error) {
	//TODO implement me
	panic("implement me")
}

func (p *paymentsRepository) UpdateOne(ctx context.Context, payment *payment_domain.Payments) error {
	//TODO implement me
	panic("implement me")
}

func (p *paymentsRepository) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	//TODO implement me
	panic("implement me")
}

func (p *paymentsRepository) GetAll(ctx context.Context) ([]payment_domain.Payments, error) {
	//TODO implement me
	panic("implement me")
}

func (p *paymentsRepository) GetByDateRange(ctx context.Context, startDate, endDate time.Time) ([]payment_domain.Payments, error) {
	//TODO implement me
	panic("implement me")
}

func (p *paymentsRepository) GetTotalPaymentsAmount(ctx context.Context) (int32, error) {
	//TODO implement me
	panic("implement me")
}
