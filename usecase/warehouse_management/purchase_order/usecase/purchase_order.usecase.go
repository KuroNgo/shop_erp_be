package purchase_order_usecase

import (
	"context"
	purchaseorderdomain "shop_erp_mono/domain/warehouse_management/purchase_order"
	"shop_erp_mono/repository"
	"time"
)

type purchaseOrderUseCase struct {
	contextTimeout          time.Duration
	purchaseOrderRepository purchaseorderdomain.IPurchaseOrderRepository
}

func (p *purchaseOrderUseCase) GetByID(ctx context.Context, id string) (*purchaseorderdomain.PurchaseOrder, error) {
	//TODO implement me
	panic("implement me")
}

func (p *purchaseOrderUseCase) Create(ctx context.Context, order *purchaseorderdomain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (p *purchaseOrderUseCase) Update(ctx context.Context, id string, order *purchaseorderdomain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (p *purchaseOrderUseCase) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (p *purchaseOrderUseCase) GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]purchaseorderdomain.PurchaseOrder, error) {
	//TODO implement me
	panic("implement me")
}

func (p *purchaseOrderUseCase) GetBySupplierID(ctx context.Context, supplierID string) ([]purchaseorderdomain.PurchaseOrder, error) {
	//TODO implement me
	panic("implement me")
}

func (p *purchaseOrderUseCase) UpdateStatus(ctx context.Context, id string, status string) error {
	//TODO implement me
	panic("implement me")
}

func NewPurchaseOrderUseCase(contextTimeout time.Duration, purchaseOrderRepository purchaseorderdomain.IPurchaseOrderRepository) purchaseorderdomain.IPurchaseOrderUseCase {
	return &purchaseOrderUseCase{contextTimeout: contextTimeout, purchaseOrderRepository: purchaseOrderRepository}
}
