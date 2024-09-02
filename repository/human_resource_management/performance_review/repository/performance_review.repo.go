package performance_review_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	performancereviewdomain "shop_erp_mono/domain/human_resource_management/performance_review"
)

type performanceReviewRepository struct {
	database                    *mongo.Database
	collectionPerformanceReview string
}

func NewPerformanceReviewRepository(db *mongo.Database, collectionPerformanceReview string) performancereviewdomain.IPerformanceReviewRepository {
	return &performanceReviewRepository{database: db, collectionPerformanceReview: collectionPerformanceReview}
}

func (p performanceReviewRepository) CreateOne(ctx context.Context, input *performancereviewdomain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (p performanceReviewRepository) DeleteOne(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (p performanceReviewRepository) UpdateOne(ctx context.Context, input *performancereviewdomain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (p performanceReviewRepository) GetOneByID(ctx context.Context, id string) (performancereviewdomain.Output, error) {
	//TODO implement me
	panic("implement me")
}

func (p performanceReviewRepository) GetOneByEmailEmployee(ctx context.Context, name string) (performancereviewdomain.Output, error) {
	//TODO implement me
	panic("implement me")
}

func (p performanceReviewRepository) GetAll(ctx context.Context) ([]performancereviewdomain.Output, error) {
	//TODO implement me
	panic("implement me")
}
