package invoices_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IInvoicesRepository interface {
	CreateInvoice(ctx context.Context, budget *Invoices) (Invoices, error)
	GetInvoiceByID(ctx context.Context, id primitive.ObjectID) (Invoices, error)
	GetInvoiceByName(ctx context.Context, name string) (Invoices, error)
	UpdateInvoice(ctx context.Context, budget *Invoices) (Invoices, error)
	DeleteInvoice(ctx context.Context, id primitive.ObjectID) error
	ListInvoices(ctx context.Context) ([]Invoices, error)
}

type IInvoicesUseCase interface {
	CreateInvoice(ctx context.Context, input *Input) (InvoicesResponse, error)
	GetInvoice(ctx context.Context, id string) (InvoicesResponse, error)
	GetInvoiceByName(ctx context.Context, name string) (InvoicesResponse, error)
	UpdateInvoice(ctx context.Context, id string, input *Input) (InvoicesResponse, error)
	DeleteInvoice(ctx context.Context, id string) error
	ListInvoices(ctx context.Context) ([]InvoicesResponse, error)
}
