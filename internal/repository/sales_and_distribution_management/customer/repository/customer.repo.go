package customer_repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	customerdomain "shop_erp_mono/internal/domain/sales_and_distribution_management/customer"
)

type customerRepository struct {
	database           *mongo.Database
	customerCollection string
}

func NewCustomerRepository(database *mongo.Database, customerCollection string) customerdomain.ICustomerRepository {
	return &customerRepository{database: database, customerCollection: customerCollection}
}

func (c *customerRepository) CreateOne(ctx context.Context, customer *customerdomain.Customer) error {
	customerCollection := c.database.Collection(c.customerCollection)

	_, err := customerCollection.InsertOne(ctx, customer)
	if err != nil {
		return err
	}

	return nil
}

func (c *customerRepository) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	customerCollection := c.database.Collection(c.customerCollection)

	filter := bson.M{"_id": id}
	_, err := customerCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (c *customerRepository) UpdateOne(ctx context.Context, customer *customerdomain.Customer) error {
	customerCollection := c.database.Collection(c.customerCollection)

	filter := bson.M{"_id": customer.ID}
	update := bson.M{"$set": bson.M{
		"first_name":   customer.FirstName,
		"last_name":    customer.LastName,
		"email":        customer.Email,
		"phone_number": customer.PhoneNumber,
		"address":      customer.Address,
		"city":         customer.City,
		"updated_at":   customer.UpdatedAt,
	}}
	_, err := customerCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (c *customerRepository) GetOneByID(ctx context.Context, id primitive.ObjectID) (*customerdomain.Customer, error) {
	customerCollection := c.database.Collection(c.customerCollection)

	filter := bson.M{"_id": id}

	var customer *customerdomain.Customer
	err := customerCollection.FindOne(ctx, filter).Decode(&customer)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return customer, nil
}

func (c *customerRepository) GetOneByName(ctx context.Context, name string) (*customerdomain.Customer, error) {
	customerCollection := c.database.Collection(c.customerCollection)

	filter := bson.M{"name": name}

	var customer *customerdomain.Customer
	err := customerCollection.FindOne(ctx, filter).Decode(&customer)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return customer, nil
}

func (c *customerRepository) GetAll(ctx context.Context) ([]customerdomain.Customer, error) {
	customerCollection := c.database.Collection(c.customerCollection)

	filter := bson.M{}
	cursor, err := customerCollection.Find(ctx, filter)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	var customers []customerdomain.Customer
	customers = make([]customerdomain.Customer, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var customer customerdomain.Customer
		if err = cursor.Decode(&customer); err != nil {
			return nil, err
		}

		customers = append(customers, customer)
	}

	return customers, nil
}
