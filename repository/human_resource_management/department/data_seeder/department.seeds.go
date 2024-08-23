package data_seeder

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	departmentsdomain "shop_erp_mono/domain/human_resource_management/departments"
	"time"
)

var (
	marketingDept = departmentsdomain.Department{
		ID:          primitive.NewObjectID(),
		Name:        "Marketing",
		Description: "Responsible for promoting and selling the company's products and services.",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	hrDept = departmentsdomain.Department{
		ID:          primitive.NewObjectID(),
		Name:        "Human Resources",
		Description: "Responsible for managing employee relations, recruitment, and company culture.",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
)

func SeedDepartment(ctx context.Context, client *mongo.Client) error {
	collectionUser := client.Database("shopERP").Collection("department")

	count, err := collectionUser.CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}

	if count == 0 {
		_, err = collectionUser.InsertOne(ctx, marketingDept)
		if err != nil {
			return err
		}
		_, err = collectionUser.InsertOne(ctx, hrDept)
		if err != nil {
			return err
		}
	}

	return nil
}
