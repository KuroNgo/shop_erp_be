package payment_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type IPaymentsRepository interface {
	CreateOne(ctx context.Context, payment *Payments) error
	GetByID(ctx context.Context, id primitive.ObjectID) (Payments, error)
	GetByInvoiceID(ctx context.Context, invoiceID primitive.ObjectID) ([]Payments, error)
	GetByAccountID(ctx context.Context, accountID primitive.ObjectID) ([]Payments, error)
	UpdateOne(ctx context.Context, payment *Payments) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	GetAll(ctx context.Context) ([]Payments, error)
	GetByDateRange(ctx context.Context, startDate, endDate time.Time) ([]Payments, error)
	GetTotalPaymentsAmount(ctx context.Context) (int32, error)
}

type IPaymentsUseCase interface {
	CreateOne(ctx context.Context, input *Input) error
	GetByID(ctx context.Context, id string) (PaymentsResponse, error)
	GetByInvoiceID(ctx context.Context, invoiceID string) ([]PaymentsResponse, error)
	GetByAccountID(ctx context.Context, accountID string) ([]PaymentsResponse, error)
	UpdateOne(ctx context.Context, id string, input *Input) error
	DeleteOne(ctx context.Context, id string) error
	GetAll(ctx context.Context) ([]PaymentsResponse, error)
	GetByDateRange(ctx context.Context, startDate, endDate string) ([]PaymentsResponse, error)
	GetTotalPaymentsAmount(ctx context.Context) (int32, error)
}
