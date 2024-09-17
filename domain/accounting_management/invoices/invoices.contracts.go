package invoices_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type InvoicesRepository interface {
	CreateInvoice(ctx context.Context, invoice *Invoices) error
	GetInvoiceByID(ctx context.Context, id primitive.ObjectID) (*Invoices, error)
	GetInvoiceByName(ctx context.Context, name string) (*Invoices, error)
	UpdateInvoice(ctx context.Context, invoice *Invoices) error
	DeleteInvoice(ctx context.Context, id primitive.ObjectID) error
	ListInvoices(ctx context.Context) ([]Invoices, error)
	GetInvoicesByDateRange(ctx context.Context, startDate, endDate time.Time) ([]Invoices, error)
	GetOverdueInvoices(ctx context.Context) ([]Invoices, error)
	MarkInvoiceAsPaid(ctx context.Context, id primitive.ObjectID) error
}

type InvoicesUseCase interface {
	CreateInvoice(ctx context.Context, input *Input) (InvoicesResponse, error)
	GetInvoice(ctx context.Context, id string) (InvoicesResponse, error)
	GetInvoiceByName(ctx context.Context, name string) (InvoicesResponse, error)
	UpdateInvoice(ctx context.Context, id string, input *Input) (InvoicesResponse, error)
	DeleteInvoice(ctx context.Context, id string) error
	ListInvoices(ctx context.Context) ([]InvoicesResponse, error)
	GetInvoicesByDateRange(ctx context.Context, startDate, endDate string) ([]InvoicesResponse, error)
	GetOverdueInvoices(ctx context.Context) ([]InvoicesResponse, error)
	MarkInvoiceAsPaid(ctx context.Context, id string) error
	SendInvoiceReminder(ctx context.Context, id string) error
	GenerateInvoiceReport(ctx context.Context, startDate, endDate string) (InvoiceReport, error)
}
