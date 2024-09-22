package invoice_usecase

import (
	"context"
	"encoding/json"
	"github.com/allegro/bigcache/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	invoices_domain "shop_erp_mono/domain/sales_and_distribution_management/invoices"
	sale_orders_domain "shop_erp_mono/domain/sales_and_distribution_management/sale_orders"
	"shop_erp_mono/usecase/sales_and_distribution_management/invoices/validate"
	"sync"
	"time"
)

type invoiceUseCase struct {
	contextTimeout       time.Duration
	invoiceRepository    invoices_domain.InvoiceRepository
	salesOrderRepository sale_orders_domain.ISalesOrderRepository
	cache                *bigcache.BigCache
}

var (
	wg sync.WaitGroup
	mu sync.Mutex
)

func NewInvoiceUseCase(contextTimeout time.Duration, invoiceRepository invoices_domain.InvoiceRepository,
	salesOrderRepository sale_orders_domain.ISalesOrderRepository, cacheTTL time.Duration) invoices_domain.InvoiceUseCase {
	config := bigcache.Config{
		Shards:           1024,
		LifeWindow:       cacheTTL,
		MaxEntrySize:     512,
		CleanWindow:      1 * time.Minute,
		HardMaxCacheSize: 8192,
	}

	cache, err := bigcache.New(context.Background(), config)
	if err != nil {
		return nil
	}
	return &invoiceUseCase{contextTimeout: contextTimeout, cache: cache, invoiceRepository: invoiceRepository, salesOrderRepository: salesOrderRepository}
}

func (i *invoiceUseCase) CreateOne(ctx context.Context, input *invoices_domain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, i.contextTimeout)
	defer cancel()

	orderID, err := primitive.ObjectIDFromHex(input.OrderID)
	if err != nil {
		return err
	}

	if err = validate.Invoices(input); err != nil {
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

	err = i.cache.Delete("invoices")
	if err != nil {
		return err
	}

	mu.Lock()
	defer mu.Unlock()

	return i.invoiceRepository.CreateOne(ctx, invoice)
}

func (i *invoiceUseCase) GetByID(ctx context.Context, id string) (*invoices_domain.InvoiceResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, i.contextTimeout)
	defer cancel()

	data, err := i.cache.Get(id)
	if err != nil {
		return nil, err
	}
	if data != nil {
		var response *invoices_domain.InvoiceResponse
		err = json.Unmarshal(data, response)
		return response, nil
	}

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

	data, err = json.Marshal(response)
	err = i.cache.Set(id, data)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (i *invoiceUseCase) GetByOrderID(ctx context.Context, orderID string) ([]invoices_domain.InvoiceResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, i.contextTimeout)
	defer cancel()

	data, err := i.cache.Get(orderID)
	if err != nil {
		return nil, err
	}
	if data != nil {
		var response []invoices_domain.InvoiceResponse
		err = json.Unmarshal(data, &response)
		return response, nil
	}

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

	data, err = json.Marshal(responses)
	err = i.cache.Set("invoices", data)
	if err != nil {
		return nil, err
	}
	return responses, nil
}

func (i *invoiceUseCase) GetByStatus(ctx context.Context, status string) ([]invoices_domain.InvoiceResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, i.contextTimeout)
	defer cancel()

	data, err := i.cache.Get(status)
	if err != nil {
		return nil, err
	}
	if data != nil {
		var response []invoices_domain.InvoiceResponse
		err = json.Unmarshal(data, &response)
		return response, nil
	}

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

	data, err = json.Marshal(responses)
	err = i.cache.Set("invoices", data)
	if err != nil {
		return nil, err
	}
	return responses, nil
}

func (i *invoiceUseCase) UpdateOne(ctx context.Context, id string, input *invoices_domain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, i.contextTimeout)
	defer cancel()

	if err := validate.Invoices(input); err != nil {
		return err
	}

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

	wg.Add(2)
	go func() {
		defer wg.Done()
		err = i.cache.Delete(id)
		if err != nil {
			return
		}
	}()
	go func() {
		defer wg.Done()
		err = i.cache.Delete("invoices")
		if err != nil {
			return
		}
	}()
	wg.Wait()

	return i.invoiceRepository.UpdateOne(ctx, invoice)
}

func (i *invoiceUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, i.contextTimeout)
	defer cancel()

	invoiceID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	wg.Add(2)
	go func() {
		defer wg.Done()
		err = i.cache.Delete(id)
		if err != nil {
			return
		}
	}()
	go func() {
		defer wg.Done()
		err = i.cache.Delete("invoices")
		if err != nil {
			return
		}
	}()
	wg.Wait()
	return i.invoiceRepository.DeleteOne(ctx, invoiceID)
}

func (i *invoiceUseCase) GetAll(ctx context.Context) ([]invoices_domain.InvoiceResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, i.contextTimeout)
	defer cancel()

	data, err := i.cache.Get("invoices")
	if err != nil {
		return nil, err
	}

	if data != nil {
		var response []invoices_domain.InvoiceResponse
		err = json.Unmarshal(data, &response)
		if err != nil {
			return nil, err
		}
		return response, nil
	}
	invoiceData, err := i.invoiceRepository.GetAll(ctx)
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

	data, err = json.Marshal(responses)
	if err != nil {
		return nil, err
	}
	err = i.cache.Set("invoices", data)
	if err != nil {
		return nil, err
	}

	return responses, nil
}
