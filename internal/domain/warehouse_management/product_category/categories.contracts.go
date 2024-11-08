package category_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ICategoryRepository interface {
	CreateOne(ctx context.Context, category Category) error
	GetByID(ctx context.Context, id primitive.ObjectID) (*Category, error)
	GetByName(ctx context.Context, name string) (*Category, error)
	UpdateOne(ctx context.Context, id primitive.ObjectID, category Category) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	GetAll(ctx context.Context) ([]Category, error)
}

type ICategoryUseCase interface {
	CreateOne(ctx context.Context, input *Input) error
	GetByID(ctx context.Context, id string) (*CategoryResponse, error)
	GetByName(ctx context.Context, name string) (*CategoryResponse, error)
	GetAll(ctx context.Context) ([]CategoryResponse, error)
	UpdateOne(ctx context.Context, id string, input *Input) error
	DeleteOne(ctx context.Context, id string) error
}
