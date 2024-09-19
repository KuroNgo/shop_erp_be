package contracts_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IContractsRepository interface {
	CreateOne(ctx context.Context, contract *Contract) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	UpdateOne(ctx context.Context, id primitive.ObjectID, contract *Contract) error
	GetByID(ctx context.Context, id primitive.ObjectID) (Contract, error)
	GetByEmployeeID(ctx context.Context, employeeID primitive.ObjectID) (Contract, error)
	GetAll(ctx context.Context) ([]Contract, error)
}

type IContractsUseCase interface {
	CreateOne(ctx context.Context, input *Input) error
	DeleteOne(ctx context.Context, id string) error
	UpdateOne(ctx context.Context, id string, input *Input) error
	GetByID(ctx context.Context, id string) (Output, error)
	GetByEmail(ctx context.Context, email string) (Output, error)
	GetAll(ctx context.Context) ([]Output, error)
}
