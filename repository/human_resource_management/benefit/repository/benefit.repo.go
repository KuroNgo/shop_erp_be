package benefit_repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	benefitsdomain "shop_erp_mono/domain/human_resource_management/benefits"
	employeesdomain "shop_erp_mono/domain/human_resource_management/employees"
	"shop_erp_mono/repository/human_resource_management/benefit/validate"
	"time"
)

type benefitRepository struct {
	database           *mongo.Database
	collectionBenefit  string
	collectionEmployee string
}

func NewBenefitRepository(db *mongo.Database, collectionBenefit string, collectionEmployee string) benefitsdomain.IBenefitRepository {
	return &benefitRepository{database: db, collectionBenefit: collectionBenefit, collectionEmployee: collectionEmployee}
}

func (b benefitRepository) CreateOne(ctx context.Context, input *benefitsdomain.Input) error {
	collectionBenefit := b.database.Collection(b.collectionBenefit)
	collectionEmployee := b.database.Collection(b.collectionEmployee)

	if err := validate.IsNilBenefit(input); err != nil {
		return err
	}

	var employee employeesdomain.Employee
	filterEmployee := bson.M{"email": input.EmployeeEmail}
	if err := collectionEmployee.FindOne(ctx, filterEmployee).Decode(&employee); err != nil {
		return err
	}

	benefit := benefitsdomain.Benefit{
		ID:          primitive.NewObjectID(),
		EmployeeID:  employee.ID,
		BenefitType: input.BenefitType,
		Amount:      input.Amount,
		StartDate:   input.StartDate,
		EndDate:     input.EndDate,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	_, err := collectionBenefit.InsertOne(ctx, benefit)
	if err != nil {
		return errors.New(err.Error() + "error the inserting benefit's information into database")
	}

	return nil
}

func (b benefitRepository) DeleteOne(ctx context.Context, id string) error {
	collectionBenefit := b.database.Collection(b.collectionBenefit)

	benefitID, _ := primitive.ObjectIDFromHex(id)
	if benefitID == primitive.NilObjectID {
		return errors.New("id do not null")
	}

	filter := bson.M{"_id": benefitID}

	_, err := collectionBenefit.DeleteOne(ctx, filter)
	if err != nil {
		return errors.New(err.Error() + "error the deleting benefit's information into database")
	}

	return nil
}

func (b benefitRepository) UpdateOne(ctx context.Context, id string, input *benefitsdomain.Input) error {
	collectionBenefit := b.database.Collection(b.collectionBenefit)
	collectionEmployee := b.database.Collection(b.collectionEmployee)

	benefitID, _ := primitive.ObjectIDFromHex(id)
	if benefitID == primitive.NilObjectID {
		return errors.New("id do not null")
	}

	var employee employeesdomain.Employee
	filterEmployee := bson.M{"email": input.EmployeeEmail}
	if err := collectionEmployee.FindOne(ctx, filterEmployee).Decode(&employee); err != nil {
		return err
	}

	filter := bson.M{"_id": benefitID}
	update := bson.M{"$set": bson.M{
		"employee_id":  employee.ID,
		"benefit_type": input.BenefitType,
		"amount":       input.Amount,
		"start_date":   input.StartDate,
		"end_date":     input.EndDate,
	}}

	_, err := collectionBenefit.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New(err.Error() + "error the updating benefit's information into database")
	}

	return nil
}

func (b benefitRepository) GetOneByID(ctx context.Context, id string) (benefitsdomain.Output, error) {
	collectionBenefit := b.database.Collection(b.collectionBenefit)
	collectionEmployee := b.database.Collection(b.collectionEmployee)

	benefitID, _ := primitive.ObjectIDFromHex(id)
	if benefitID == primitive.NilObjectID {
		return benefitsdomain.Output{}, errors.New("id do not nil")
	}

	var benefit benefitsdomain.Benefit
	filter := bson.M{"_id": benefitID}
	if err := collectionBenefit.FindOne(ctx, filter).Decode(&benefit); err != nil {
		return benefitsdomain.Output{}, err
	}

	var employee employeesdomain.Employee
	filterEmployee := bson.M{"_id": benefit.EmployeeID}
	if err := collectionEmployee.FindOne(ctx, filterEmployee).Decode(&employee); err != nil {
		return benefitsdomain.Output{}, err
	}

	output := benefitsdomain.Output{
		Benefit:  benefit,
		Employee: employee.Email,
	}

	return output, nil
}

func (b benefitRepository) GetOneByEmail(ctx context.Context, email string) (benefitsdomain.Output, error) {
	collectionBenefit := b.database.Collection(b.collectionBenefit)
	collectionEmployee := b.database.Collection(b.collectionEmployee)

	var employee employeesdomain.Employee
	filterEmployee := bson.M{"email": email}
	if err := collectionEmployee.FindOne(ctx, filterEmployee).Decode(&employee); err != nil {
		return benefitsdomain.Output{}, err
	}

	var benefit benefitsdomain.Benefit
	filter := bson.M{"employee_id": employee.ID}
	if err := collectionBenefit.FindOne(ctx, filter).Decode(&benefit); err != nil {
		return benefitsdomain.Output{}, err
	}

	output := benefitsdomain.Output{
		Benefit:  benefit,
		Employee: employee.Email,
	}

	return output, nil
}

func (b benefitRepository) GetAll(ctx context.Context) ([]benefitsdomain.Output, error) {
	collectionBenefit := b.database.Collection(b.collectionBenefit)
	collectionEmployee := b.database.Collection(b.collectionEmployee)

	filter := bson.M{}
	cursor, err := collectionBenefit.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	var benefits []benefitsdomain.Output
	benefits = make([]benefitsdomain.Output, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var benefit benefitsdomain.Benefit
		if err = cursor.Decode(&benefit); err != nil {
			return nil, err
		}

		var employee employeesdomain.Employee
		filterEmployee := bson.M{"_id": benefit.EmployeeID}
		if err := collectionEmployee.FindOne(ctx, filterEmployee).Decode(&employee); err != nil {
			return nil, err
		}

		output := benefitsdomain.Output{
			Benefit:  benefit,
			Employee: employee.Email,
		}

		benefits = append(benefits, output)
	}

	return benefits, nil
}
