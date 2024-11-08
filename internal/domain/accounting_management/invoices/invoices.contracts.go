package invoices_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type InvoicesRepository interface {
	CreateOne(ctx context.Context, invoice *Invoices) error
	GetOneByID(ctx context.Context, id primitive.ObjectID) (*Invoices, error)
	GetOneByName(ctx context.Context, name string) (*Invoices, error)
	UpdateOne(ctx context.Context, invoice *Invoices) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	GetAll(ctx context.Context) ([]Invoices, error)
	GetByDateRange(ctx context.Context, startDate, endDate time.Time) ([]Invoices, error)
	GetOverdueInvoices(ctx context.Context) ([]Invoices, error)
	MarkInvoiceAsPaid(ctx context.Context, id primitive.ObjectID) error
}

type InvoicesUseCase interface {
	CreateOne(ctx context.Context, input *Input) (InvoicesResponse, error)
	GetInOne(ctx context.Context, id string) (InvoicesResponse, error)
	GetOneByName(ctx context.Context, name string) (InvoicesResponse, error)
	UpdateOne(ctx context.Context, id string, input *Input) (InvoicesResponse, error)
	DeleteOne(ctx context.Context, id string) error
	GetAll(ctx context.Context) ([]InvoicesResponse, error)
	GetByDateRange(ctx context.Context, startDate, endDate string) ([]InvoicesResponse, error)
	GetOverdueInvoices(ctx context.Context) ([]InvoicesResponse, error)
	MarkInvoiceAsPaid(ctx context.Context, id string) error
	SendInvoiceReminder(ctx context.Context, id string) error
	GenerateInvoiceReport(ctx context.Context, startDate, endDate string) (InvoiceReport, error)
}
