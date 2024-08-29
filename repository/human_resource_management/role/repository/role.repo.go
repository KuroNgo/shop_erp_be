package role_repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	roledomain "shop_erp_mono/domain/human_resource_management/role"
	"shop_erp_mono/repository/human_resource_management/role/validate"
	"time"
)

type roleRepository struct {
	database       *mongo.Database
	collectionRole string
}

func NewRoleRepository(db *mongo.Database, collectionRole string) roledomain.IRoleRepository {
	return &roleRepository{database: db, collectionRole: collectionRole}
}

func (r roleRepository) CreateOneRole(ctx context.Context, input *roledomain.Input) error {
	collectionRole := r.database.Collection(r.collectionRole)

	err := validate.IsNilTitle(input.Title)
	if err != nil {
		return err
	}

	err = validate.IsNilDescription(input.Description)
	if err != nil {
		return err
	}

	_, err = collectionRole.InsertOne(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (r roleRepository) GetByTitleRole(ctx context.Context, title string) (roledomain.Output, error) {
	collectionRole := r.database.Collection(r.collectionRole)

	filter := bson.M{"title": title}
	var role roledomain.Role
	if err := collectionRole.FindOne(ctx, filter).Decode(&role); err != nil {
		return roledomain.Output{}, err
	}

	output := roledomain.Output{
		Role: role,
	}
	return output, nil
}

func (r roleRepository) GetByIDRole(ctx context.Context, id string) (roledomain.Output, error) {
	collectionRole := r.database.Collection(r.collectionRole)

	roleID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return roledomain.Output{}, err
	}

	filter := bson.M{"_id": roleID}
	var role roledomain.Role
	if err := collectionRole.FindOne(ctx, filter).Decode(&role); err != nil {
		return roledomain.Output{}, err
	}

	output := roledomain.Output{
		Role: role,
	}
	return output, nil
}

func (r roleRepository) GetAllRole(ctx context.Context) ([]roledomain.Output, error) {
	collectionRole := r.database.Collection(r.collectionRole)

	filter := bson.M{}
	cursor, err := collectionRole.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	var roles []roledomain.Output
	roles = make([]roledomain.Output, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var role roledomain.Role
		if err = cursor.Decode(&role); err != nil {
			return nil, err
		}

		output := roledomain.Output{
			Role: role,
		}
		roles = append(roles, output)
	}

	return roles, nil
}

func (r roleRepository) UpdateOneRole(ctx context.Context, id string, input *roledomain.Input) error {
	collectionRole := r.database.Collection(r.collectionRole)

	if err := validate.IsNilTitle(input.Title); err != nil {
		return err
	}

	if err := validate.IsNilDescription(input.Description); err != nil {
		return err
	}

	roleID, _ := primitive.ObjectIDFromHex(id)
	if roleID == primitive.NilObjectID {
		return errors.New("id do not nil")
	}

	filter := bson.M{"_id": roleID}
	update := bson.M{"$set": bson.M{
		"title":       input.Title,
		"description": input.Description,
		"updated_at":  time.Now(),
	}}
	_, err := collectionRole.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New(err.Error() + "error in the updating role's information into database ")
	}

	return nil
}

func (r roleRepository) DeleteOneRole(ctx context.Context, id string) error {
	collectionRole := r.database.Collection(r.collectionRole)

	if err := validate.IsNilID(id); err != nil {
		return err
	}

	roleID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": roleID}
	_, err := collectionRole.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
