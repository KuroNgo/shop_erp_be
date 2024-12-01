package candidate_usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/allegro/bigcache/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongo_driven "go.mongodb.org/mongo-driver/mongo"
	"log"
	candidatedomain "shop_erp_mono/internal/domain/human_resource_management/candidate"
	employeesdomain "shop_erp_mono/internal/domain/human_resource_management/employees"
	"shop_erp_mono/internal/repository"
	"shop_erp_mono/internal/usecase/human_resource_management/candidate/validate"
	"strconv"
	"strings"
	"time"
)

type candidateUseCase struct {
	contextTimeout      time.Duration
	candidateRepository candidatedomain.ICandidateRepository
	employeeRepository  employeesdomain.IEmployeeRepository
	cache               *bigcache.BigCache
	client              *mongo_driven.Client
}

func NewCandidateUseCase(contextTimeout time.Duration, candidateRepository candidatedomain.ICandidateRepository,
	employeeRepository employeesdomain.IEmployeeRepository, cacheTTL time.Duration,
	client *mongo_driven.Client) candidatedomain.ICandidateUseCase {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(cacheTTL))
	if err != nil {
		return nil
	}
	return &candidateUseCase{contextTimeout: contextTimeout, cache: cache, candidateRepository: candidateRepository,
		employeeRepository: employeeRepository, client: client}
}

func (c *candidateUseCase) CreateOne(ctx context.Context, candidate *candidatedomain.Candidate) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	if err := validate.Candidate(candidate); err != nil {
		return err
	}

	err := c.candidateRepository.CreateOne(ctx, candidate)
	if err != nil {
		return err
	}

	if err := c.cache.Delete("candidates"); err != nil {
		log.Printf("failed to delete candidates cache: %v", err)
	}

	return nil
}

func (c *candidateUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	candidateID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	err = c.candidateRepository.DeleteOne(ctx, candidateID)
	if err != nil {
		return err
	}

	err = c.cache.Delete("candidates")
	if err != nil {
		log.Printf("failed to delete candidates cache: %v", err)
	}
	err = c.cache.Delete(id)
	if err != nil {
		log.Printf("failed to delete candidates cache: %v", err)
	}

	return nil
}

func (c *candidateUseCase) UpdateOne(ctx context.Context, id string, candidate *candidatedomain.Candidate) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	candidateID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	candidateData, err := c.candidateRepository.GetByID(ctx, candidateID)
	if err != nil {
		return err
	}

	if candidateData.Status == "onboarding" {
		return errors.New("can not update with candidate have done process")
	}

	err = c.candidateRepository.UpdateOne(ctx, candidateID, candidate)
	if err != nil {
		return err
	}

	err = c.cache.Delete("candidates")
	if err != nil {
		log.Printf("failed to delete candidates cache: %v", err)
	}
	err = c.cache.Delete(id)
	if err != nil {
		log.Printf("failed to delete candidates cache: %v", err)
	}

	return nil
}

func (c *candidateUseCase) UpdateStatus(ctx context.Context, id string, status string) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	candidateID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	candidateData, err := c.candidateRepository.GetByID(ctx, candidateID)
	if err != nil {
		return err
	}

	if candidateData.Status == "onboarding" {
		return errors.New("can not update with candidate have done process")
	}

	session, err := c.client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	callback := func(sessionCtx mongo_driven.SessionContext) (interface{}, error) {
		if status == "onboarding" {
			// Get a information candidate
			candidate, err := c.candidateRepository.GetByID(sessionCtx, candidateID)
			if err != nil {
				return nil, err
			}

			firstname, lastname, err := c.splitFullName(candidate.Name)
			if err != nil {
				return nil, err

			}

			email, err := c.createEmailEmployee(candidate.Name)
			if err != nil {
				return nil, err
			}

			// From candidate's information to employee
			employee := employeesdomain.Employee{
				FirstName: firstname,
				LastName:  lastname,
				Gender:    candidate.Gender,
				Address:   candidate.Address,
				AvatarURL: candidate.ImageURL,
				Email:     email,
				Phone:     candidate.Phone,
				RoleID:    candidate.RoleHire,
				//DepartmentID: candidateID., get info from roleID
				//SalaryID: candidate., get info from roleID
				Active:    "active",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}

			err = c.employeeRepository.CreateOne(sessionCtx, &employee)
			if err != nil {
				return nil, err
			}
		}

		return nil, err
	}

	// Run the transaction
	_, err = session.WithTransaction(ctx, callback)
	if err != nil {
		return err
	}

	err = session.CommitTransaction(ctx)
	if err != nil {
		return err
	}

	err = c.candidateRepository.UpdateStatus(ctx, candidateID, status)
	if err != nil {
		return err
	}

	err = c.cache.Delete("candidates")
	if err != nil {
		log.Printf("failed to delete candidates cache: %v", err)
	}
	err = c.cache.Delete(id)
	if err != nil {
		log.Printf("failed to delete candidates cache: %v", err)
	}

	return nil
}

func (c *candidateUseCase) splitFullName(name string) (string, string, error) {
	// Split the full name into an array of words
	fullName := strings.Split(name, " ")
	if len(fullName) < 2 {
		return "", "", fmt.Errorf("name must contain at least first and last name")
	}

	lastName := fullName[len(fullName)-1]
	firstAndMiddleName := strings.Join(fullName[:len(fullName)-1], " ")

	return firstAndMiddleName, lastName, nil
}

func (c *candidateUseCase) createEmailEmployee(name string) (string, error) {
	fullName := strings.Split(name, " ")
	lastName := fullName[len(fullName)-1]

	initials := ""
	for i := 0; i < len(fullName)-1; i++ {
		initials += string(fullName[i][0])
	}

	email := strings.ToLower(lastName + initials + "@gmail.com")
	counter := 0

	for {
		count, err := c.employeeRepository.CountEmployeeByEmail(context.Background(), email)
		if err != nil {
			return "", err
		}

		if count == 0 {
			break
		}

		counter++
		email = strings.ToLower(lastName + initials + strconv.Itoa(counter) + "@gmail.com")
	}

	return email, nil
}

func (c *candidateUseCase) GetByID(ctx context.Context, id string) (*candidatedomain.Candidate, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	data, err := c.cache.Get(id)
	if err != nil {
		log.Printf("failed to get candidates cache: %v", err)
	}
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
		log.Printf("failed to set candidates cache: %v", err)
	}

	return response, nil
}

func (c *candidateUseCase) GetByEmail(ctx context.Context, email string) (*candidatedomain.Candidate, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	data, err := c.cache.Get(email)
	if err != nil {
		log.Printf("failed to get candidates cache: %v", err)
	}
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
		log.Printf("failed to set candidates cache: %v", err)
	}

	return candidateData, nil
}

func (c *candidateUseCase) GetAll(ctx context.Context) ([]candidatedomain.Candidate, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	data, err := c.cache.Get("candidates")
	if err != nil {
		log.Printf("failed to get candidates cache: %v", err)
	}
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
		log.Printf("failed to set candidates cache: %v", err)
	}

	return candidateData, nil
}

func (c *candidateUseCase) GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]candidatedomain.Candidate, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	data, err := c.cache.Get(pagination.Page)
	if err != nil {
		log.Printf("failed to get candidates cache: %v", err)
	}
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
		log.Printf("failed to set candidates cache: %v", err)
	}

	return candidateData, nil
}
