package performance_review_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	performancereviewdomain "shop_erp_mono/domain/human_resource_management/performance_review"
	"time"
)

type performanceReviewRepository struct {
	database                    *mongo.Database
	collectionPerformanceReview string
}

func NewPerformanceReviewRepository(db *mongo.Database, collectionPerformanceReview string) performancereviewdomain.IPerformanceReviewRepository {
	return &performanceReviewRepository{database: db, collectionPerformanceReview: collectionPerformanceReview}
}

func (p performanceReviewRepository) CreateOne(ctx context.Context, performanceReview *performancereviewdomain.PerformanceReview) error {
	collectionPerformanceReview := p.database.Collection(p.collectionPerformanceReview)

	_, err := collectionPerformanceReview.InsertOne(ctx, performanceReview)
	if err != nil {
		return err
	}

	return nil
}

func (p performanceReviewRepository) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	collectionPerformanceReview := p.database.Collection(p.collectionPerformanceReview)

	filter := bson.M{"_id": id}
	_, err := collectionPerformanceReview.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (p performanceReviewRepository) UpdateOne(ctx context.Context, id primitive.ObjectID,
	performanceReview *performancereviewdomain.PerformanceReview) error {
	collectionPerformanceReview := p.database.Collection(p.collectionPerformanceReview)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"employee_id":       performanceReview.EmployeeID,
		"review_date":       performanceReview.ReviewDate,
		"reviewer_id":       performanceReview.ReviewerID,
		"performance_score": performanceReview.PerformanceScore,
		"comments":          performanceReview.Comments,
		"updated_at":        time.Now(),
	}}
	_, err := collectionPerformanceReview.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (p performanceReviewRepository) GetOneByID(ctx context.Context, id primitive.ObjectID) (performancereviewdomain.PerformanceReview, error) {
	collectionPerformanceReview := p.database.Collection(p.collectionPerformanceReview)

	var performanceReview performancereviewdomain.PerformanceReview
	filter := bson.M{"_id": id}
	if err := collectionPerformanceReview.FindOne(ctx, filter).Decode(&performanceReview); err != nil {
		return performancereviewdomain.PerformanceReview{}, err
	}

	return performanceReview, nil
}

func (p performanceReviewRepository) GetOneByEmployeeID(ctx context.Context, employeeID primitive.ObjectID) (performancereviewdomain.PerformanceReview, error) {
	collectionPerformanceReview := p.database.Collection(p.collectionPerformanceReview)

	var performanceReview performancereviewdomain.PerformanceReview
	filter := bson.M{"employee_id": employeeID}
	if err := collectionPerformanceReview.FindOne(ctx, filter).Decode(&performanceReview); err != nil {
		return performancereviewdomain.PerformanceReview{}, err
	}

	return performanceReview, nil
}

func (p performanceReviewRepository) GetAll(ctx context.Context) ([]performancereviewdomain.PerformanceReview, error) {
	collectionPerformanceReview := p.database.Collection(p.collectionPerformanceReview)

	filter := bson.M{}
	cursor, err := collectionPerformanceReview.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	var performanceReviews []performancereviewdomain.PerformanceReview
	performanceReviews = make([]performancereviewdomain.PerformanceReview, 0, cursor.RemainingBatchLength())

	for cursor.Next(ctx) {
		var performanceReview performancereviewdomain.PerformanceReview
		if err = cursor.Decode(&performanceReviews); err != nil {
			return nil, err
		}

		performanceReviews = append(performanceReviews, performanceReview)
	}

	return performanceReviews, nil
}
