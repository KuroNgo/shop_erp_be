package performance_review_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IPerformanceReviewRepository interface {
	CreateOne(ctx context.Context, performanceReview *PerformanceReview) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	UpdateOne(ctx context.Context, id primitive.ObjectID, performanceReview *PerformanceReview) error
	GetOneByID(ctx context.Context, id primitive.ObjectID) (PerformanceReview, error)
	GetOneByEmployeeID(ctx context.Context, employeeID primitive.ObjectID) (PerformanceReview, error)
	GetAll(ctx context.Context) ([]PerformanceReview, error)
}

type IPerformanceReviewUseCase interface {
	CreateOne(ctx context.Context, input *Input) error
	DeleteOne(ctx context.Context, id string) error
	UpdateOne(ctx context.Context, id string, input *Input) error
	GetOneByID(ctx context.Context, id string) (Output, error)
	GetOneByEmailEmployee(ctx context.Context, name string) (Output, error)
	GetAll(ctx context.Context) ([]Output, error)
}
