package user

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	userdomain "shop_erp_mono/domain/human_resource_management/user"
	"time"
)

type userUseCase struct {
	contextTimeout time.Duration
	userRepository userdomain.IUserRepository
}

func NewUserUseCase(contextTimeout time.Duration, userRepository userdomain.IUserRepository) userdomain.IUserUseCase {
	return &userUseCase{contextTimeout: contextTimeout, userRepository: userRepository}
}

func (u userUseCase) FetchMany(ctx context.Context) (userdomain.Response, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	data, err := u.userRepository.FetchMany(ctx)
	if err != nil {
		return userdomain.Response{}, err
	}

	return data, nil
}

func (u userUseCase) GetByEmail(ctx context.Context, email string) (*userdomain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	data, err := u.userRepository.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u userUseCase) GetByID(ctx context.Context, id primitive.ObjectID) (*userdomain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	data, err := u.userRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u userUseCase) CheckVerify(ctx context.Context, verificationCode string) bool {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	res := u.userRepository.CheckVerify(ctx, verificationCode)
	return res
}

func (u userUseCase) GetByVerificationCode(ctx context.Context, verificationCode string) (*userdomain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	data, err := u.userRepository.GetByVerificationCode(ctx, verificationCode)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u userUseCase) Create(ctx context.Context, user *userdomain.User) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	err := u.userRepository.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (u userUseCase) UpsertOne(ctx context.Context, email string, user *userdomain.User) (*userdomain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	data, err := u.userRepository.UpsertOne(ctx, email, user)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u userUseCase) Login(ctx context.Context, request userdomain.SignIn) (*userdomain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	data, err := u.userRepository.Login(ctx, request)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u userUseCase) Update(ctx context.Context, user *userdomain.UpdateUser) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	err := u.userRepository.Update(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (u userUseCase) UpdatePassword(ctx context.Context, user *userdomain.User) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	err := u.userRepository.UpdatePassword(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (u userUseCase) UpdateVerify(ctx context.Context, user *userdomain.User) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	_, err := u.userRepository.UpdateVerify(ctx, user)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (u userUseCase) UpdateVerifyForChangePassword(ctx context.Context, user *userdomain.User) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	_, err := u.userRepository.UpdateVerify(ctx, user)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (u userUseCase) UpdateImage(ctx context.Context, userID string, imageURL string) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	err := u.userRepository.UpdateImage(ctx, userID, imageURL)
	if err != nil {
		return err
	}

	return nil
}

func (u userUseCase) DeleteOne(ctx context.Context, userID string) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	err := u.userRepository.DeleteOne(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}

func (u userUseCase) UniqueVerificationCode(ctx context.Context, verificationCode string) bool {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	res := u.userRepository.UniqueVerificationCode(ctx, verificationCode)
	return res
}
