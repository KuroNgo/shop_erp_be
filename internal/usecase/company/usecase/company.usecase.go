package company_usecase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	companydomain "shop_erp_mono/internal/domain/company"
	"shop_erp_mono/internal/repository"
	"time"
)

type companyUseCase struct {
	contextTimeout    time.Duration
	companyRepository companydomain.ICompanyRepository
}

func NewCompanyUseCase(contextTimeout time.Duration, companyRepository companydomain.ICompanyRepository) companydomain.ICompanyUseCase {
	return &companyUseCase{contextTimeout: contextTimeout, companyRepository: companyRepository}
}

func (c *companyUseCase) CreateOne(ctx context.Context, company *companydomain.Company) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()
	return nil
}

func (c *companyUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()
	return nil
}

func (c *companyUseCase) DeleteSoft(ctx context.Context, id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()
	return nil
}

func (c *companyUseCase) UpdateOne(ctx context.Context, id string, company *companydomain.Company) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()
	return nil
}

func (c *companyUseCase) UpdateStatus(ctx context.Context, id string, status string) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()
	return nil
}

func (c *companyUseCase) GetByID(ctx context.Context, id string) (companydomain.CompanyResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()
	panic("implement me")
}

func (c *companyUseCase) GetByName(ctx context.Context, name string) (companydomain.CompanyResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()
	panic("implement me")
}

func (c *companyUseCase) GetAll(ctx context.Context) ([]companydomain.CompanyResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()
	panic("implement me")
}

func (c *companyUseCase) GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]companydomain.CompanyResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()
	panic("implement me")
}
