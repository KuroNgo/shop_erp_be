package performance_review_usecase

import (
	"context"
	performancereviewdomain "shop_erp_mono/domain/human_resource_management/performance_review"
	"time"
)

type performanceReviewUseCase struct {
	contextTimeout              time.Duration
	performanceReviewRepository performancereviewdomain.IPerformanceReviewRepository
}

func NewPerformanceReviewUseCase(contextTimeout time.Duration, performanceReviewRepository performancereviewdomain.IPerformanceReviewRepository) performancereviewdomain.IPerformanceReviewUseCase {
	return &performanceReviewUseCase{contextTimeout: contextTimeout, performanceReviewRepository: performanceReviewRepository}
}

func (p performanceReviewUseCase) CreateOne(ctx context.Context, input *performancereviewdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()
	return p.performanceReviewRepository.CreateOne(ctx, input)
}

func (p performanceReviewUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()
	return p.performanceReviewRepository.DeleteOne(ctx, id)
}

func (p performanceReviewUseCase) UpdateOne(ctx context.Context, input *performancereviewdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()
	return p.performanceReviewRepository.UpdateOne(ctx, input)
}

func (p performanceReviewUseCase) GetOneByID(ctx context.Context, id string) (performancereviewdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()
	return p.performanceReviewRepository.GetOneByID(ctx, id)
}

func (p performanceReviewUseCase) GetOneByEmailEmployee(ctx context.Context, name string) (performancereviewdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()
	return p.performanceReviewRepository.GetOneByEmailEmployee(ctx, name)
}

func (p performanceReviewUseCase) GetAll(ctx context.Context) ([]performancereviewdomain.Output, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()
	return p.performanceReviewRepository.GetAll(ctx)
}
