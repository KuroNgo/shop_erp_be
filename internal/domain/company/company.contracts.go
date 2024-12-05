package company_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"shop_erp_mono/internal/repository"
)

type ICompanyRepository interface {
	CreateOne(ctx context.Context, company *Company) error
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
	DeleteSoft(ctx context.Context, id primitive.ObjectID) error
	UpdateOne(ctx context.Context, id primitive.ObjectID, company *Company) error
	UpdateStatus(ctx context.Context, id primitive.ObjectID, status string) error
	GetByID(ctx context.Context, id primitive.ObjectID) (Company, error)
	GetByName(ctx context.Context, name string) (Company, error)
	GetAll(ctx context.Context) ([]Company, error)
	GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]Company, error)
}

type ICompanyUseCase interface {
	CreateOne(ctx context.Context, company *Company) error
	DeleteOne(ctx context.Context, id string) error
	DeleteSoft(ctx context.Context, id primitive.ObjectID) error
	UpdateOne(ctx context.Context, id string, company *Company) error
	UpdateStatus(ctx context.Context, id string, status string) error
	GetByID(ctx context.Context, id string) (CompanyResponse, error)
	GetByName(ctx context.Context, name string) (CompanyResponse, error)
	GetAll(ctx context.Context) ([]CompanyResponse, error)
	GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]CompanyResponse, error)
}
