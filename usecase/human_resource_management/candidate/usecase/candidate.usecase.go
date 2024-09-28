package candidate_usecase

import (
	"context"
	"encoding/json"
	"github.com/allegro/bigcache/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	candidatedomain "shop_erp_mono/domain/human_resource_management/candidate"
	employeesdomain "shop_erp_mono/domain/human_resource_management/employees"
	"shop_erp_mono/repository"
	"shop_erp_mono/usecase/human_resource_management/candidate/validate"
	"sync"
	"time"
)

type candidateUseCase struct {
	contextTimeout      time.Duration
	candidateRepository candidatedomain.ICandidateRepository
	employeeRepository  employeesdomain.IEmployeeRepository
	cache               *bigcache.BigCache
}

func NewCandidateUseCase(contextTimeout time.Duration, candidateRepository candidatedomain.ICandidateRepository,
	employeeRepository employeesdomain.IEmployeeRepository, cacheTTL time.Duration) candidatedomain.ICandidateUseCase {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(cacheTTL))
	if err != nil {
		return nil
	}
	return &candidateUseCase{contextTimeout: contextTimeout, cache: cache, candidateRepository: candidateRepository, employeeRepository: employeeRepository}
}

func (c *candidateUseCase) CreateOne(ctx context.Context, candidate *candidatedomain.Candidate) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	if err := validate.Candidate(candidate); err != nil {
		return err
	}

	errCh := make(chan error, 1)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := c.cache.Delete("candidates")
		if err != nil {
			errCh <- err
			return
		}
	}()

	go func() {
		wg.Wait()
		close(errCh)
	}()

	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
	case <-ctx.Done():
		return ctx.Err()
	}

	return c.candidateRepository.CreateOne(ctx, candidate)
}

func (c *candidateUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	candidateID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	errCh := make(chan error, 1)
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		err = c.cache.Delete("candidates")
		if err != nil {
			errCh <- err
			return
		}
	}()

	go func() {
		defer wg.Done()
		err = c.cache.Delete(id)
		if err != nil {
			errCh <- err
			return
		}
	}()

	go func() {
		wg.Wait()
		close(errCh)
	}()

	select {
	case err = <-errCh:
		if err != nil {
			return err
		}
	case <-ctx.Done():
		return ctx.Err()
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

	errCh := make(chan error, 1)
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		err = c.cache.Delete("candidates")
		if err != nil {
			errCh <- err
			return
		}
	}()

	go func() {
		defer wg.Done()
		err = c.cache.Delete(id)
		if err != nil {
			errCh <- err
			return
		}
	}()

	go func() {
		wg.Wait()
		close(errCh)
	}()

	select {
	case err = <-errCh:
		if err != nil {
			return err
		}
	case <-ctx.Done():
		return ctx.Err()
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

	errCh := make(chan error, 1)
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		err = c.cache.Delete("candidates")
		if err != nil {
			errCh <- err
			return
		}
	}()

	go func() {
		defer wg.Done()
		err = c.cache.Delete(id)
		if err != nil {
			errCh <- err
			return
		}
	}()

	go func() {
		wg.Wait()
		close(errCh)
	}()

	select {
	case err = <-errCh:
		if err != nil {
			return err
		}
	case <-ctx.Done():
		return ctx.Err()
	}
	return c.candidateRepository.UpdateStatus(ctx, candidateID, status)
}

func (c *candidateUseCase) GetByID(ctx context.Context, id string) (*candidatedomain.Candidate, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	data, _ := c.cache.Get(id)
	if data != nil {
		var response *candidatedomain.Candidate
		err := json.Unmarshal(data, response)
		if err != nil {
			return nil, err
		}
		return response, nil
	}

	candidateID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	response, err := c.candidateRepository.GetByID(ctx, candidateID)
	if err != nil {
		return nil, err
	}

	data, err = json.Marshal(response)
	if err != nil {
		return nil, err
	}

	err = c.cache.Set(id, data)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *candidateUseCase) GetByEmail(ctx context.Context, email string) (*candidatedomain.Candidate, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	data, _ := c.cache.Get(email)
	if data != nil {
		var response *candidatedomain.Candidate
		err := json.Unmarshal(data, response)
		if err != nil {
			return nil, err
		}
		return response, nil
	}

	candidateData, err := c.candidateRepository.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	data, err = json.Marshal(candidateData)
	if err != nil {
		return nil, err
	}

	err = c.cache.Set(email, data)
	if err != nil {
		return nil, err
	}

	return candidateData, nil
}

func (c *candidateUseCase) GetAll(ctx context.Context) ([]candidatedomain.Candidate, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	data, _ := c.cache.Get("candidates")
	if data != nil {
		var response []candidatedomain.Candidate
		err := json.Unmarshal(data, &response)
		if err != nil {
			return nil, err
		}
		return response, nil
	}

	candidateData, err := c.candidateRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	data, err = json.Marshal(candidateData)
	if err != nil {
		return nil, err
	}

	err = c.cache.Set("candidates", data)
	if err != nil {
		return nil, err
	}

	return candidateData, nil
}

func (c *candidateUseCase) GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]candidatedomain.Candidate, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	data, _ := c.cache.Get(pagination.Page)
	if data != nil {
		var response []candidatedomain.Candidate
		err := json.Unmarshal(data, &response)
		if err != nil {
			return nil, err
		}
		return response, nil
	}

	candidateData, err := c.candidateRepository.GetAllWithPagination(ctx, pagination)
	if err != nil {
		return nil, err
	}

	data, err = json.Marshal(candidateData)
	if err != nil {
		return nil, err
	}

	err = c.cache.Set(pagination.Page, data)
	if err != nil {
		return nil, err
	}

	return candidateData, nil
}
