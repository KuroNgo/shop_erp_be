package contract_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	contractsdomain "shop_erp_mono/domain/human_resource_management/contracts"
)

type contractRepository struct {
	database           *mongo.Database
	collectionContract string
	collectionEmployee string
}

func NewContractRepository(db *mongo.Database, collectionContract string, collectionEmployee string) contractsdomain.IContractsRepository {
	return &contractRepository{database: db, collectionContract: collectionContract, collectionEmployee: collectionEmployee}
}

func (c contractRepository) CreateOne(ctx context.Context, input *contractsdomain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (c contractRepository) DeleteOne(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (c contractRepository) UpdateOne(ctx context.Context, input *contractsdomain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (c contractRepository) GetOneByID(ctx context.Context, id string) (contractsdomain.Output, error) {
	//TODO implement me
	panic("implement me")
}

func (c contractRepository) GetOneByEmail(ctx context.Context, email string) (contractsdomain.Output, error) {
	//TODO implement me
	panic("implement me")
}

func (c contractRepository) GetAllByEmployeeID(ctx context.Context, employeeID string) ([]contractsdomain.Output, error) {
	//TODO implement me
	panic("implement me")
}
