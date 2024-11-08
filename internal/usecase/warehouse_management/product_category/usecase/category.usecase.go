package category_usecase

import (
	"context"
	"errors"
	"github.com/allegro/bigcache/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	productdomain "shop_erp_mono/internal/domain/warehouse_management/product"
	categorydomain "shop_erp_mono/internal/domain/warehouse_management/product_category"
	"shop_erp_mono/internal/usecase/warehouse_management/product_category/validate"
	"time"
)

type categoryUseCase struct {
	contextTimeout     time.Duration
	categoryRepository categorydomain.ICategoryRepository
	productRepository  productdomain.IProductRepository
	cache              *bigcache.BigCache
}

func NewCategoryUseCase(contextTimeout time.Duration, categoryRepository categorydomain.ICategoryRepository,
	productRepository productdomain.IProductRepository, cacheTTL time.Duration) categorydomain.ICategoryUseCase {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(cacheTTL))
	if err != nil {
		return nil
	}
	return &categoryUseCase{contextTimeout: contextTimeout, cache: cache, categoryRepository: categoryRepository, productRepository: productRepository}
}

func (c *categoryUseCase) CreateOne(ctx context.Context, input *categorydomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	if err := validate.Category(input); err != nil {
		return err
	}

	category := categorydomain.Category{
		ID:           primitive.NewObjectID(),
		CategoryName: input.CategoryName,
		Description:  input.Description,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	return c.categoryRepository.CreateOne(ctx, category)
}

func (c *categoryUseCase) GetByID(ctx context.Context, id string) (*categorydomain.CategoryResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	categoryID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	categoryData, err := c.categoryRepository.GetByID(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	response := &categorydomain.CategoryResponse{
		Category: *categoryData,
	}

	return response, nil
}

func (c *categoryUseCase) GetByName(ctx context.Context, name string) (*categorydomain.CategoryResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	categoryData, err := c.categoryRepository.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}

	response := &categorydomain.CategoryResponse{
		Category: *categoryData,
	}

	return response, nil
}

func (c *categoryUseCase) UpdateOne(ctx context.Context, id string, input *categorydomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	if err := validate.Category(input); err != nil {
		return err
	}

	categoryID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	category := categorydomain.Category{
		CategoryName: input.CategoryName,
		Description:  input.Description,
		UpdatedAt:    time.Now(),
	}

	return c.categoryRepository.UpdateOne(ctx, categoryID, category)
}

func (c *categoryUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	categoryID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	productCount, err := c.productRepository.CountCategory(ctx, categoryID)
	if err != nil {
		return err
	}

	if productCount > 0 {
		return errors.New("cannot delete product_category with linked products")
	}

	return c.categoryRepository.DeleteOne(ctx, categoryID)
}

func (c *categoryUseCase) GetAll(ctx context.Context) ([]categorydomain.CategoryResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	categoryData, err := c.categoryRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []categorydomain.CategoryResponse
	responses = make([]categorydomain.CategoryResponse, 0, len(categoryData))
	for _, category := range categoryData {
		response := categorydomain.CategoryResponse{
			Category: category,
		}

		responses = append(responses, response)
	}

	return responses, nil
}
