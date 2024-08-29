package department_repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	departmentsdomain "shop_erp_mono/domain/human_resource_management/departments"
	departmentvalidate "shop_erp_mono/repository/human_resource_management/department/validate"
	"time"
)

type departmentRepository struct {
	collectionDepartment string
	database             *mongo.Database
}

func NewDepartmentRepository(db *mongo.Database, collectionDepartment string) departmentsdomain.IDepartmentRepository {
	return &departmentRepository{database: db, collectionDepartment: collectionDepartment}
}

func (d departmentRepository) CreateOne(ctx context.Context, input *departmentsdomain.Input) error {
	collectionDepartment := d.database.Collection(d.collectionDepartment)

	if err := departmentvalidate.IsNilDepartment(input); err != nil {
		return err
	}

	department := departmentsdomain.Department{
		ID:          primitive.NewObjectID(),
		Name:        input.Name,
		Description: input.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	_, err := collectionDepartment.InsertOne(ctx, department)
	if err != nil {
		return errors.New(err.Error() + "error in the inserting information's department into database")
	}

	return nil
}

func (d departmentRepository) DeleteOne(ctx context.Context, id string) error {
	collectionDepartment := d.database.Collection(d.collectionDepartment)

	departmentID, _ := primitive.ObjectIDFromHex(id)

	if departmentID == primitive.NilObjectID {
		return errors.New("error in the department's ID with delete in database, this is do not nil")
	}

	filter := bson.M{"_id": departmentID}
	_, err := collectionDepartment.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (d departmentRepository) UpdateOne(ctx context.Context, id string, input *departmentsdomain.Input) error {
	collectionDepartment := d.database.Collection(d.collectionDepartment)

	if err := departmentvalidate.IsNilDepartment(input); err != nil {
		return err
	}

	departmentID, _ := primitive.ObjectIDFromHex(id)
	if departmentID == primitive.NilObjectID {
		return errors.New("id do not nil")
	}

	filter := bson.M{"_id": departmentID}
	update := bson.M{"$set": bson.M{
		"name":        input.Name,
		"description": input.Description,
		"updated_at":  time.Now(),
	}}

	_, err := collectionDepartment.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (d departmentRepository) GetOneByID(ctx context.Context, id string) (departmentsdomain.Output, error) {
	collectionDepartment := d.database.Collection(d.collectionDepartment)

	departmentID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": departmentID}
	var department departmentsdomain.Department
	if err := collectionDepartment.FindOne(ctx, filter).Decode(&department); err != nil {
		return departmentsdomain.Output{}, err
	}

	output := departmentsdomain.Output{
		Department: department,
	}

	return output, nil
}

func (d departmentRepository) GetOneByName(ctx context.Context, name string) (departmentsdomain.Output, error) {
	collectionDepartment := d.database.Collection(d.collectionDepartment)

	filter := bson.M{"name": name}
	var department departmentsdomain.Department
	if err := collectionDepartment.FindOne(ctx, filter).Decode(&department); err != nil {
		return departmentsdomain.Output{}, err
	}

	output := departmentsdomain.Output{
		Department: department,
	}
	return output, nil
}

func (d departmentRepository) GetAll(ctx context.Context) ([]departmentsdomain.Output, error) {
	collectionDepartment := d.database.Collection(d.collectionDepartment)

	filter := bson.M{}
	cursor, err := collectionDepartment.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	var departments []departmentsdomain.Output
	departments = make([]departmentsdomain.Output, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var department departmentsdomain.Department
		if err = cursor.Decode(&department); err != nil {
			return nil, err
		}

		output := departmentsdomain.Output{
			Department: department,
		}
		departments = append(departments, output)
	}

	return departments, nil
}
