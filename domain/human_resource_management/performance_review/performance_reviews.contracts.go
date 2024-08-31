package performance_review_domain

import "context"

type IPerformanceReviewRepository interface {
	CreateOne(ctx context.Context, input *Input) error
	DeleteOne(ctx context.Context, id string) error
	UpdateOne(ctx context.Context, input *Input) error
	GetOneByID(ctx context.Context, id string) (Output, error)
	GetOneByEmailEmployee(ctx context.Context, name string) (Output, error)
	GetAll(ctx context.Context) ([]Output, error)
}

type IPerformanceReviewUseCase interface {
	CreateOne(ctx context.Context, input *Input) error
	DeleteOne(ctx context.Context, id string) error
	UpdateOne(ctx context.Context, input *Input) error
	GetOneByID(ctx context.Context, id string) (Output, error)
	GetOneByEmailEmployee(ctx context.Context, name string) (Output, error)
	GetAll(ctx context.Context) ([]Output, error)
}
