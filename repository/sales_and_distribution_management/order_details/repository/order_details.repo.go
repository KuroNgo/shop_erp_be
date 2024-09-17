package order_detail_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	order_details_domain "shop_erp_mono/domain/sales_and_distribution_management/order_details"
)

type orderDetailRepository struct {
	database              *mongo.Database
	orderDetailCollection string
}

func NewOrderDetailRepository(database *mongo.Database, orderDetailCollection string) order_details_domain.IOrderDetailRepository {
	return &orderDetailRepository{database: database, orderDetailCollection: orderDetailCollection}
}

func (o *orderDetailRepository) CreateOne(ctx context.Context, detail order_details_domain.OrderDetail) error {
	//TODO implement me
	panic("implement me")
}

func (o *orderDetailRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*order_details_domain.OrderDetail, error) {
	//TODO implement me
	panic("implement me")
}

func (o *orderDetailRepository) GetByOrderID(ctx context.Context, orderID primitive.ObjectID) ([]order_details_domain.OrderDetail, error) {
	//TODO implement me
	panic("implement me")
}

func (o *orderDetailRepository) GetByProductID(ctx context.Context, productID primitive.ObjectID) ([]order_details_domain.OrderDetail, error) {
	//TODO implement me
	panic("implement me")
}

func (o *orderDetailRepository) UpdateOne(ctx context.Context, id primitive.ObjectID, updatedDetail order_details_domain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (o *orderDetailRepository) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	//TODO implement me
	panic("implement me")
}

func (o *orderDetailRepository) List(ctx context.Context) ([]order_details_domain.OrderDetail, error) {
	//TODO implement me
	panic("implement me")
}
