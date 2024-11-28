package user_seeder

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	userdomain "shop_erp_mono/internal/domain/human_resource_management/user"
	"shop_erp_mono/pkg/shared/constant"
	"shop_erp_mono/pkg/shared/password"
	"time"
)

var user = userdomain.User{
	ID:           primitive.NewObjectID(),
	Username:     "admin",
	Email:        "admin@admin.com",
	PasswordHash: "12345",
	Phone:        "0329245971",
	Verified:     true,
	Provider:     "app",
	Role:         constant.RoleSuperAdmin,
	CreatedAt:    time.Now(),
	UpdatedAt:    time.Now(),
}

func SeedUser(ctx context.Context, client *mongo.Client) error {
	collectionUser := client.Database("shopERP").Collection("user")

	count, err := collectionUser.CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}

	user.PasswordHash, err = password.HashPassword(user.PasswordHash)
	if err != nil {
		return err
	}

	if count == 0 {
		_, err = collectionUser.InsertOne(ctx, user)
		if err != nil {
			return err
		}
	}

	return nil
}
