package purchase_order_detail_usecase

import (
	"context"
	purchaseorderdetaildomain "shop_erp_mono/domain/warehouse_management/purchase_order_detail"
	"shop_erp_mono/repository"
	"time"
)

type productOrderDetailUseCase struct {
	contextTimeout               time.Duration
	productOrderDetailRepository purchaseorderdetaildomain.IPurchaseOrderDetailRepository
}

func NewProductOrderDetailRepository(contextTimeout time.Duration, productOrderDetailRepository purchaseorderdetaildomain.IPurchaseOrderDetailRepository) purchaseorderdetaildomain.IPurchaseOrderDetailUseCase {
	return &productOrderDetailUseCase{contextTimeout: contextTimeout, productOrderDetailRepository: productOrderDetailRepository}
}

func (p *productOrderDetailUseCase) GetByID(ctx context.Context, id string) (*purchaseorderdetaildomain.PurchaseOrderDetail, error) {
	//TODO implement me
	panic("implement me")
}

func (p *productOrderDetailUseCase) GetByPurchaseOrderID(ctx context.Context, purchaseOrderID string) ([]purchaseorderdetaildomain.PurchaseOrderDetail, error) {
	//TODO implement me
	panic("implement me")
}

func (p *productOrderDetailUseCase) Create(ctx context.Context, input *purchaseorderdetaildomain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (p *productOrderDetailUseCase) Update(ctx context.Context, id string, input *purchaseorderdetaildomain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (p *productOrderDetailUseCase) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (p *productOrderDetailUseCase) GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]purchaseorderdetaildomain.PurchaseOrderDetail, error) {
	//TODO implement me
	panic("implement me")
}

func (p *productOrderDetailUseCase) GetAll(ctx context.Context) ([]purchaseorderdetaildomain.PurchaseOrderDetail, error) {
	//TODO implement me
	panic("implement me")
}
