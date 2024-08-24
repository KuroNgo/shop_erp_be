package data_seeder

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	roledomain "shop_erp_mono/domain/human_resource_management/role"
	salarydomain "shop_erp_mono/domain/human_resource_management/salary"
	"time"
)

var salary = salarydomain.Input{
	Role:         "Admin",
	BaseSalary:   1500.00,
	UnitCurrency: "USD",
	Bonus:        200.00,
	Deductions:   100.00,
	NetSalary:    1600.00,
}

func SeedSalary(ctx context.Context, client *mongo.Client) error {
	collectionSalary := client.Database("shopERP").Collection("salary")
	collectionRole := client.Database("shopERP").Collection("role")

	filter := bson.M{"title": salary.Role}
	var role roledomain.Role
	if err := collectionRole.FindOne(ctx, filter).Decode(&role); err != nil {
		return err
	}

	salaryData := salarydomain.Salary{
		ID:           primitive.NewObjectID(),
		RoleID:       role.ID,
		UnitCurrency: salary.UnitCurrency,
		BaseSalary:   salary.BaseSalary,
		Bonus:        salary.Bonus,
		Deductions:   salary.Deductions,
		NetSalary:    salary.NetSalary,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	count, err := collectionSalary.CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}

	if count == 0 {
		_, err = collectionSalary.InsertOne(ctx, salaryData)
		if err != nil {
			return err
		}
	}

	return nil
}
