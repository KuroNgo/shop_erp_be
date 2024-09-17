package payment_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	paymentsdomain "shop_erp_mono/domain/sales_and_distribution_management/payments"
)

type paymentRepository struct {
	database          *mongo.Database
	paymentCollection string
}

func NewPaymentRepository(database *mongo.Database, paymentCollection string) paymentsdomain.IPaymentRepository {
	return &paymentRepository{database: database, paymentCollection: paymentCollection}
}

func (p *paymentRepository) CreateOne(ctx context.Context, payment paymentsdomain.Payment) error {
	//TODO implement me
	panic("implement me")
}

func (p *paymentRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*paymentsdomain.Payment, error) {
	//TODO implement me
	panic("implement me")
}

func (p *paymentRepository) GetByOrderID(ctx context.Context, orderID primitive.ObjectID) ([]paymentsdomain.Payment, error) {
	//TODO implement me
	panic("implement me")
}

func (p *paymentRepository) GetByStatus(ctx context.Context, status string) ([]paymentsdomain.Payment, error) {
	//TODO implement me
	panic("implement me")
}

func (p *paymentRepository) UpdateOne(ctx context.Context, id primitive.ObjectID, updatedPayment paymentsdomain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (p *paymentRepository) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	//TODO implement me
	panic("implement me")
}

func (p *paymentRepository) List(ctx context.Context) ([]paymentsdomain.Payment, error) {
	//TODO implement me
	panic("implement me")
}
