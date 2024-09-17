package payments_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IPaymentRepository interface {
	CreateOne(ctx context.Context, payment Payment) error
	GetByID(ctx context.Context, id primitive.ObjectID) (*Payment, error)
	GetByOrderID(ctx context.Context, orderID primitive.ObjectID) ([]Payment, error)
	GetByStatus(ctx context.Context, status string) ([]Payment, error)
	UpdateOne(ctx context.Context, id primitive.ObjectID, updatedPayment Input) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	List(ctx context.Context) ([]Payment, error)
}

type IPaymentUseCase interface {
	CreateOne(ctx context.Context, input *Input) error
	GetByID(ctx context.Context, id string) (*PaymentResponse, error)
	GetByOrderID(ctx context.Context, orderID string) ([]PaymentResponse, error)
	GetByStatus(ctx context.Context, status string) ([]PaymentResponse, error)
	UpdateOne(ctx context.Context, id string, input *Input) error
	DeleteOne(ctx context.Context, id string) error
	List(ctx context.Context) ([]PaymentResponse, error)
}
