package purchase_order_usecase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	purchaseorderdomain "shop_erp_mono/domain/warehouse_management/purchase_order"
	supplierdomain "shop_erp_mono/domain/warehouse_management/supplier"
	"shop_erp_mono/repository"
	"shop_erp_mono/usecase/warehouse_management/purchase_order/validate"
	"time"
)

type purchaseOrderUseCase struct {
	contextTimeout          time.Duration
	purchaseOrderRepository purchaseorderdomain.IPurchaseOrderRepository
	supplierRepository      supplierdomain.ISupplierRepository
}

func NewPurchaseOrderUseCase(contextTimeout time.Duration, purchaseOrderRepository purchaseorderdomain.IPurchaseOrderRepository, supplierRepository supplierdomain.ISupplierRepository) purchaseorderdomain.IPurchaseOrderUseCase {
	return &purchaseOrderUseCase{contextTimeout: contextTimeout, purchaseOrderRepository: purchaseOrderRepository, supplierRepository: supplierRepository}
}

func (p *purchaseOrderUseCase) GetByID(ctx context.Context, id string) (*purchaseorderdomain.PurchaseOrderResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	purchaseOrderID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	purchaseOrderData, err := p.purchaseOrderRepository.GetByID(ctx, purchaseOrderID)
	if err != nil {
		return nil, err
	}

	response := &purchaseorderdomain.PurchaseOrderResponse{
		PurchaseOrder: *purchaseOrderData,
	}

	return response, nil
}

func (p *purchaseOrderUseCase) CreateOne(ctx context.Context, input *purchaseorderdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	if err := validate.PurchaseOrder(input); err != nil {
		return err
	}

	supplierData, err := p.supplierRepository.GetByName(ctx, input.Supplier)
	if err != nil {
		return err
	}

	order := &purchaseorderdomain.PurchaseOrder{
		ID:          primitive.NewObjectID(),
		SupplierID:  supplierData.ID,
		OrderNumber: input.OrderNumber,
		OrderDate:   input.OrderDate,
		TotalAmount: input.TotalAmount,
		Status:      input.Status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return p.purchaseOrderRepository.CreateOne(ctx, order)
}

func (p *purchaseOrderUseCase) UpdateOne(ctx context.Context, id string, input *purchaseorderdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	if err := validate.PurchaseOrder(input); err != nil {
		return err
	}

	purchaseOrderID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	supplierData, err := p.supplierRepository.GetByName(ctx, input.Supplier)
	if err != nil {
		return err
	}

	order := &purchaseorderdomain.PurchaseOrder{
		ID:          purchaseOrderID,
		SupplierID:  supplierData.ID,
		OrderNumber: input.OrderNumber,
		OrderDate:   input.OrderDate,
		TotalAmount: input.TotalAmount,
		Status:      input.Status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return p.purchaseOrderRepository.UpdateOne(ctx, order)
}

func (p *purchaseOrderUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	purchaseOrderID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return p.purchaseOrderRepository.DeleteOne(ctx, purchaseOrderID)
}

func (p *purchaseOrderUseCase) GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]purchaseorderdomain.PurchaseOrderResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	purchaseOrderData, err := p.purchaseOrderRepository.GetAllWithPagination(ctx, pagination)
	if err != nil {
		return nil, err
	}

	var responses []purchaseorderdomain.PurchaseOrderResponse
	responses = make([]purchaseorderdomain.PurchaseOrderResponse, 0, len(purchaseOrderData))
	for _, purchaseOrder := range purchaseOrderData {
		response := purchaseorderdomain.PurchaseOrderResponse{
			PurchaseOrder: purchaseOrder,
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (p *purchaseOrderUseCase) GetBySupplierID(ctx context.Context, supplierID string) ([]purchaseorderdomain.PurchaseOrderResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	idSupplier, err := primitive.ObjectIDFromHex(supplierID)
	if err != nil {
		return nil, err
	}

	purchaseOrderData, err := p.purchaseOrderRepository.GetBySupplierID(ctx, idSupplier)
	if err != nil {
		return nil, err
	}

	var responses []purchaseorderdomain.PurchaseOrderResponse
	responses = make([]purchaseorderdomain.PurchaseOrderResponse, 0, len(purchaseOrderData))
	for _, purchaseOrder := range purchaseOrderData {
		response := purchaseorderdomain.PurchaseOrderResponse{
			PurchaseOrder: purchaseOrder,
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (p *purchaseOrderUseCase) UpdateStatus(ctx context.Context, id string, status string) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	purchaseOrderID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	order := &purchaseorderdomain.PurchaseOrder{
		ID:        purchaseOrderID,
		Status:    status,
		UpdatedAt: time.Now(),
	}

	return p.purchaseOrderRepository.UpdateOne(ctx, order)
}
