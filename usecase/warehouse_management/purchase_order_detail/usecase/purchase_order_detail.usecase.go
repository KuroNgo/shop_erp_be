package purchase_order_detail_usecase

import (
	"context"
	purchaseorderdetaildomain "shop_erp_mono/domain/warehouse_management/purchase_order_detail"
	"shop_erp_mono/repository"
	"time"
)

type purchaseOrderDetailUseCase struct {
	contextTimeout                time.Duration
	purchaseOrderDetailRepository purchaseorderdetaildomain.IPurchaseOrderDetailRepository
}

func NewProductOrderDetailRepository(contextTimeout time.Duration, purchaseOrderDetailRepository purchaseorderdetaildomain.IPurchaseOrderDetailRepository) purchaseorderdetaildomain.IPurchaseOrderDetailUseCase {
	return &purchaseOrderDetailUseCase{contextTimeout: contextTimeout, purchaseOrderDetailRepository: purchaseOrderDetailRepository}
}

func (p *purchaseOrderDetailUseCase) GetByID(ctx context.Context, id string) (*purchaseorderdetaildomain.PurchaseOrderDetail, error) {
	//TODO implement me
	panic("implement me")
}

func (p *purchaseOrderDetailUseCase) GetByPurchaseOrderID(ctx context.Context, purchaseOrderID string) ([]purchaseorderdetaildomain.PurchaseOrderDetail, error) {
	//TODO implement me
	panic("implement me")
}

func (p *purchaseOrderDetailUseCase) Create(ctx context.Context, input *purchaseorderdetaildomain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (p *purchaseOrderDetailUseCase) Update(ctx context.Context, id string, input *purchaseorderdetaildomain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (p *purchaseOrderDetailUseCase) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (p *purchaseOrderDetailUseCase) GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]purchaseorderdetaildomain.PurchaseOrderDetail, error) {
	//TODO implement me
	panic("implement me")
}

func (p *purchaseOrderDetailUseCase) GetAll(ctx context.Context) ([]purchaseorderdetaildomain.PurchaseOrderDetail, error) {
	//TODO implement me
	panic("implement me")
}
