package user_repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	userdomain "shop_erp_mono/domain/human_resource_management/user"
	"sync"
	"time"
)

type userRepository struct {
	database       *mongo.Database
	collectionUser string
}

func NewUserRepository(db *mongo.Database, collectionUser string) userdomain.IUserRepository {
	return &userRepository{database: db, collectionUser: collectionUser}
}

var (
	wg sync.WaitGroup
	mu sync.Mutex
)

func (u userRepository) FetchMany(ctx context.Context) ([]userdomain.User, error) {
	collectionUser := u.database.Collection(u.collectionUser)

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

	return users, nil
}

func (u userRepository) GetByEmail(ctx context.Context, email string) (userdomain.User, error) {
	collectionUser := u.database.Collection(u.collectionUser)

	filter := bson.M{"email": email}
	var user userdomain.User
	if err := collectionUser.FindOne(ctx, filter).Decode(&user); err != nil {
		return userdomain.User{}, errors.New(err.Error() + "error in the finding user into the database")
	}

	return user, nil
}

func (u userRepository) GetByID(ctx context.Context, id primitive.ObjectID) (userdomain.User, error) {
	collectionUser := u.database.Collection(u.collectionUser)

	filter := bson.M{"_id": id}
	var user userdomain.User
	if err := collectionUser.FindOne(ctx, filter).Decode(&user); err != nil {
		return userdomain.User{}, errors.New(err.Error() + "error in the finding user into the database")
	}

	return user, nil
}

func (u userRepository) GetByVerificationCode(ctx context.Context, verificationCode string) (userdomain.User, error) {
	collectionUser := u.database.Collection(u.collectionUser)

	filter := bson.M{"verification_code": verificationCode}
	var user userdomain.User
	if err := collectionUser.FindOne(ctx, filter).Decode(&user); err != nil {
		return userdomain.User{}, errors.New(err.Error() + "error in the finding user's data into database")
	}

	return user, nil
}

func (u userRepository) CreateOne(ctx context.Context, user *userdomain.User) error {
	collectionUser := u.database.Collection(u.collectionUser)

	_, err := collectionUser.InsertOne(ctx, user)
	if err != nil {
		return errors.New(err.Error() + "error in the inserting user into the database ")
	}

	return nil
}

func (u userRepository) UpdateOne(ctx context.Context, user *userdomain.User) error {
	collectionUser := u.database.Collection(u.collectionUser)

	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": user}

	_, err := collectionUser.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New(err.Error() + "error in the updating user into database")
	}

	return nil
}

func (u userRepository) UpdatePassword(ctx context.Context, user *userdomain.User) error {
	collectionUser := u.database.Collection(u.collectionUser)

	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": bson.M{"password_hash": user.PasswordHash}}

	_, err := collectionUser.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New(err.Error() + "error in the updating user into database")
	}

	return nil

}

func (u userRepository) UpdateVerify(ctx context.Context, user *userdomain.User) error {
	collectionUser := u.database.Collection(u.collectionUser)

	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": bson.M{"verify": user.Verified}}

	_, err := collectionUser.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New(err.Error() + "error in the updating user's data into database")
	}

	return nil
}

func (u userRepository) UpsertOne(ctx context.Context, user *userdomain.User) error {
	collectionUser := u.database.Collection(u.collectionUser)

	filter := bson.M{"email": user.Email}
	update := bson.M{"$set": bson.M{
		"username":   user.Username,
		"updated_at": time.Now(),
	}, "$setOnInsert": bson.M{
		"createdAt": time.Now(),
	}}

	opts := options.Update().SetUpsert(true)
	_, err := collectionUser.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return errors.New(err.Error() + "error in the updating user's data into database")
	}

	return nil
}

func (u userRepository) UpdateImage(ctx context.Context, user *userdomain.User) error {
	collectionUser := u.database.Collection(u.collectionUser)

	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": bson.M{"avatar_url": user.AvatarURL}}

	_, err := collectionUser.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New(err.Error() + "error in the updating user's data into database")
	}

	return nil
}

func (u userRepository) DeleteOne(ctx context.Context, userID primitive.ObjectID) error {
	collectionUser := u.database.Collection(u.collectionUser)

	filter := bson.M{"_id": userID}

	_, err := collectionUser.DeleteOne(ctx, filter)
	if err != nil {
		return errors.New(err.Error() + "error the deleting user's data into the database")
	}

	return nil
}
