package taxes_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ITaxesRepository interface {
	Create(ctx context.Context, tax *Taxes) error
	GetByID(ctx context.Context, id primitive.ObjectID) (Taxes, error)
	GetByName(ctx context.Context, name string) (Taxes, error)
	Update(ctx context.Context, tax *Taxes) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	List(ctx context.Context) ([]Taxes, error)
}

type ITaxesUseCase interface {
	CreateTax(ctx context.Context, input *Input) error
	GetTaxByID(ctx context.Context, id string) (TaxesResponse, error)
	GetTaxByName(ctx context.Context, name string) (TaxesResponse, error)
	UpdateTax(ctx context.Context, id string, input *Input) error
	DeleteTax(ctx context.Context, id string) error
	ListTaxes(ctx context.Context) ([]TaxesResponse, error)
}
