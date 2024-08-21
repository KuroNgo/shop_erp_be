package user_repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	userdomain "shop_erp_mono/domain/human_resource_management/user"
	"shop_erp_mono/pkg/password"
	user_validate "shop_erp_mono/repository/human_resource_management/user/validate"
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

func (u userRepository) FetchMany(ctx context.Context) (userdomain.Response, error) {
	collectionUser := u.database.Collection(u.collectionUser)

	errCh := make(chan error)
	filter := bson.M{}
	cursor, err := collectionUser.Find(ctx, filter)
	if err != nil {
		return userdomain.Response{}, errors.New(err.Error() + "error in the finding user into the database")
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			errCh <- err
			return
		}
	}(cursor, ctx)

	var users []userdomain.User
	users = make([]userdomain.User, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var user userdomain.User
		if err = cursor.Decode(&user); err != nil {
			return userdomain.Response{}, err
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			users = append(users, user)
		}()
	}
	wg.Wait()

	response := userdomain.Response{
		User: users,
	}

	select {
	case err = <-errCh:
		return userdomain.Response{}, err
	default:
		return response, nil
	}
}

func (u userRepository) GetByEmail(ctx context.Context, email string) (*userdomain.User, error) {
	collectionUser := u.database.Collection(u.collectionUser)

	filter := bson.M{"email": email}
	var user *userdomain.User
	if err := collectionUser.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, errors.New(err.Error() + "error in the finding user into the database")
	}

	return user, nil
}

func (u userRepository) GetByID(ctx context.Context, id string) (*userdomain.User, error) {
	collectionUser := u.database.Collection(u.collectionUser)

	userID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": userID}
	var user *userdomain.User
	if err := collectionUser.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, errors.New(err.Error() + "error in the finding user into the database")
	}

	return user, nil
}

func (u userRepository) GetByVerificationCode(ctx context.Context, verificationCode string) (*userdomain.User, error) {
	collectionUser := u.database.Collection(u.collectionUser)

	filter := bson.M{"verification_code": verificationCode}
	var user *userdomain.User
	if err := collectionUser.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, errors.New(err.Error() + "error in the finding user's data into database")
	}

	return user, nil
}

func (u userRepository) CheckVerify(ctx context.Context, verificationCode string) bool {
	collectionUser := u.database.Collection(u.collectionUser)

	filter := bson.M{"verification_code": verificationCode}
	count, err := collectionUser.CountDocuments(ctx, filter)
	if err != nil {
		return false
	}

	if count > 0 {
		return true
	}

	return false
}

func (u userRepository) Login(ctx context.Context, request userdomain.SignIn) (*userdomain.User, error) {
	user, err := u.GetByEmail(ctx, request.Email)
	if err != nil {
		return nil, err
	}

	err = password.VerifyPassword(user.PasswordHash, request.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u userRepository) Create(ctx context.Context, user *userdomain.User) error {
	collectionUser := u.database.Collection(u.collectionUser)

	if err := user_validate.IsInvalidUser(user); err != nil {
		return err
	}

	filter := bson.M{"email": user.Email}
	count, err := collectionUser.CountDocuments(ctx, filter)
	if count > 0 {
		return err
	}

	_, err = collectionUser.InsertOne(ctx, user)
	if err != nil {
		return errors.New(err.Error() + "error in the inserting user into the database ")
	}

	return nil
}

func (u userRepository) Update(ctx context.Context, user *userdomain.UpdateUser) error {
	collectionUser := u.database.Collection(u.collectionUser)

	userData := userdomain.User{
		Username:  user.Username,
		AvatarURL: user.AvatarURL,
		UpdatedAt: time.Now(),
	}

	if err := user_validate.IsNilUsername(&userData); err != nil {
		return err
	}

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

	err := user_validate.IsNilPasswordHash(user)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": bson.M{"password": user.PasswordHash}}

	_, err = collectionUser.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New(err.Error() + "error in the updating user into database")
	}

	return nil

}

func (u userRepository) UpdateVerify(ctx context.Context, user *userdomain.User) (*mongo.UpdateResult, error) {
	collectionUser := u.database.Collection(u.collectionUser)

	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": bson.M{"verify": user.Verified}}

	updateResult, err := collectionUser.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, errors.New(err.Error() + "error in the updating user's data into database")
	}

	return updateResult, nil
}

func (u userRepository) UpdateVerifyForChangePassword(ctx context.Context, user *userdomain.User) (*mongo.UpdateResult, error) {
	collectionUser := u.database.Collection(u.collectionUser)

	filter := bson.D{{Key: "_id", Value: user.ID}}
	update := bson.D{{Key: "$set", Value: bson.M{
		"verify":     user.Verified,
		"updated_at": user.UpdatedAt,
	}}}

	data, err := collectionUser.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u userRepository) UpsertOne(ctx context.Context, email string, user *userdomain.User) (*userdomain.User, error) {
	collectionUser := u.database.Collection(u.collectionUser)

	err := user_validate.IsNilUsername(user)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"email": email}
	update := bson.M{"$set": bson.M{
		"username":   user.Username,
		"updated_at": time.Now(),
	}, "$setOnInsert": bson.M{
		"createdAt": time.Now(),
	}}

	opts := options.Update().SetUpsert(true)
	_, err = collectionUser.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return nil, errors.New(err.Error() + "error in the updating user's data into database")
	}

	return user, nil
}

func (u userRepository) UpdateImage(ctx context.Context, userID string, imageURL string) error {
	collectionUser := u.database.Collection(u.collectionUser)

	err := user_validate.IsNilImage(imageURL)
	if err != nil {
		return err
	}

	idUser, _ := primitive.ObjectIDFromHex(userID)
	filter := bson.M{"_id": idUser}
	update := bson.M{"$set": bson.M{"image_url": imageURL}}

	_, err = collectionUser.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New(err.Error() + "error in the updating user's data into database")
	}

	return nil
}

func (u userRepository) DeleteOne(ctx context.Context, userID string) error {
	collectionUser := u.database.Collection(u.collectionUser)

	idUser, _ := primitive.ObjectIDFromHex(userID)
	filter := bson.M{"_id": idUser}

	_, err := collectionUser.DeleteOne(ctx, filter)
	if err != nil {
		return errors.New(err.Error() + "error the deleting user's data into the database")
	}

	return nil
}

func (u userRepository) UniqueVerificationCode(ctx context.Context, verificationCode string) bool {
	collectionUser := u.database.Collection(u.collectionUser)

	filter := bson.M{"verification_code": verificationCode}
	count, err := collectionUser.CountDocuments(ctx, filter)
	if err != nil || count > 0 {
		return false
	}
	return true
}
