package user_usecase

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mime/multipart"
	"shop_erp_mono/bootstrap"
	userdomain "shop_erp_mono/domain/human_resource_management/user"
	"shop_erp_mono/pkg/token"
	"time"
)

type userUseCase struct {
	database       *bootstrap.Database
	contextTimeout time.Duration
	userRepository userdomain.IUserRepository
}

func NewUserUseCase(database *bootstrap.Database, contextTimeout time.Duration, userRepository userdomain.IUserRepository) userdomain.IUserUseCase {
	return &userUseCase{database: database, contextTimeout: contextTimeout, userRepository: userRepository}
}

func (u userUseCase) SignUp(ctx context.Context, input *userdomain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (u userUseCase) LoginUser(ctx context.Context, signIn *userdomain.Input) (userdomain.OutputLogin, error) {
	//TODO implement me
	panic("implement me")
}

func (u userUseCase) UpdateOne(ctx context.Context, userID string, input *userdomain.Input, file *multipart.FileHeader) error {
	//ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	//defer cancel()
	//
	//idUser, err := primitive.ObjectIDFromHex(userID)
	//if err != nil {
	//	return err
	//}
	//
	//userData, err := u.userRepository.GetByID(ctx, idUser)
	//if err != nil {
	//	return err
	//}
	//
	//TODO implement me
	panic("implement me")
}

func (u userUseCase) UpdatePassword(ctx context.Context, input *userdomain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (u userUseCase) UpdateVerify(ctx context.Context, input *userdomain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (u userUseCase) UpdateVerifyForChangePassword(ctx context.Context, input *userdomain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (u userUseCase) UpsertOne(ctx context.Context, email string, input *userdomain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (u userUseCase) UpdateImage(ctx context.Context, input *userdomain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (u userUseCase) DeleteOne(ctx context.Context, idUser string) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	userID, err := primitive.ObjectIDFromHex(idUser)
	if err != nil {
		return err
	}

	return u.userRepository.DeleteOne(ctx, userID)
}

func (u userUseCase) FetchMany(ctx context.Context) ([]userdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	userData, err := u.userRepository.FetchMany(ctx)
	if err != nil {
		return nil, err
	}

	var outputs []userdomain.Output
	outputs = make([]userdomain.Output, 0, len(userData))
	for _, user := range userData {
		output := userdomain.Output{
			User: user,
		}

		outputs = append(outputs, output)
	}

	return outputs, nil
}

func (u userUseCase) GetByEmail(ctx context.Context, email string) (userdomain.Output, error) {
	//TODO implement me
	panic("implement me")
}

func (u userUseCase) GetByID(ctx context.Context, id string) (userdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	sub, err := token.ValidateToken(id, u.database.AccessTokenPublicKey)
	if err != nil {
		return userdomain.Output{}, err
	}

	idUser, err := primitive.ObjectIDFromHex(fmt.Sprint(sub))
	if err != nil {
		return userdomain.Output{}, err
	}

	userData, err := u.userRepository.GetByID(ctx, idUser)
	if err != nil {
		return userdomain.Output{}, err
	}

	output := userdomain.Output{
		User: userData,
	}

	return output, nil
}

func (u userUseCase) GetByVerificationCode(ctx context.Context, verificationCode string) (userdomain.Output, error) {
	//TODO implement me
	panic("implement me")
}
