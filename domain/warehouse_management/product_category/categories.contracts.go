package category_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ICategoryRepository interface {
	Create(ctx context.Context, category Category) error
	GetByID(ctx context.Context, id primitive.ObjectID) (*Category, error)
	GetByName(ctx context.Context, name string) (*Category, error)
	Update(ctx context.Context, id primitive.ObjectID, input Category) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	GetAll(ctx context.Context) ([]Category, error)
}

type ICategoryUseCase interface {
	CreateCategory(ctx context.Context, input Input) error
	GetByIDCategory(ctx context.Context, id string) (*CategoryResponse, error)
	UpdateCategory(ctx context.Context, id string, input Input) error
	DeleteCategory(ctx context.Context, id string) error
	GetAllCategories(ctx context.Context) ([]CategoryResponse, error)
}
