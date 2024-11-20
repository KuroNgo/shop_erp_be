package performance_review_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IPerformanceReviewRepository interface {
	CreateOne(ctx context.Context, performanceReview *PerformanceReview) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	UpdateOne(ctx context.Context, id primitive.ObjectID, performanceReview *PerformanceReview) error
	GetByID(ctx context.Context, id primitive.ObjectID) (PerformanceReview, error)
	GetByEmployeeID(ctx context.Context, employeeID primitive.ObjectID) (PerformanceReview, error)
	GetAll(ctx context.Context) ([]PerformanceReview, error)
}

type IPerformanceReviewUseCase interface {
	CreateOneWithEmailEmployee(ctx context.Context, input *Input1) error
	CreateOneWithIDEmployee(ctx context.Context, input *Input2) error

	DeleteOne(ctx context.Context, id string) error

	UpdateOneWithEmailEmployee(ctx context.Context, id string, input *Input1) error
	UpdateOneWithIDEmployee(ctx context.Context, id string, input *Input2) error

	GetByID(ctx context.Context, id string) (Output, error)
	GetByEmailEmployee(ctx context.Context, name string) (Output, error)
	GetAll(ctx context.Context) ([]Output, error)
}
