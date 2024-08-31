package contract_usecase

import (
	"context"
	contractsdomain "shop_erp_mono/domain/human_resource_management/contracts"
	"time"
)

type contractUseCase struct {
	contextTimeout     time.Duration
	contractRepository contractsdomain.IContractsRepository
}

func NewContractUseCase(contextTimeout time.Duration, contractRepository contractsdomain.IContractsRepository) contractsdomain.IContractsUseCase {
	return &contractUseCase{contextTimeout: contextTimeout, contractRepository: contractRepository}
}

func (c contractUseCase) CreateOne(ctx context.Context, input *contractsdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	err := c.contractRepository.CreateOne(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (c contractUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	err := c.contractRepository.DeleteOne(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (c contractUseCase) UpdateOne(ctx context.Context, id string, input *contractsdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	err := c.contractRepository.UpdateOne(ctx, id, input)
	if err != nil {
		return err
	}

	return nil
}

func (c contractUseCase) GetOneByID(ctx context.Context, id string) (contractsdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	data, err := c.contractRepository.GetOneByID(ctx, id)
	if err != nil {
		return contractsdomain.Output{}, err
	}

	return data, nil
}

func (c contractUseCase) GetOneByEmail(ctx context.Context, email string) (contractsdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	data, err := c.contractRepository.GetOneByEmail(ctx, email)
	if err != nil {
		return contractsdomain.Output{}, err
	}

	return data, nil
}

func (c contractUseCase) GetAll(ctx context.Context) ([]contractsdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	data, err := c.contractRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}
