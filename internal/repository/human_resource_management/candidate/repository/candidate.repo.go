package candidate_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	candidatedomain "shop_erp_mono/internal/domain/human_resource_management/candidate"
	"shop_erp_mono/internal/repository"
)

type candidateRepository struct {
	database            *mongo.Database
	candidateCollection string
}

func NewCandidateRepository(database *mongo.Database, candidateCollection string) candidatedomain.ICandidateRepository {
	return &candidateRepository{database: database, candidateCollection: candidateCollection}
}

func (c *candidateRepository) CreateOne(ctx context.Context, candidate *candidatedomain.Candidate) error {
	candidateCollection := c.database.Collection(c.candidateCollection)

	_, err := candidateCollection.InsertOne(ctx, candidate)
	if err != nil {
		return err
	}

	return nil
}

func (c *candidateRepository) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	candidateCollection := c.database.Collection(c.candidateCollection)

	filter := bson.M{"_id": id}
	_, err := candidateCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (c *candidateRepository) UpdateOne(ctx context.Context, id primitive.ObjectID, candidate *candidatedomain.Candidate) error {
	candidateCollection := c.database.Collection(c.candidateCollection)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": candidate}
	_, err := candidateCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (c *candidateRepository) UpdateStatus(ctx context.Context, id primitive.ObjectID, status string) error {
	candidateCollection := c.database.Collection(c.candidateCollection)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": status}
	_, err := candidateCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (c *candidateRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*candidatedomain.Candidate, error) {
	candidateCollection := c.database.Collection(c.candidateCollection)

	filter := bson.M{"_id": id}
	var candidate candidatedomain.Candidate
	if err := candidateCollection.FindOne(ctx, filter).Decode(&candidate); err != nil {
		return nil, err
	}

	return &candidate, nil
}

func (c *candidateRepository) GetByEmail(ctx context.Context, email string) (*candidatedomain.Candidate, error) {
	candidateCollection := c.database.Collection(c.candidateCollection)

	filter := bson.M{"email": email}
	var candidate candidatedomain.Candidate
	if err := candidateCollection.FindOne(ctx, filter).Decode(&candidate); err != nil {
		return nil, err
	}

	return &candidate, nil
}

func (c *candidateRepository) GetAll(ctx context.Context) ([]candidatedomain.Candidate, error) {
	candidateCollection := c.database.Collection(c.candidateCollection)

	filter := bson.M{}
	cursor, err := candidateCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var candidates []candidatedomain.Candidate
	candidates = make([]candidatedomain.Candidate, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var candidate candidatedomain.Candidate
		if err = cursor.Decode(&candidate); err != nil {
			return nil, err
		}

		candidates = append(candidates, candidate)
	}

	return candidates, nil
}

func (c *candidateRepository) GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]candidatedomain.Candidate, error) {
	candidateCollection := c.database.Collection(c.candidateCollection)

	filter := bson.M{}
	cursor, err := repository.Paginate(ctx, candidateCollection, filter, pagination)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var candidates []candidatedomain.Candidate
	candidates = make([]candidatedomain.Candidate, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var candidate candidatedomain.Candidate
		if err = cursor.Decode(&candidate); err != nil {
			return nil, err
		}

		candidates = append(candidates, candidate)
	}

	return candidates, nil
}
