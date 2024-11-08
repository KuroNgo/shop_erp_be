package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
)

type Pagination struct {
	Limit string `json:"limit"`
	Page  string `json:"page"`
}

func Paginate(ctx context.Context, collection *mongo.Collection, filter interface{}, pagination Pagination) (*mongo.Cursor, error) {
	page, err := strconv.ParseInt(pagination.Page, 10, 64)
	limit, err := strconv.ParseInt(pagination.Limit, 10, 64)

	// Calculate skip and limit
	skip := (page - 1) * limit

	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip(skip)

	// Acting query with pagination
	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}

	return cursor, nil
}
