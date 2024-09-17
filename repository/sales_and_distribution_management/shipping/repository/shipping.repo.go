package shipping_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	shippingdomain "shop_erp_mono/domain/sales_and_distribution_management/shipping"
	"time"
)

type shippingRepository struct {
	database           *mongo.Database
	shippingCollection string
}

func NewShippingRepository(database *mongo.Database, shippingCollection string) shippingdomain.IShippingRepository {
	return &shippingRepository{database: database, shippingCollection: shippingCollection}
}

func (s *shippingRepository) CreateOne(ctx context.Context, shipping shippingdomain.Shipping) error {
	shippingCollection := s.database.Collection(s.shippingCollection)

	_, err := shippingCollection.InsertOne(ctx, shipping)
	if err != nil {
		return err
	}

	return nil
}

func (s *shippingRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*shippingdomain.Shipping, error) {
	shippingCollection := s.database.Collection(s.shippingCollection)

	filter := bson.D{{"_id", id}}
	var shipping shippingdomain.Shipping
	if err := shippingCollection.FindOne(ctx, filter).Decode(&shipping); err != nil {
		return nil, err
	}

	return &shipping, nil
}

func (s *shippingRepository) GetByOrderID(ctx context.Context, orderID primitive.ObjectID) (*shippingdomain.Shipping, error) {
	shippingCollection := s.database.Collection(s.shippingCollection)

	filter := bson.D{{"order_id", orderID}}
	var shipping shippingdomain.Shipping
	if err := shippingCollection.FindOne(ctx, filter).Decode(&shipping); err != nil {
		return nil, err
	}

	return &shipping, nil
}

func (s *shippingRepository) UpdateOne(ctx context.Context, shipping shippingdomain.Shipping) error {
	shippingCollection := s.database.Collection(s.shippingCollection)

	filter := bson.D{{"_id", shipping.ID}}
	update := bson.D{
		{"$set", bson.D{
			{"order_id", shipping.OrderID},
			{"shipping_method", shipping.ShippingMethod},
			{"shipping_date", shipping.ShippingDate},
			{"estimated_delivery", shipping.EstimatedDelivery},
			{"actual_delivery", shipping.ActualDelivery},
			{"tracking_number", shipping.TrackingNumber},
			{"updated_at", time.Now()},
		}},
	}

	_, err := shippingCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (s *shippingRepository) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	shippingCollection := s.database.Collection(s.shippingCollection)

	filter := bson.D{{"_id", id}}
	_, err := shippingCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (s *shippingRepository) List(ctx context.Context) ([]shippingdomain.Shipping, error) {
	shippingCollection := s.database.Collection(s.shippingCollection)

	filter := bson.D{}
	cursor, err := shippingCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var shippings []shippingdomain.Shipping
	for cursor.Next(ctx) {
		var shipping shippingdomain.Shipping
		if err := cursor.Decode(&shipping); err != nil {
			return nil, err
		}

		shippings = append(shippings, shipping)
	}
	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return shippings, nil
}

func (s *shippingRepository) GetByStatus(ctx context.Context, status string) ([]shippingdomain.Shipping, error) {
	shippingCollection := s.database.Collection(s.shippingCollection)

	filter := bson.D{{"status", status}}
	cursor, err := shippingCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var shippings []shippingdomain.Shipping
	for cursor.Next(ctx) {
		var shipping shippingdomain.Shipping
		if err := cursor.Decode(&shipping); err != nil {
			return nil, err
		}

		shippings = append(shippings, shipping)
	}
	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return shippings, nil
}

func (s *shippingRepository) UpdateDeliveryStatus(ctx context.Context, id primitive.ObjectID, status string, actualDelivery *time.Time) error {
	shippingCollection := s.database.Collection(s.shippingCollection)

	filter := bson.D{{"_id", id}}
	update := bson.D{
		{"$set", bson.D{
			{"status", status},
			{"actual_delivery", actualDelivery},
			{"updated_at", time.Now()},
		}},
	}

	_, err := shippingCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}
