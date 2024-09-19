package benefit_repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	benefitsdomain "shop_erp_mono/domain/human_resource_management/benefits"
	"time"
)

type benefitRepository struct {
	database          *mongo.Database
	collectionBenefit string
}

func NewBenefitRepository(db *mongo.Database, collectionBenefit string) benefitsdomain.IBenefitRepository {
	return &benefitRepository{database: db, collectionBenefit: collectionBenefit}
}

func (b *benefitRepository) CreateOne(ctx context.Context, benefit *benefitsdomain.Benefit) error {
	collectionBenefit := b.database.Collection(b.collectionBenefit)

	_, err := collectionBenefit.InsertOne(ctx, benefit)
	if err != nil {
		return errors.New(err.Error() + "error the inserting benefit's information into database")
	}

	return nil
}

func (b *benefitRepository) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	collectionBenefit := b.database.Collection(b.collectionBenefit)

	if id == primitive.NilObjectID {
		return errors.New("id do not null")
	}

	filter := bson.M{"_id": id}

	_, err := collectionBenefit.DeleteOne(ctx, filter)
	if err != nil {
		return errors.New(err.Error() + "error the deleting benefit's information into database")
	}

	return nil
}

func (b *benefitRepository) UpdateOne(ctx context.Context, id primitive.ObjectID, benefit *benefitsdomain.Benefit) error {
	collectionBenefit := b.database.Collection(b.collectionBenefit)

	if id == primitive.NilObjectID {
		return errors.New("id do not null")
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"employee_id":  benefit.EmployeeID,
		"benefit_type": benefit.BenefitType,
		"amount":       benefit.Amount,
		"start_date":   benefit.StartDate,
		"end_date":     benefit.EndDate,
		"updated_at":   time.Now(),
	}}

	_, err := collectionBenefit.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New(err.Error() + "error the updating benefit's information into database")
	}

	return nil
}

func (b *benefitRepository) GetByID(ctx context.Context, id primitive.ObjectID) (benefitsdomain.Benefit, error) {
	collectionBenefit := b.database.Collection(b.collectionBenefit)

	if id == primitive.NilObjectID {
		return benefitsdomain.Benefit{}, errors.New("id do not nil")
	}

	var benefit benefitsdomain.Benefit
	filter := bson.M{"_id": id}
	if err := collectionBenefit.FindOne(ctx, filter).Decode(&benefit); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return benefitsdomain.Benefit{}, nil
		}
		return benefitsdomain.Benefit{}, err
	}

	return benefit, nil
}

func (b *benefitRepository) GetByEmployeeID(ctx context.Context, employeeID primitive.ObjectID) (benefitsdomain.Benefit, error) {
	collectionBenefit := b.database.Collection(b.collectionBenefit)

	var benefit benefitsdomain.Benefit
	filter := bson.M{"employee_id": employeeID}
	if err := collectionBenefit.FindOne(ctx, filter).Decode(&benefit); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return benefitsdomain.Benefit{}, nil
		}
		return benefitsdomain.Benefit{}, err
	}

	return benefit, nil
}

func (b *benefitRepository) GetAll(ctx context.Context) ([]benefitsdomain.Benefit, error) {
	collectionBenefit := b.database.Collection(b.collectionBenefit)

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

	var benefits []benefitsdomain.Benefit
	benefits = make([]benefitsdomain.Benefit, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var benefit benefitsdomain.Benefit
		if err = cursor.Decode(&benefit); err != nil {
			return nil, err
		}

		benefits = append(benefits, benefit)
	}

	// Check for any errors encountered during iteration
	if err = cursor.Err(); err != nil {
		return nil, err
	}
	return benefits, nil
}
