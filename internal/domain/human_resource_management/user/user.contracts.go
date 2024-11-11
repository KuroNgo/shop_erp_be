package user_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mime/multipart"
)

type IUserRepository interface {
	FetchMany(ctx context.Context) ([]User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	GetByID(ctx context.Context, id primitive.ObjectID) (User, error)
	GetByVerificationCode(ctx context.Context, verificationCode string) (User, error)

	UpdateOne(ctx context.Context, user *User) error
	UpdatePassword(ctx context.Context, user *User) error
	UpdateVerify(ctx context.Context, user *User) error
	UpdateVerificationCode(ctx context.Context, user *User) error
	UpsertOne(ctx context.Context, user *User) (*User, error)
	UpdateImage(ctx context.Context, user *User) error

	CreateOne(ctx context.Context, user *User) error
	DeleteOne(ctx context.Context, userID primitive.ObjectID) error
}

type IUserUseCase interface {
	FetchMany(ctx context.Context) ([]Output, error)
	GetByEmail(ctx context.Context, email string) (Output, error)
	GetByIDForCheckCookie(ctx context.Context, accessToken string) (*Output, error)
	GetByID(ctx context.Context, idUser string) (*Output, error)
	GetByVerificationCode(ctx context.Context, verificationCode string) (Output, error)

	UpdateOne(ctx context.Context, userID string, input *Input, file *multipart.FileHeader) error
	UpdateVerify(ctx context.Context, id string, input *Input) error
	UpdateImage(ctx context.Context, id string, input *Input) error

	SignUp(ctx context.Context, file *multipart.FileHeader, input *Input) error
	LoginUser(ctx context.Context, signIn *SignIn) (OutputLogin, error)
	LoginGoogle(ctx context.Context, code string) (*Output, *OutputLoginGoogle, error)
	DeleteOne(ctx context.Context, idUser string) error
	RefreshToken(ctx context.Context, refreshToken string) (*OutputLogin, error)

	ForgetPassword(ctx context.Context, email string) error
	UpdateVerifyForChangePassword(ctx context.Context, verificationCode string) error
	UpdatePassword(ctx context.Context, id string, input *ChangePasswordInput) error
}
