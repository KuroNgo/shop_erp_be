package base_salary_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	basesalarydomain "shop_erp_mono/internal/domain/human_resource_management/salary_base"
)

type baseSalaryRepository struct {
	database             *mongo.Database
	baseSalaryCollection string
}

func NewBaseSalaryRepository(database *mongo.Database, baseSalaryCollection string) basesalarydomain.ISalaryRepository {
	return &baseSalaryRepository{database: database, baseSalaryCollection: baseSalaryCollection}
}

func (b baseSalaryRepository) CreateOne(ctx context.Context, baseSalary *basesalarydomain.BaseSalary) error {
	collectionBaseSalary := b.database.Collection(b.baseSalaryCollection)

	_, err := collectionBaseSalary.InsertOne(ctx, baseSalary)
	if err != nil {
		return err
	}

	return nil
}

func (b baseSalaryRepository) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	collectionBaseSalary := b.database.Collection(b.baseSalaryCollection)

	filter := bson.M{"_id": id}
	_, err := collectionBaseSalary.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (b baseSalaryRepository) UpdateOne(ctx context.Context, id primitive.ObjectID, baseSalary *basesalarydomain.BaseSalary) error {
	collectionBaseSalary := b.database.Collection(b.baseSalaryCollection)

	filter := bson.M{"_id": id}
	update := bson.M{
		"role_id":       baseSalary.RoleID,
		"unit_currency": baseSalary.UnitCurrency,
		"base_salary":   baseSalary.BaseSalaries,
		"updated_at":    baseSalary.UpdatedAt,
	}

	_, err := collectionBaseSalary.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (b baseSalaryRepository) GetByID(ctx context.Context, id primitive.ObjectID) (basesalarydomain.BaseSalary, error) {
	collectionBaseSalary := b.database.Collection(b.baseSalaryCollection)

	filter := bson.M{"_id": id}
	var baseSalary basesalarydomain.BaseSalary
	err := collectionBaseSalary.FindOne(ctx, filter).Decode(&baseSalary)
	if err != nil {
		return basesalarydomain.BaseSalary{}, err
	}

	return baseSalary, nil
}

func (b baseSalaryRepository) GetByRoleID(ctx context.Context, roleID primitive.ObjectID) (basesalarydomain.BaseSalary, error) {
	collectionBaseSalary := b.database.Collection(b.baseSalaryCollection)

	filter := bson.M{"role_id": roleID}
	var baseSalary basesalarydomain.BaseSalary
	err := collectionBaseSalary.FindOne(ctx, filter).Decode(&baseSalary)
	if err != nil {
		return basesalarydomain.BaseSalary{}, err
	}

	return baseSalary, nil
}

func (b baseSalaryRepository) GetAll(ctx context.Context) ([]basesalarydomain.BaseSalary, error) {
	collectionBaseSalary := b.database.Collection(b.baseSalaryCollection)

	filter := bson.M{}
	cursor, err := collectionBaseSalary.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	var salaries []basesalarydomain.BaseSalary
	salaries = make([]basesalarydomain.BaseSalary, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var salary basesalarydomain.BaseSalary
		if err = cursor.Decode(&salary); err != nil {
			return nil, err
		}

		salaries = append(salaries, salary)
	}

	// Check for any errors encountered during iteration
	if err = cursor.Err(); err != nil {
		return nil, err
	}
	return salaries, err
}
