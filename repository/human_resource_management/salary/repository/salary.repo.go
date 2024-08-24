package salary_repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	roledomain "shop_erp_mono/domain/human_resource_management/role"
	salarydomain "shop_erp_mono/domain/human_resource_management/salary"
	"shop_erp_mono/repository/human_resource_management/salary/validate"
	"time"
)

type salaryRepository struct {
	database         *mongo.Database
	collectionSalary string
	collectionRole   string
}

func NewSalaryRepository(db *mongo.Database, collectionSalary string, collectionRole string) salarydomain.ISalaryRepository {
	return &salaryRepository{database: db, collectionSalary: collectionSalary, collectionRole: collectionRole}
}

func (s salaryRepository) CreateOne(ctx context.Context, salary *salarydomain.Input) error {
	collectionSalary := s.database.Collection(s.collectionSalary)
	collectionRole := s.database.Collection(s.collectionRole)

	if err := validate.IsNilSalary(salary); err != nil {
		return err
	}

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

	filterRole := bson.M{"role_id": role.ID}
	count, err := collectionSalary.CountDocuments(ctx, filterRole)
	if count > 0 {
		return errors.New("the value do exist in database")
	}

	_, err = collectionSalary.InsertOne(ctx, salaryData)
	if err != nil {
		return errors.New(err.Error() + "error in the creating salary's information into database")
	}

	return nil
}

func (s salaryRepository) DeleteOne(ctx context.Context, id string) error {
	collectionSalary := s.database.Collection(s.collectionSalary)

	salaryID, _ := primitive.ObjectIDFromHex(id)
	if salaryID == primitive.NilObjectID {
		return errors.New("error in the deleting with value nil")
	}

	filter := bson.M{"_id": salaryID}
	_, err := collectionSalary.DeleteOne(ctx, filter)
	if err != nil {
		return errors.New(err.Error() + "error in the deleting salary's information into database")
	}

	return nil
}

func (s salaryRepository) UpdateOne(ctx context.Context, salary *salarydomain.Input) error {
	collectionSalary := s.database.Collection(s.collectionSalary)
	collectionRole := s.database.Collection(s.collectionRole)

	if err := validate.IsNilSalary(salary); err != nil {
		return errors.New("error the updating salary's information into database")
	}

	filterRole := bson.M{"title": salary.Role}
	var role roledomain.Role
	if err := collectionRole.FindOne(ctx, filterRole).Decode(&role); err != nil {
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

	filter := bson.M{"_id": salaryData.ID}
	update := bson.M{"$set": bson.M{
		"role_id":       salaryData.RoleID,
		"unit_currency": salary.UnitCurrency,
		"base_salary":   salary.BaseSalary,
		"bonus":         salary.Bonus,
		"deductions":    salary.Deductions,
		"net_salary":    salary.NetSalary,
		"updated_at":    salaryData.UpdatedAt,
	}}

	_, err := collectionSalary.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New(err.Error() + "error in the creating salary's information into database")
	}

	return nil
}

func (s salaryRepository) GetOneByID(ctx context.Context, id string) (salarydomain.Output, error) {
	collectionSalary := s.database.Collection(s.collectionSalary)
	collectionRole := s.database.Collection(s.collectionRole)

	salaryID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": salaryID}
	var salary salarydomain.Salary
	if err := collectionSalary.FindOne(ctx, filter).Decode(&salary); err != nil {
		return salarydomain.Output{}, errors.New(err.Error() + "error in the finding salary's information into database")
	}

	filterRole := bson.M{"_id": salary.RoleID}
	var role roledomain.Role
	if err := collectionRole.FindOne(ctx, filterRole).Decode(&role); err != nil {
		return salarydomain.Output{}, errors.New(err.Error() + "error in the finding role's information into database")
	}

	output := salarydomain.Output{
		Salary: salary,
		Role:   role.Title,
	}

	return output, nil
}

func (s salaryRepository) GetOneByRole(ctx context.Context, name string) (salarydomain.Output, error) {
	collectionSalary := s.database.Collection(s.collectionSalary)
	collectionRole := s.database.Collection(s.collectionRole)

	filterRole := bson.M{"title": name}
	var role roledomain.Role
	if err := collectionRole.FindOne(ctx, filterRole).Decode(&role); err != nil {
		return salarydomain.Output{}, errors.New(err.Error() + "error in the finding role's information into database")
	}

	filter := bson.M{"role_id": role.ID}
	var salary salarydomain.Salary
	if err := collectionSalary.FindOne(ctx, filter).Decode(&salary); err != nil {
		return salarydomain.Output{}, errors.New(err.Error() + "error in the finding salary's information into database")
	}

	output := salarydomain.Output{
		Salary: salary,
		Role:   role.Title,
	}

	return output, nil
}

func (s salaryRepository) GetAll(ctx context.Context) ([]salarydomain.Output, error) {
	collectionSalary := s.database.Collection(s.collectionSalary)
	collectionRole := s.database.Collection(s.collectionRole)

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

	var salaries []salarydomain.Output
	salaries = make([]salarydomain.Output, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var salary salarydomain.Salary
		if err = cursor.Decode(&salary); err != nil {
			return nil, err
		}

		filterRole := bson.M{"_id": salary.RoleID}
		var role roledomain.Role
		if err := collectionRole.FindOne(ctx, filterRole).Decode(&role); err != nil {
			return nil, errors.New(err.Error() + "error in the finding role's information into database")
		}

		output := salarydomain.Output{
			Salary: salary,
			Role:   role.Title,
		}

		salaries = append(salaries, output)
	}

	return salaries, nil
}
