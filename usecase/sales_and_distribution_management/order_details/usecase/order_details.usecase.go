package order_detail_usecase

import (
	"context"
	order_details_domain "shop_erp_mono/domain/sales_and_distribution_management/order_details"
	"time"
)

type orderDetailUseCase struct {
	contextTimeout        time.Duration
	orderDetailRepository order_details_domain.IOrderDetailRepository
}

func NewOrderDetailUseCase(contextTimeout time.Duration, orderDetailRepository order_details_domain.IOrderDetailRepository) order_details_domain.IOrderDetailUseCase {
	return &orderDetailUseCase{contextTimeout: contextTimeout, orderDetailRepository: orderDetailRepository}
}

func (o *orderDetailUseCase) CreateOne(ctx context.Context, input *order_details_domain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (o *orderDetailUseCase) GetByID(ctx context.Context, id string) (*order_details_domain.OrderDetailResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (o *orderDetailUseCase) GetByOrderID(ctx context.Context, orderID string) ([]order_details_domain.OrderDetailResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (o *orderDetailUseCase) GetByProductID(ctx context.Context, productID string) ([]order_details_domain.OrderDetailResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (o *orderDetailUseCase) UpdateOne(ctx context.Context, id string, input *order_details_domain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (o *orderDetailUseCase) DeleteOne(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (o *orderDetailUseCase) List(ctx context.Context) ([]order_details_domain.OrderDetailResponse, error) {
	//TODO implement me
	panic("implement me")
}
