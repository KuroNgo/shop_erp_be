package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Pagination struct {
	Limit int64 `json:"limit"`
	Page  int64 `json:"page"`
}

func Paginate(ctx context.Context, collection *mongo.Collection, filter interface{}, pagination Pagination) (*mongo.Cursor, error) {
	// Calculate skip and limit
	skip := (pagination.Page - 1) * pagination.Limit

	findOptions := options.Find()
	findOptions.SetLimit(pagination.Limit)
	findOptions.SetSkip(skip)

	// Acting query with pagination
	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}

	return cursor, nil
}
