package user_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mime/multipart"
)

type IUserRepository interface {
	CreateOne(ctx context.Context, user *User) error
	UpdateOne(ctx context.Context, user *User) error
	UpdatePassword(ctx context.Context, user *User) error
	UpdateVerify(ctx context.Context, user *User) error
	UpsertOne(ctx context.Context, user *User) error
	UpdateImage(ctx context.Context, user *User) error
	DeleteOne(ctx context.Context, userID primitive.ObjectID) error
	FetchMany(ctx context.Context) ([]User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	GetByID(ctx context.Context, id primitive.ObjectID) (User, error)
	GetByVerificationCode(ctx context.Context, verificationCode string) (User, error)
}

type IUserUseCase interface {
	SignUp(ctx context.Context, input *Input) error
	LoginUser(ctx context.Context, signIn *SignIn) (OutputLogin, error)
	UpdateOne(ctx context.Context, userID string, input *Input, file *multipart.FileHeader) error
	UpdatePassword(ctx context.Context, input *Input) error
	UpdateVerify(ctx context.Context, input *Input) error
	UpdateVerifyForChangePassword(ctx context.Context, input *Input) error
	UpsertOne(ctx context.Context, email string, input *Input) error
	UpdateImage(ctx context.Context, input *Input) error
	DeleteOne(ctx context.Context, idUser string) error
	FetchMany(ctx context.Context) ([]Output, error)
	GetByEmail(ctx context.Context, email string) (Output, error)
	GetByIDForCheckCookie(ctx context.Context, idUser string) (*Output, error)
	GetByID(ctx context.Context, idUser string) (*Output, error)
	GetByVerificationCode(ctx context.Context, verificationCode string) (Output, error)
}
