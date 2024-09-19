package sales_order_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	saleordersdomain "shop_erp_mono/domain/sales_and_distribution_management/sale_orders"
	"time"
)

type saleOrderRepository struct {
	database            *mongo.Database
	saleOrderCollection string
}

func NewSaleOrderRepository(database *mongo.Database, saleOrderCollection string) saleordersdomain.ISalesOrderRepository {
	return &saleOrderRepository{database: database, saleOrderCollection: saleOrderCollection}
}

func (s *saleOrderRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*saleordersdomain.SalesOrder, error) {
	salesOrderCollection := s.database.Collection(s.saleOrderCollection)

	filter := bson.M{"_id": id}
	var salesOrder saleordersdomain.SalesOrder
	if err := salesOrderCollection.FindOne(ctx, filter).Decode(&salesOrder); err != nil {
		return nil, err
	}

	return &salesOrder, nil
}

func (s *saleOrderRepository) GetByCustomerID(ctx context.Context, customerID primitive.ObjectID) ([]saleordersdomain.SalesOrder, error) {
	salesOrderCollection := s.database.Collection(s.saleOrderCollection)

	filter := bson.M{"customer_id": customerID}
	cursor, err := salesOrderCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []saleordersdomain.SalesOrder
	orders = make([]saleordersdomain.SalesOrder, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var order saleordersdomain.SalesOrder
		if err = cursor.Decode(&order); err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (s *saleOrderRepository) GetByStatus(ctx context.Context, status string) ([]saleordersdomain.SalesOrder, error) {
	salesOrderCollection := s.database.Collection(s.saleOrderCollection)

	filter := bson.M{"status": status}
	cursor, err := salesOrderCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []saleordersdomain.SalesOrder
	orders = make([]saleordersdomain.SalesOrder, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var order saleordersdomain.SalesOrder
		if err = cursor.Decode(&order); err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (s *saleOrderRepository) GetAll(ctx context.Context) ([]saleordersdomain.SalesOrder, error) {
	salesOrderCollection := s.database.Collection(s.saleOrderCollection)

	filter := bson.M{}
	cursor, err := salesOrderCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []saleordersdomain.SalesOrder
	orders = make([]saleordersdomain.SalesOrder, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var order saleordersdomain.SalesOrder
		if err = cursor.Decode(&order); err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (s *saleOrderRepository) CreateOne(ctx context.Context, order saleordersdomain.SalesOrder) error {
	salesOrderCollection := s.database.Collection(s.saleOrderCollection)

	_, err := salesOrderCollection.InsertOne(ctx, order)
	if err != nil {
		return err
	}

	return nil
}

func (s *saleOrderRepository) UpdateOne(ctx context.Context, order saleordersdomain.SalesOrder) error {
	salesOrderCollection := s.database.Collection(s.saleOrderCollection)

	filter := bson.M{"_id": order.ID}
	update := bson.M{"$set": bson.M{
		"order_number":     order.OrderNumber,
		"customer_id":      order.CustomerID,
		"order_date":       order.OrderDate,
		"shipping_address": order.ShippingAddress,
		"total_amount":     order.TotalAmount,
		"status":           order.Status,
		"updated_at":       time.Now(),
	}}
	_, err := salesOrderCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (s *saleOrderRepository) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	salesOrderCollection := s.database.Collection(s.saleOrderCollection)

	filter := bson.M{"_id": id}
	_, err := salesOrderCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
