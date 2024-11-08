package taxes_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	taxes_domain "shop_erp_mono/domain/accounting_management/taxes"
)

type taxesRepository struct {
	database        *mongo.Database
	taxesCollection string
}

func NewTaxesRepository(database *mongo.Database, taxesCollection string) taxes_domain.ITaxesRepository {
	return &taxesRepository{database: database, taxesCollection: taxesCollection}
}

func (t *taxesRepository) CreateOne(ctx context.Context, tax *taxes_domain.Taxes) error {
	//TODO implement me
	panic("implement me")
}

func (t *taxesRepository) GetByID(ctx context.Context, id primitive.ObjectID) (taxes_domain.Taxes, error) {
	//TODO implement me
	panic("implement me")
}

func (t *taxesRepository) GetByName(ctx context.Context, name string) (taxes_domain.Taxes, error) {
	//TODO implement me
	panic("implement me")
}

func (t *taxesRepository) UpdateOne(ctx context.Context, tax *taxes_domain.Taxes) error {
	//TODO implement me
	panic("implement me")
}

func (t *taxesRepository) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	//TODO implement me
	panic("implement me")
}

func (t *taxesRepository) GetAll(ctx context.Context) ([]taxes_domain.Taxes, error) {
	//TODO implement me
	panic("implement me")
}
