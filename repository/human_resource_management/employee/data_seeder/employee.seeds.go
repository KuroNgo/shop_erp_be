package data_seeder

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	employeesdomain "shop_erp_mono/domain/human_resource_management/employees"
	"time"
)

var employee = employeesdomain.Employee{
	ID:        primitive.NewObjectID(),
	FirstName: "admin",
	LastName:  "",
	Email:     "admin@admin.com",
	Gender:    "Male",
	Phone:     "0329245971",
	Address:   "HCM city",
	IsActive:  true,
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

func SeedEmployee(ctx context.Context, client *mongo.Client) error {
	collectionEmployee := client.Database("shopERP").Collection("employee")

	count, err := collectionEmployee.CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}

	if count == 0 {
		_, err = collectionEmployee.InsertOne(ctx, employee)
		if err != nil {
			return err
		}
	}

	return nil
}
