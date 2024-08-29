package benefit_usecase

import (
	"context"
	benefitsdomain "shop_erp_mono/domain/human_resource_management/benefits"
	"time"
)

type benefitUseCase struct {
	contextTimeout    time.Duration
	benefitRepository benefitsdomain.IBenefitRepository
}

func NewBenefitUseCase(contextTimeout time.Duration, benefitRepository benefitsdomain.IBenefitRepository) benefitsdomain.IBenefitUseCase {
	return &benefitUseCase{contextTimeout: contextTimeout, benefitRepository: benefitRepository}
}

func (b benefitUseCase) CreateOne(ctx context.Context, input *benefitsdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()

	err := b.benefitRepository.CreateOne(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (b benefitUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()

	err := b.benefitRepository.DeleteOne(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (b benefitUseCase) UpdateOne(ctx context.Context, id string, input *benefitsdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()

	err := b.benefitRepository.UpdateOne(ctx, id, input)
	if err != nil {
		return err
	}

	return nil
}

func (b benefitUseCase) GetOneByID(ctx context.Context, id string) (benefitsdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()

	data, err := b.benefitRepository.GetOneByID(ctx, id)
	if err != nil {
		return benefitsdomain.Output{}, err
	}

	return data, nil
}

func (b benefitUseCase) GetOneByEmail(ctx context.Context, email string) (benefitsdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()

	data, err := b.benefitRepository.GetOneByEmail(ctx, email)
	if err != nil {
		return benefitsdomain.Output{}, err
	}

	return data, nil
}

func (b benefitUseCase) GetAll(ctx context.Context) ([]benefitsdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()

	data, err := b.benefitRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}
