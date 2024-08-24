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
	Title:       "Admin",
	Description: "The admin role has full access and control over the system.",
	CreatedAt:   time.Now(),
	UpdatedAt:   time.Now(),
}

func SeedRole(ctx context.Context, client *mongo.Client) error {
	collectionRole := client.Database("shopERP").Collection("role")

	count, err := collectionRole.CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}

	if count == 0 {
		_, err = collectionRole.InsertOne(ctx, role)
		if err != nil {
			return err
		}
	}

	return nil
}
