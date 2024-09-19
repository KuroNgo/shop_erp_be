package invoice_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InvoiceRepository interface {
	CreateOne(ctx context.Context, invoice Invoice) error
	GetByID(ctx context.Context, id primitive.ObjectID) (*Invoice, error)
	GetByOrderID(ctx context.Context, orderID primitive.ObjectID) ([]Invoice, error)
	GetByStatus(ctx context.Context, status string) ([]Invoice, error)
	UpdateOne(ctx context.Context, invoice Invoice) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	List(ctx context.Context) ([]Invoice, error)
}

type InvoiceUseCase interface {
	CreateOne(ctx context.Context, input *Input) error
	GetByID(ctx context.Context, id string) (*InvoiceResponse, error)
	GetByOrderID(ctx context.Context, orderID string) ([]InvoiceResponse, error)
	GetByStatus(ctx context.Context, status string) ([]InvoiceResponse, error)
	UpdateOne(ctx context.Context, id string, input *Input) error
	DeleteOne(ctx context.Context, id string) error
	List(ctx context.Context) ([]InvoiceResponse, error)
}
