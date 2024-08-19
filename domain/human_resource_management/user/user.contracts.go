package user_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

//go:generate mockery --name IUserRepository
type IUserRepository interface {
	FetchMany(ctx context.Context) (Response, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByID(ctx context.Context, id string) (*User, error)
	GetByVerificationCode(ctx context.Context, verificationCode string) (*User, error)
	CheckVerify(ctx context.Context, verificationCode string) bool

	Login(ctx context.Context, request SignIn) (*User, error)
	Create(ctx context.Context, user *SignUp) error
	Update(ctx context.Context, user *User) error
	UpdatePassword(ctx context.Context, user *User) error
	UpdateVerify(ctx context.Context, user *User) (*mongo.UpdateResult, error)
	UpdateVerifyForChangePassword(ctx context.Context, user *User) (*mongo.UpdateResult, error)
	UpsertOne(ctx context.Context, email string, user *User) (*User, error)
	UpdateImage(ctx context.Context, userID string, imageURL string) error
	DeleteOne(ctx context.Context, userID string) error
	UniqueVerificationCode(ctx context.Context, verificationCode string) bool
}

//go:generate mockery --name IUserUseCase
type IUserUseCase interface {
	FetchMany(ctx context.Context) (Response, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByID(ctx context.Context, id string) (*User, error)
	CheckVerify(ctx context.Context, verificationCode string) bool
	GetByVerificationCode(ctx context.Context, verificationCode string) (*User, error)

	Create(ctx context.Context, user *SignUp) error
	UpsertUser(ctx context.Context, email string, user *User) (*User, error)
	Login(ctx context.Context, request SignIn) (*User, error)
	Update(ctx context.Context, user *User) error
	UpdatePassword(ctx context.Context, user *User) error
	UpdateVerify(ctx context.Context, user *User) (*mongo.UpdateResult, error)
	UpdateVerifyForChangePassword(ctx context.Context, user *User) (*mongo.UpdateResult, error)
	UpdateImage(ctx context.Context, userID string, imageURL string) error
	Delete(ctx context.Context, userID string) error
	UniqueVerificationCode(ctx context.Context, verificationCode string) bool
}
