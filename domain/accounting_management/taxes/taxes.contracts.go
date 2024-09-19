package taxes_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ITaxesRepository interface {
	CreateOne(ctx context.Context, tax *Taxes) error
	GetByID(ctx context.Context, id primitive.ObjectID) (Taxes, error)
	GetByName(ctx context.Context, name string) (Taxes, error)
	UpdateOne(ctx context.Context, tax *Taxes) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	GetAll(ctx context.Context) ([]Taxes, error)
}

type ITaxesUseCase interface {
	CreateOne(ctx context.Context, input *Input) error
	GetByID(ctx context.Context, id string) (TaxesResponse, error)
	GetByName(ctx context.Context, name string) (TaxesResponse, error)
	UpdateOne(ctx context.Context, id string, input *Input) error
	DeleteOne(ctx context.Context, id string) error
	GetAll(ctx context.Context) ([]TaxesResponse, error)
}
