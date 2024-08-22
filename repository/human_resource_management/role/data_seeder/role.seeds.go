package data_seeder

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	roledomain "shop_erp_mono/domain/human_resource_management/role"
	"time"
)

var role = roledomain.Role{
	ID:          primitive.NewObjectID(),
	Title:       "admin",
	Description: "admin have a role base in this system",
	CreatedAt:   time.Now(),
	UpdatedAt:   time.Now(),
}

func SeedRole(ctx context.Context, client *mongo.Client) error {
	collectionUser := client.Database("shopERP").Collection("role")

	count, err := collectionUser.CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}

	if count == 0 {
		_, err = collectionUser.InsertOne(ctx, role)
		if err != nil {
			return err
		}
	}

	return nil
}
