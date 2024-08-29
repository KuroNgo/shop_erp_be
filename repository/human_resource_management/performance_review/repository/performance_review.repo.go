package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	performance_review_domain "shop_erp_mono/domain/human_resource_management/performance_review"
)

type performanceReviewRepository struct {
	database                    *mongo.Database
	collectionPerformanceReview string
	collectionEmployee          string
}

func (p performanceReviewRepository) CreateOne(ctx context.Context, input *performance_review_domain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (p performanceReviewRepository) DeleteOne(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (p performanceReviewRepository) UpdateOne(ctx context.Context, input *performance_review_domain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (p performanceReviewRepository) GetOneByID(ctx context.Context, id string) (performance_review_domain.Output, error) {
	//TODO implement me
	panic("implement me")
}

func (p performanceReviewRepository) GetOneByName(ctx context.Context, name string) (performance_review_domain.Output, error) {
	//TODO implement me
	panic("implement me")
}

func (p performanceReviewRepository) GetAllByEmployeeID(ctx context.Context, employeeID string) ([]performance_review_domain.Output, error) {
	//TODO implement me
	panic("implement me")
}

func NewPerformanceReviewRepository(db *mongo.Database, collectionPerformanceReview string, collectionEmployee string) performance_review_domain.IPerformanceReviewRepository {
	return &performanceReviewRepository{database: db, collectionPerformanceReview: collectionPerformanceReview, collectionEmployee: collectionEmployee}
}
