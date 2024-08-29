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

func (b benefitUseCase) CreateOne(ctx context.Context, input *benefitsdomain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (b benefitUseCase) DeleteOne(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (b benefitUseCase) UpdateOne(ctx context.Context, input *benefitsdomain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (b benefitUseCase) GetOneByID(ctx context.Context, id string) (benefitsdomain.Output, error) {
	//TODO implement me
	panic("implement me")
}

func (b benefitUseCase) GetOneByEmail(ctx context.Context, email string) (benefitsdomain.Output, error) {
	//TODO implement me
	panic("implement me")
}

func (b benefitUseCase) GetAll(ctx context.Context) ([]benefitsdomain.Output, error) {
	//TODO implement me
	panic("implement me")
}

func NewBenefitUseCase(contextTimeout time.Duration, benefitRepository benefitsdomain.IBenefitRepository) benefitsdomain.IBenefitUseCase {
	return &benefitUseCase{contextTimeout: contextTimeout, benefitRepository: benefitRepository}
}
