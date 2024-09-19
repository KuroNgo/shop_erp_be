package benefits_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IBenefitRepository interface {
	CreateOne(ctx context.Context, benefit *Benefit) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	UpdateOne(ctx context.Context, id primitive.ObjectID, benefit *Benefit) error
	GetByID(ctx context.Context, id primitive.ObjectID) (Benefit, error)
	GetByEmployeeID(ctx context.Context, employeeID primitive.ObjectID) (Benefit, error)
	GetAll(ctx context.Context) ([]Benefit, error)
}

type IBenefitUseCase interface {
	CreateOne(ctx context.Context, input *Input) error
	DeleteOne(ctx context.Context, id string) error
	UpdateOne(ctx context.Context, id string, input *Input) error
	GetByID(ctx context.Context, id string) (Output, error)
	GetByEmail(ctx context.Context, email string) (Output, error)
	GetAll(ctx context.Context) ([]Output, error)
}
