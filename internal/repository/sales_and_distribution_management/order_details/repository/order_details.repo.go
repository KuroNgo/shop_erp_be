package order_detail_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	orderdetailsdomain "shop_erp_mono/internal/domain/sales_and_distribution_management/order_details"
	"time"
)

type orderDetailRepository struct {
	database              *mongo.Database
	orderDetailCollection string
}

func NewOrderDetailRepository(database *mongo.Database, orderDetailCollection string) orderdetailsdomain.IOrderDetailRepository {
	return &orderDetailRepository{database: database, orderDetailCollection: orderDetailCollection}
}

func (o *orderDetailRepository) CreateOne(ctx context.Context, detail orderdetailsdomain.OrderDetail) error {
	orderDetailCollection := o.database.Collection(o.orderDetailCollection)

	_, err := orderDetailCollection.InsertOne(ctx, detail)
	if err != nil {
		return err
	}

	return nil
}

func (o *orderDetailRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*orderdetailsdomain.OrderDetail, error) {
	orderDetailCollection := o.database.Collection(o.orderDetailCollection)

	filter := bson.M{"_id": id}
	var orderDetail orderdetailsdomain.OrderDetail
	err := orderDetailCollection.FindOne(ctx, filter).Decode(&orderDetail)
	if err != nil {
		return nil, err
	}

	return &orderDetail, nil
}

func (o *orderDetailRepository) GetByOrderID(ctx context.Context, orderID primitive.ObjectID) ([]orderdetailsdomain.OrderDetail, error) {
	orderDetailCollection := o.database.Collection(o.orderDetailCollection)

	filter := bson.M{"order_id": orderID}
	cursor, err := orderDetailCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orderDetails []orderdetailsdomain.OrderDetail
	orderDetails = make([]orderdetailsdomain.OrderDetail, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var order orderdetailsdomain.OrderDetail
		if err = cursor.Decode(&order); err != nil {
			return nil, err
		}

		orderDetails = append(orderDetails, order)
	}

	return orderDetails, nil
}

func (o *orderDetailRepository) GetByProductID(ctx context.Context, productID primitive.ObjectID) ([]orderdetailsdomain.OrderDetail, error) {
	orderDetailCollection := o.database.Collection(o.orderDetailCollection)

	filter := bson.M{"product_id": productID}
	cursor, err := orderDetailCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orderDetails []orderdetailsdomain.OrderDetail
	orderDetails = make([]orderdetailsdomain.OrderDetail, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var order orderdetailsdomain.OrderDetail
		if err = cursor.Decode(&order); err != nil {
			return nil, err
		}

		orderDetails = append(orderDetails, order)
	}

	return orderDetails, nil
}

func (o *orderDetailRepository) UpdateOne(ctx context.Context, detail orderdetailsdomain.OrderDetail) error {
	orderDetailCollection := o.database.Collection(o.orderDetailCollection)

	filter := bson.M{"_id": detail.ID}
	update := bson.M{"$set": bson.M{
		"order_id":    detail.OrderID,
		"product_id":  detail.ProductID,
		"quantity":    detail.Quantity,
		"unit_price":  detail.UnitPrice,
		"total_price": detail.TotalPrice,
		"updated_at":  time.Now(),
	}}

	_, err := orderDetailCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (o *orderDetailRepository) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	orderDetailCollection := o.database.Collection(o.orderDetailCollection)

	filter := bson.M{"_id": id}
	_, err := orderDetailCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (o *orderDetailRepository) GetAll(ctx context.Context) ([]orderdetailsdomain.OrderDetail, error) {
	orderDetailCollection := o.database.Collection(o.orderDetailCollection)

	filter := bson.M{}
	cursor, err := orderDetailCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orderDetails []orderdetailsdomain.OrderDetail
	orderDetails = make([]orderdetailsdomain.OrderDetail, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var order orderdetailsdomain.OrderDetail
		if err = cursor.Decode(&order); err != nil {
			return nil, err
		}

		orderDetails = append(orderDetails, order)
	}

	return orderDetails, nil
}
