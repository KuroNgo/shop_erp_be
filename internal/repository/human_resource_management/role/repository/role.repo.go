package role_repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	roledomain "shop_erp_mono/internal/domain/human_resource_management/role"
	"time"
)

type roleRepository struct {
	database       *mongo.Database
	collectionRole string
}

func NewRoleRepository(db *mongo.Database, collectionRole string) roledomain.IRoleRepository {
	return &roleRepository{database: db, collectionRole: collectionRole}
}

func (r *roleRepository) CreateOne(ctx context.Context, role *roledomain.Role) error {
	collectionRole := r.database.Collection(r.collectionRole)

	_, err := collectionRole.InsertOne(ctx, role)
	if err != nil {
		return err
	}

	return nil
}

func (r *roleRepository) GetByTitle(ctx context.Context, title string) (roledomain.Role, error) {
	collectionRole := r.database.Collection(r.collectionRole)

	filter := bson.M{"title": title}
	var role roledomain.Role
	if err := collectionRole.FindOne(ctx, filter).Decode(&role); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return roledomain.Role{}, nil
		}
		return roledomain.Role{}, err
	}

	return role, nil
}

func (r *roleRepository) GetByID(ctx context.Context, id primitive.ObjectID) (roledomain.Role, error) {
	collectionRole := r.database.Collection(r.collectionRole)

	filter := bson.M{"_id": id}
	var role roledomain.Role
	if err := collectionRole.FindOne(ctx, filter).Decode(&role); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return roledomain.Role{}, nil
		}
		return roledomain.Role{}, err
	}

	return role, nil
}

func (r *roleRepository) GetByStatus(ctx context.Context, status string) ([]roledomain.Role, error) {
	collectionRole := r.database.Collection(r.collectionRole)

	filter := bson.M{"status": status}
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

	var roles []roledomain.Role
	roles = make([]roledomain.Role, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var role roledomain.Role
		if err = cursor.Decode(&role); err != nil {
			return nil, err
		}

		roles = append(roles, role)
	}

	// Check for any errors encountered during iteration
	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}

func (r *roleRepository) GetByName(ctx context.Context, name string) (roledomain.Role, error) {
	collectionRole := r.database.Collection(r.collectionRole)

	filter := bson.M{"name": name}
	var role roledomain.Role
	if err := collectionRole.FindOne(ctx, filter).Decode(&role); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return roledomain.Role{}, nil
		}
		return roledomain.Role{}, err
	}

	return role, nil
}

func (r *roleRepository) GetByLevel(ctx context.Context, level int) ([]roledomain.Role, error) {
	collectionRole := r.database.Collection(r.collectionRole)

	filter := bson.M{"level": level}
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

	var roles []roledomain.Role
	roles = make([]roledomain.Role, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var role roledomain.Role
		if err = cursor.Decode(&role); err != nil {
			return nil, err
		}

		roles = append(roles, role)
	}

	// Check for any errors encountered during iteration
	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}

func (r *roleRepository) GetByEnable(ctx context.Context, enable int) ([]roledomain.Role, error) {
	collectionRole := r.database.Collection(r.collectionRole)

	filter := bson.M{"enable": enable}
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

	var roles []roledomain.Role
	roles = make([]roledomain.Role, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var role roledomain.Role
		if err = cursor.Decode(&role); err != nil {
			return nil, err
		}

		roles = append(roles, role)
	}

	// Check for any errors encountered during iteration
	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}

func (r *roleRepository) GetAll(ctx context.Context) ([]roledomain.Role, error) {
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

	var roles []roledomain.Role
	roles = make([]roledomain.Role, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var role roledomain.Role
		if err = cursor.Decode(&role); err != nil {
			return nil, err
		}

		roles = append(roles, role)
	}

	// Check for any errors encountered during iteration
	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}

func (r *roleRepository) UpdateOne(ctx context.Context, id primitive.ObjectID, role *roledomain.Role) error {
	collectionRole := r.database.Collection(r.collectionRole)

	if id == primitive.NilObjectID {
		return errors.New("id do not nil")
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"name":        role.Name,
		"description": role.Description,
		"updated_at":  time.Now(),
	}}
	_, err := collectionRole.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New(err.Error() + "error in the updating role's information into database ")
	}

	return nil
}

func (r *roleRepository) UpdateStatus(ctx context.Context, id primitive.ObjectID, status string) error {
	collectionRole := r.database.Collection(r.collectionRole)

	if id == primitive.NilObjectID {
		return errors.New("id do not nil")
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"status":     status,
		"updated_at": time.Now(),
	}}
	_, err := collectionRole.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New(err.Error() + "error in the updating role's information into database ")
	}

	return nil
}

func (r *roleRepository) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	collectionRole := r.database.Collection(r.collectionRole)

	filter := bson.M{"_id": id}
	_, err := collectionRole.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (r *roleRepository) DeleteSoft(ctx context.Context, id primitive.ObjectID) error {
	collectionRole := r.database.Collection(r.collectionRole)

	if id == primitive.NilObjectID {
		return errors.New("id do not nil")
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"enable":     0,
		"updated_at": time.Now(),
	}}
	_, err := collectionRole.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New(err.Error() + "error in the updating role's information into database ")
	}

	return nil
}

func (r *roleRepository) CountRole(ctx context.Context) (int64, error) {
	collectionRole := r.database.Collection(r.collectionRole)

	filter := bson.M{}
	count, err := collectionRole.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}

	return count, nil
}
