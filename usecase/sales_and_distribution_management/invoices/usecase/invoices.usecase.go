package invoice_usecase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	invoices_domain "shop_erp_mono/domain/sales_and_distribution_management/invoices"
	sale_orders_domain "shop_erp_mono/domain/sales_and_distribution_management/sale_orders"
	"time"
)

type invoiceUseCase struct {
	contextTimeout       time.Duration
	invoiceRepository    invoices_domain.InvoiceRepository
	salesOrderRepository sale_orders_domain.ISalesOrderRepository
}

func NewInvoiceUseCase(contextTimeout time.Duration, invoiceRepository invoices_domain.InvoiceRepository, salesOrderRepository sale_orders_domain.ISalesOrderRepository) invoices_domain.InvoiceUseCase {
	return &invoiceUseCase{contextTimeout: contextTimeout, invoiceRepository: invoiceRepository, salesOrderRepository: salesOrderRepository}
}

func (i *invoiceUseCase) CreateOne(ctx context.Context, input *invoices_domain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, i.contextTimeout)
	defer cancel()

	orderID, err := primitive.ObjectIDFromHex(input.OrderID)
	if err != nil {
		return err
	}

	invoice := invoices_domain.Invoice{
		ID:          primitive.NewObjectID(),
		OrderID:     orderID,
		InvoiceDate: input.InvoiceDate,
		DueDate:     input.DueDate,
		AmountPaid:  input.AmountPaid,
		AmountDue:   input.AmountDue,
		Status:      input.Status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return i.invoiceRepository.CreateOne(ctx, invoice)
}

func (i *invoiceUseCase) GetByID(ctx context.Context, id string) (*invoices_domain.InvoiceResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, i.contextTimeout)
	defer cancel()

	invoiceID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	invoiceData, err := i.invoiceRepository.GetByID(ctx, invoiceID)
	if err != nil {
		return nil, err
	}

	orderData, err := i.salesOrderRepository.GetByID(ctx, invoiceData.OrderID)
	if err != nil {
		return nil, err
	}

	response := &invoices_domain.InvoiceResponse{
		Invoice: *invoiceData,
		Order:   *orderData,
	}

	return response, nil
}

func (i *invoiceUseCase) GetByOrderID(ctx context.Context, orderID string) ([]invoices_domain.InvoiceResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, i.contextTimeout)
	defer cancel()

	idOrder, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		return nil, err
	}

	invoiceData, err := i.invoiceRepository.GetByOrderID(ctx, idOrder)
	if err != nil {
		return nil, err
	}

	var responses []invoices_domain.InvoiceResponse
	responses = make([]invoices_domain.InvoiceResponse, 0, len(invoiceData))
	for _, invoice := range invoiceData {
		orderData, err := i.salesOrderRepository.GetByID(ctx, invoice.OrderID)
		if err != nil {
			return nil, err
		}

		response := invoices_domain.InvoiceResponse{
			Invoice: invoice,
			Order:   *orderData,
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (i *invoiceUseCase) GetByStatus(ctx context.Context, status string) ([]invoices_domain.InvoiceResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, i.contextTimeout)
	defer cancel()

	invoiceData, err := i.invoiceRepository.GetByStatus(ctx, status)
	if err != nil {
		return nil, err
	}

	var responses []invoices_domain.InvoiceResponse
	responses = make([]invoices_domain.InvoiceResponse, 0, len(invoiceData))
	for _, invoice := range invoiceData {
		orderData, err := i.salesOrderRepository.GetByID(ctx, invoice.OrderID)
		if err != nil {
			return nil, err
		}

		response := invoices_domain.InvoiceResponse{
			Invoice: invoice,
			Order:   *orderData,
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (i *invoiceUseCase) UpdateOne(ctx context.Context, id string, input *invoices_domain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, i.contextTimeout)
	defer cancel()

	invoiceID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	orderID, err := primitive.ObjectIDFromHex(input.OrderID)
	if err != nil {
		return err
	}

	invoice := invoices_domain.Invoice{
		ID:          invoiceID,
		OrderID:     orderID,
		InvoiceDate: input.InvoiceDate,
		DueDate:     input.DueDate,
		AmountPaid:  input.AmountPaid,
		AmountDue:   input.AmountDue,
		Status:      input.Status,
		UpdatedAt:   time.Now(),
	}

	return i.invoiceRepository.UpdateOne(ctx, invoice)
}

func (i *invoiceUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, i.contextTimeout)
	defer cancel()

	invoiceID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return i.invoiceRepository.DeleteOne(ctx, invoiceID)
}

func (i *invoiceUseCase) List(ctx context.Context) ([]invoices_domain.InvoiceResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, i.contextTimeout)
	defer cancel()

	invoiceData, err := i.invoiceRepository.List(ctx)
	if err != nil {
		return nil, err
	}

	var responses []invoices_domain.InvoiceResponse
	responses = make([]invoices_domain.InvoiceResponse, 0, len(invoiceData))
	for _, invoice := range invoiceData {
		orderData, err := i.salesOrderRepository.GetByID(ctx, invoice.OrderID)
		if err != nil {
			return nil, err
		}

		response := invoices_domain.InvoiceResponse{
			Invoice: invoice,
			Order:   *orderData,
		}

		responses = append(responses, response)
	}

	return responses, nil
}
