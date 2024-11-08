package candidate_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"shop_erp_mono/internal/repository"
)

type ICandidateRepository interface {
	CreateOne(ctx context.Context, candidate *Candidate) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	UpdateOne(ctx context.Context, id primitive.ObjectID, candidate *Candidate) error
	UpdateStatus(ctx context.Context, id primitive.ObjectID, status string) error
	GetByID(ctx context.Context, id primitive.ObjectID) (*Candidate, error)
	GetByEmail(ctx context.Context, email string) (*Candidate, error)
	GetAll(ctx context.Context) ([]Candidate, error)
	GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]Candidate, error)
}

type ICandidateUseCase interface {
	CreateOne(ctx context.Context, candidate *Candidate) error
	DeleteOne(ctx context.Context, id string) error
	UpdateOne(ctx context.Context, id string, candidate *Candidate) error
	UpdateStatus(ctx context.Context, id string, status string) error
	GetByID(ctx context.Context, id string) (*Candidate, error)
	GetByEmail(ctx context.Context, email string) (*Candidate, error)
	GetAll(ctx context.Context) ([]Candidate, error)
	GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]Candidate, error)
}
