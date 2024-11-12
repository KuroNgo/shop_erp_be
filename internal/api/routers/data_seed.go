package routers

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	department_seeder "shop_erp_mono/internal/repository/human_resource_management/department/data_seeder"
	user_seeder "shop_erp_mono/internal/repository/human_resource_management/user/data_seeder"
)

func DataSeeds(ctx context.Context, client *mongo.Client) error {
	err := user_seeder.SeedUser(ctx, client)
	if err != nil {
		return err
	}

	err = department_seeder.SeedDepartment(ctx, client)
	if err != nil {
		return err
	}

	return nil
}
