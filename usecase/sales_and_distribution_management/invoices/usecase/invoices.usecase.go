package invoice_usecase

import (
	"context"
	invoices_domain "shop_erp_mono/domain/sales_and_distribution_management/invoices"
	"time"
)

type invoiceUseCase struct {
	contextTimeout    time.Duration
	invoiceRepository invoices_domain.InvoiceRepository
}

func NewInvoiceUseCase(contextTimeout time.Duration, invoiceRepository invoices_domain.InvoiceRepository) invoices_domain.InvoiceUseCase {
	return &invoiceUseCase{contextTimeout: contextTimeout, invoiceRepository: invoiceRepository}
}

func (i *invoiceUseCase) CreateOne(ctx context.Context, input *invoices_domain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (i *invoiceUseCase) GetByID(ctx context.Context, id string) (*invoices_domain.InvoiceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (i *invoiceUseCase) GetByOrderID(ctx context.Context, orderID string) ([]invoices_domain.InvoiceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (i *invoiceUseCase) GetByStatus(ctx context.Context, status string) ([]invoices_domain.InvoiceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (i *invoiceUseCase) UpdateOne(ctx context.Context, id string, input *invoices_domain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (i *invoiceUseCase) DeleteOne(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (i *invoiceUseCase) List(ctx context.Context) ([]invoices_domain.InvoiceResponse, error) {
	//TODO implement me
	panic("implement me")
}
