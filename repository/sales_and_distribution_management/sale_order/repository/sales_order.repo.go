package sales_order_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	saleordersdomain "shop_erp_mono/domain/sales_and_distribution_management/sale_orders"
)

type saleOrderRepository struct {
	database            *mongo.Database
	saleOrderCollection string
}

func NewSaleOrderRepository(database *mongo.Database, saleOrderCollection string) saleordersdomain.ISalesOrderRepository {
	return &saleOrderRepository{database: database, saleOrderCollection: saleOrderCollection}
}

func (s *saleOrderRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*saleordersdomain.SalesOrderResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *saleOrderRepository) GetByCustomerID(ctx context.Context, customerID primitive.ObjectID) ([]saleordersdomain.SalesOrderResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *saleOrderRepository) GetByStatus(ctx context.Context, status string) ([]saleordersdomain.SalesOrderResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *saleOrderRepository) List(ctx context.Context, filters map[string]interface{}) ([]saleordersdomain.SalesOrderResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *saleOrderRepository) CreateOne(ctx context.Context, order saleordersdomain.SalesOrder) error {
	//TODO implement me
	panic("implement me")
}

func (s *saleOrderRepository) UpdateOne(ctx context.Context, id primitive.ObjectID, updatedOrder saleordersdomain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (s *saleOrderRepository) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	//TODO implement me
	panic("implement me")
}
