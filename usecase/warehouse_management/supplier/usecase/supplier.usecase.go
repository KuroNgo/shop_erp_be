package supplier_usecase

import (
	"context"
	"github.com/allegro/bigcache/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	supplierdomain "shop_erp_mono/domain/warehouse_management/supplier"
	"shop_erp_mono/repository"
	"shop_erp_mono/usecase/warehouse_management/supplier/validate"
	"time"
)

type supplierUseCase struct {
	contextTimeout     time.Duration
	supplierRepository supplierdomain.ISupplierRepository
	cache              *bigcache.BigCache
}

func NewSupplierUseCase(contextTimeout time.Duration, supplierRepository supplierdomain.ISupplierRepository, cacheTTL time.Duration) supplierdomain.ISupplierUseCase {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(cacheTTL))
	if err != nil {
		return nil
	}
	return &supplierUseCase{contextTimeout: contextTimeout, cache: cache, supplierRepository: supplierRepository}
}

func (s *supplierUseCase) CreateOne(ctx context.Context, input *supplierdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	if err := validate.Supplier(input); err != nil {
		return err
	}

	supplier := supplierdomain.Supplier{
		ID:            primitive.NewObjectID(),
		SupplierName:  input.SupplierName,
		ContactPerson: input.ContactPerson,
		PhoneNumber:   input.PhoneNumber,
		Email:         input.Email,
		Address:       input.Address,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	return s.supplierRepository.CreateOne(ctx, supplier)
}

func (s *supplierUseCase) GetByID(ctx context.Context, id string) (*supplierdomain.SupplierResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	supplierID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	supplierData, err := s.supplierRepository.GetByID(ctx, supplierID)
	if err != nil {
		return nil, err
	}

	response := &supplierdomain.SupplierResponse{
		Supplier: *supplierData,
	}

	return response, nil
}

func (s *supplierUseCase) GetByName(ctx context.Context, name string) (*supplierdomain.SupplierResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	supplierData, err := s.supplierRepository.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}

	response := &supplierdomain.SupplierResponse{
		Supplier: *supplierData,
	}

	return response, nil
}

func (s *supplierUseCase) GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]supplierdomain.SupplierResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	supplierData, err := s.supplierRepository.GetAllWithPagination(ctx, pagination)
	if err != nil {
		return nil, err
	}

	var responses []supplierdomain.SupplierResponse
	responses = make([]supplierdomain.SupplierResponse, 0, len(supplierData))
	for _, supplier := range supplierData {
		response := supplierdomain.SupplierResponse{
			Supplier: supplier,
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (s *supplierUseCase) GetAll(ctx context.Context) ([]supplierdomain.SupplierResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	supplierData, err := s.supplierRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []supplierdomain.SupplierResponse
	responses = make([]supplierdomain.SupplierResponse, 0, len(supplierData))
	for _, supplier := range supplierData {
		response := supplierdomain.SupplierResponse{
			Supplier: supplier,
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (s *supplierUseCase) UpdateOne(ctx context.Context, id string, input *supplierdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	if err := validate.Supplier(input); err != nil {
		return err
	}

	supplierID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	supplier := &supplierdomain.Supplier{
		ID:            supplierID,
		SupplierName:  input.SupplierName,
		ContactPerson: input.ContactPerson,
		PhoneNumber:   input.PhoneNumber,
		Email:         input.Email,
		Address:       input.Address,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	return s.supplierRepository.UpdateOne(ctx, supplier)
}

func (s *supplierUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	supplierID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return s.supplierRepository.DeleteOne(ctx, supplierID)
}
