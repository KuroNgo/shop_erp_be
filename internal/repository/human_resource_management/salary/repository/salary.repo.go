package salary_repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	salarydomain "shop_erp_mono/internal/domain/human_resource_management/salary"
	"time"
)

type salaryRepository struct {
	database         *mongo.Database
	collectionSalary string
}

func NewSalaryRepository(db *mongo.Database, collectionSalary string) salarydomain.ISalaryRepository {
	return &salaryRepository{database: db, collectionSalary: collectionSalary}
}

func (s *salaryRepository) CreateOne(ctx context.Context, salary *salarydomain.Salary) error {
	collectionSalary := s.database.Collection(s.collectionSalary)

	filterRole := bson.M{"role_id": salary.ID}
	count, err := collectionSalary.CountDocuments(ctx, filterRole)
	if count > 0 {
		return errors.New("the value do exist in database")
	}

	_, err = collectionSalary.InsertOne(ctx, salary)
	if err != nil {
		return errors.New(err.Error() + "error in the creating salary's information into database")
	}

	return nil
}

func (s *salaryRepository) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	collectionSalary := s.database.Collection(s.collectionSalary)

	if id == primitive.NilObjectID {
		return errors.New("error in the deleting with value nil")
	}

	filter := bson.M{"_id": id}
	_, err := collectionSalary.DeleteOne(ctx, filter)
	if err != nil {
		return errors.New(err.Error() + "error in the deleting salary's information into database")
	}

	return nil
}

func (s *salaryRepository) UpdateOne(ctx context.Context, id primitive.ObjectID, salary *salarydomain.Salary) error {
	collectionSalary := s.database.Collection(s.collectionSalary)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"unit_currency": salary.UnitCurrency,
		"base_salary":   salary.BaseSalary,
		"bonus":         salary.Bonus,
		"deductions":    salary.Deductions,
		"net_salary":    salary.NetSalary,
		"updated_at":    time.Now(),
	}}

	_, err := collectionSalary.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New(err.Error() + "error in the creating salary's information into database")
	}

	return nil
}

func (s *salaryRepository) GetByID(ctx context.Context, id primitive.ObjectID) (salarydomain.Salary, error) {
	collectionSalary := s.database.Collection(s.collectionSalary)

	filter := bson.M{"_id": id}
	var salary salarydomain.Salary
	if err := collectionSalary.FindOne(ctx, filter).Decode(&salary); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return salarydomain.Salary{}, nil
		}
		return salarydomain.Salary{}, errors.New(err.Error() + "error in the finding salary's information into database")
	}

	return salary, nil
}

func (s *salaryRepository) GetByRoleID(ctx context.Context, roleID primitive.ObjectID) (salarydomain.Salary, error) {
	collectionSalary := s.database.Collection(s.collectionSalary)

	filter := bson.M{"role_id": roleID}
	var salary salarydomain.Salary
	if err := collectionSalary.FindOne(ctx, filter).Decode(&salary); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return salarydomain.Salary{}, nil
		}
		return salarydomain.Salary{}, errors.New(err.Error() + "error in the finding salary's information into database")
	}

	return salary, nil
}

func (s *salaryRepository) GetAll(ctx context.Context) ([]salarydomain.Salary, error) {
	collectionSalary := s.database.Collection(s.collectionSalary)

	filter := bson.M{}
	cursor, err := collectionSalary.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	var salaries []salarydomain.Salary
	salaries = make([]salarydomain.Salary, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var salary salarydomain.Salary
		if err = cursor.Decode(&salary); err != nil {
			return nil, err
		}

		salaries = append(salaries, salary)
	}

	// Check for any errors encountered during iteration
	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return salaries, nil
}

func (s *salaryRepository) CountSalary(ctx context.Context) (int64, error) {
	collectionSalary := s.database.Collection(s.collectionSalary)

	filter := bson.M{}
	count, err := collectionSalary.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}

	return count, nil
}
