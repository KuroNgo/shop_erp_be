package sales_order_usecase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	saleordersdomain "shop_erp_mono/domain/sales_and_distribution_management/sale_orders"
	"time"
)

type saleOrderUseCase struct {
	contextTimeout      time.Duration
	saleOrderRepository saleordersdomain.ISalesOrderRepository
}

func NewSaleOrderUseCase(contextTimeout time.Duration, saleOrderRepository saleordersdomain.ISalesOrderRepository) saleordersdomain.ISalesOrderUseCase {
	return &saleOrderUseCase{contextTimeout: contextTimeout, saleOrderRepository: saleOrderRepository}
}

func (s *saleOrderUseCase) GetByID(ctx context.Context, id primitive.ObjectID) (*saleordersdomain.SalesOrderResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *saleOrderUseCase) GetByCustomerID(ctx context.Context, customerID primitive.ObjectID) ([]saleordersdomain.SalesOrderResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *saleOrderUseCase) GetByStatus(ctx context.Context, status string) ([]saleordersdomain.SalesOrderResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *saleOrderUseCase) List(ctx context.Context) ([]saleordersdomain.SalesOrderResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *saleOrderUseCase) CreateOne(ctx context.Context, order saleordersdomain.SalesOrder) error {
	//TODO implement me
	panic("implement me")
}

func (s *saleOrderUseCase) UpdateOne(ctx context.Context, order saleordersdomain.SalesOrder) error {
	//TODO implement me
	panic("implement me")
}

func (s *saleOrderUseCase) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	//TODO implement me
	panic("implement me")
}
