package payment_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	paymentsdomain "shop_erp_mono/internal/domain/sales_and_distribution_management/payments"
	"time"
)

type paymentRepository struct {
	database          *mongo.Database
	paymentCollection string
}

func NewPaymentRepository(database *mongo.Database, paymentCollection string) paymentsdomain.IPaymentRepository {
	return &paymentRepository{database: database, paymentCollection: paymentCollection}
}

func (p *paymentRepository) CreateOne(ctx context.Context, payment paymentsdomain.Payment) error {
	paymentCollection := p.database.Collection(p.paymentCollection)

	_, err := paymentCollection.InsertOne(ctx, payment)
	if err != nil {
		return err
	}

	return nil
}

func (p *paymentRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*paymentsdomain.Payment, error) {
	paymentCollection := p.database.Collection(p.paymentCollection)

	filter := bson.M{"_id": id}
	var payment paymentsdomain.Payment
	err := paymentCollection.FindOne(ctx, filter).Decode(&payment)
	if err != nil {
		return nil, err
	}

	return &payment, nil
}

func (p *paymentRepository) GetByOrderID(ctx context.Context, orderID primitive.ObjectID) ([]paymentsdomain.Payment, error) {
	paymentCollection := p.database.Collection(p.paymentCollection)

	filter := bson.M{"order_id": orderID}
	cursor, err := paymentCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var payments []paymentsdomain.Payment
	payments = make([]paymentsdomain.Payment, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var payment paymentsdomain.Payment
		if err = cursor.Decode(&payment); err != nil {
			return nil, err
		}

		payments = append(payments, payment)
	}

	return payments, nil
}

func (p *paymentRepository) GetByStatus(ctx context.Context, status string) ([]paymentsdomain.Payment, error) {
	paymentCollection := p.database.Collection(p.paymentCollection)

	filter := bson.M{"status": status}
	cursor, err := paymentCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var payments []paymentsdomain.Payment
	payments = make([]paymentsdomain.Payment, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var payment paymentsdomain.Payment
		if err = cursor.Decode(&payment); err != nil {
			return nil, err
		}

		payments = append(payments, payment)
	}

	return payments, nil
}

func (p *paymentRepository) UpdateOne(ctx context.Context, payment paymentsdomain.Payment) error {
	paymentCollection := p.database.Collection(p.paymentCollection)

	filter := bson.M{"_id": payment.ID}
	update := bson.M{"$set": bson.M{
		"order_id":       payment.OrderID,
		"payment_date":   payment.PaymentDate,
		"payment_method": payment.PaymentMethod,
		"amount_paid":    payment.AmountPaid,
		"status":         payment.Status,
		"updated_at":     time.Now(),
	}}
	_, err := paymentCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (p *paymentRepository) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	paymentCollection := p.database.Collection(p.paymentCollection)

	filter := bson.M{"_id": id}
	_, err := paymentCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (p *paymentRepository) GetAll(ctx context.Context) ([]paymentsdomain.Payment, error) {
	paymentCollection := p.database.Collection(p.paymentCollection)

	filter := bson.M{}
	cursor, err := paymentCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var payments []paymentsdomain.Payment
	payments = make([]paymentsdomain.Payment, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var payment paymentsdomain.Payment
		if err = cursor.Decode(&payment); err != nil {
			return nil, err
		}

		payments = append(payments, payment)
	}

	return payments, nil
}
