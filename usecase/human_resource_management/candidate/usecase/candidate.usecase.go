package candidate_usecase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	candidatedomain "shop_erp_mono/domain/human_resource_management/candidate"
	employeesdomain "shop_erp_mono/domain/human_resource_management/employees"
	"shop_erp_mono/repository"
	"time"
)

type candidateUseCase struct {
	contextTimeout      time.Duration
	candidateRepository candidatedomain.ICandidateRepository
	employeeRepository  employeesdomain.IEmployeeRepository
}

func NewCandidateUseCase(contextTimeout time.Duration, candidateRepository candidatedomain.ICandidateRepository,
	employeeRepository employeesdomain.IEmployeeRepository) candidatedomain.ICandidateUseCase {
	return &candidateUseCase{contextTimeout: contextTimeout, candidateRepository: candidateRepository, employeeRepository: employeeRepository}
}

func (c *candidateUseCase) CreateOne(ctx context.Context, candidate *candidatedomain.Candidate) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	return c.candidateRepository.CreateOne(ctx, candidate)
}

func (c *candidateUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	candidateID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return c.candidateRepository.DeleteOne(ctx, candidateID)
}

func (c *candidateUseCase) UpdateOne(ctx context.Context, id string, candidate *candidatedomain.Candidate) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	candidateID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return c.candidateRepository.UpdateOne(ctx, candidateID, candidate)
}

func (c *candidateUseCase) UpdateStatus(ctx context.Context, id string, status string) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	candidateID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return c.candidateRepository.UpdateStatus(ctx, candidateID, status)
}

func (c *candidateUseCase) GetByID(ctx context.Context, id string) (*candidatedomain.Candidate, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	candidateID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return c.candidateRepository.GetByID(ctx, candidateID)
}

func (c *candidateUseCase) GetByEmail(ctx context.Context, email string) (*candidatedomain.Candidate, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	return c.candidateRepository.GetByEmail(ctx, email)
}

func (c *candidateUseCase) GetAll(ctx context.Context) ([]candidatedomain.Candidate, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	return c.candidateRepository.GetAll(ctx)
}

func (c *candidateUseCase) GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]candidatedomain.Candidate, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	return c.candidateRepository.GetAllWithPagination(ctx, pagination)
}
