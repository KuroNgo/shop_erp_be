package user_repository

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	userdomain "shop_erp_mono/internal/domain/human_resource_management/user"
)

type userRepository struct {
	database       *mongo.Database
	collectionUser string
}

func NewUserRepository(db *mongo.Database, collectionUser string) userdomain.IUserRepository {
	return &userRepository{database: db, collectionUser: collectionUser}
}

func (r *userRepository) FetchMany(ctx context.Context) ([]userdomain.User, error) {
	collectionUser := r.database.Collection(r.collectionUser)

	filter := bson.M{}
	cursor, err := collectionUser.Find(ctx, filter)
	if err != nil {
		return nil, errors.New(err.Error() + "error in the finding user into the database")
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	var users []userdomain.User
	users = make([]userdomain.User, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var user userdomain.User
		if err = cursor.Decode(&user); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	// Check for any errors encountered during iteration
	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (userdomain.User, error) {
	collectionUser := r.database.Collection(r.collectionUser)

	filter := bson.M{"email": email}
	var user userdomain.User
	if err := collectionUser.FindOne(ctx, filter).Decode(&user); err != nil {
		return userdomain.User{}, errors.New(err.Error() + "error in the finding user into the database")
	}

	return user, nil
}

func (r *userRepository) GetByID(ctx context.Context, id primitive.ObjectID) (userdomain.User, error) {
	collectionUser := r.database.Collection(r.collectionUser)

	filter := bson.M{"_id": id}
	projection := bson.M{"password_hash": 0}

	var user userdomain.User
	if err := collectionUser.FindOne(ctx, filter, options.FindOne().SetProjection(projection)).Decode(&user); err != nil {
		return userdomain.User{}, errors.New(err.Error() + "error in the finding user into the database")
	}

	return user, nil
}

func (r *userRepository) GetByVerificationCode(ctx context.Context, verificationCode string) (userdomain.User, error) {
	collectionUser := r.database.Collection(r.collectionUser)

	filter := bson.M{"verification_code": verificationCode}
	var user userdomain.User
	if err := collectionUser.FindOne(ctx, filter).Decode(&user); err != nil {
		return userdomain.User{}, errors.New(err.Error() + "error in the finding user's data into database")
	}

	return user, nil
}

func (r *userRepository) UserExists(ctx context.Context, email string) (bool, error) {
	collectionUser := r.database.Collection(r.collectionUser)
	filter := bson.M{"email": email}
	count, err := collectionUser.CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *userRepository) CreateOne(ctx context.Context, user *userdomain.User) error {
	collectionUser := r.database.Collection(r.collectionUser)

	exists, err := r.UserExists(ctx, user.Email)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("user is confirmed to exist")
	}

	_, err = collectionUser.InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("error inserting user into the database: %w", err)
	}

	return nil
}

func (r *userRepository) UpdateOne(ctx context.Context, user *userdomain.User) error {
	collectionUser := r.database.Collection(r.collectionUser)

	exists, err := r.UserExists(ctx, user.Email)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("user is confirmed to exist")
	}

	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": user}

	_, err = collectionUser.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New(err.Error() + "error in the updating user into database")
	}

	return nil
}

func (r *userRepository) UpdatePassword(ctx context.Context, user *userdomain.User) error {
	collectionUser := r.database.Collection(r.collectionUser)

	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": bson.M{"password_hash": user.PasswordHash}}

	_, err := collectionUser.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New(err.Error() + "error in the updating user into database")
	}

	return nil

}

func (r *userRepository) UpdateVerify(ctx context.Context, user *userdomain.User) error {
	collectionUser := r.database.Collection(r.collectionUser)

	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": bson.M{
		"verify":     user.Verified,
		"updated_at": user.UpdatedAt,
	}}

	_, err := collectionUser.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New(err.Error() + "error in the updating user's data into database")
	}

	return nil
}

func (r *userRepository) UpdateVerificationCode(ctx context.Context, user *userdomain.User) error {
	collectionUser := r.database.Collection(r.collectionUser)

	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": bson.M{
		"verify":            user.Verified,
		"verification_code": user.VerificationCode,
		"updated_at":        user.UpdatedAt,
	}}

	_, err := collectionUser.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New(err.Error() + "error in the updating user's data into database")
	}

	return nil
}

func (r *userRepository) UpsertOne(ctx context.Context, user *userdomain.User) (*userdomain.User, error) {
	collectionUser := r.database.Collection(r.collectionUser)

	filter := bson.M{"email": user.Email}

	// Chuẩn bị các giá trị cập nhật
	update := bson.D{{Key: "$set", Value: bson.M{
		"full_name":  user.Username,
		"email":      user.Email,
		"avatar_url": user.AvatarURL,
		"phone":      user.Phone,
		"provider":   user.Provider,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
		"role":       user.Role,
	}}}

	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
	result := collectionUser.FindOneAndUpdate(ctx, filter, update, opts)

	var updatedUser *userdomain.User
	if err := result.Decode(&updatedUser); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return updatedUser, nil
}

func (r *userRepository) UpdateImage(ctx context.Context, user *userdomain.User) error {
	collectionUser := r.database.Collection(r.collectionUser)

	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": bson.M{
		"avatar_url": user.AvatarURL,
		"asset_url":  user.AssetURL,
	}}

	_, err := collectionUser.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New(err.Error() + "error in the updating user's data into database")
	}

	return nil
}

func (r *userRepository) DeleteOne(ctx context.Context, userID primitive.ObjectID) error {
	collectionUser := r.database.Collection(r.collectionUser)

	filter := bson.M{"_id": userID}
	_, err := collectionUser.DeleteOne(ctx, filter)
	if err != nil {
		return errors.New(err.Error() + "error the deleting user's data into the database")
	}

	return nil
}
