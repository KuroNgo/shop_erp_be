package category_usecase

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	productdomain "shop_erp_mono/domain/warehouse_management/product"
	categorydomain "shop_erp_mono/domain/warehouse_management/product_category"
	"time"
)

type categoryUseCase struct {
	contextTimeout     time.Duration
	categoryRepository categorydomain.ICategoryRepository
	productRepository  productdomain.IProductRepository
}

func NewCategoryUseCase(contextTimeout time.Duration, categoryRepository categorydomain.ICategoryRepository, productRepository productdomain.IProductRepository) categorydomain.ICategoryUseCase {
	return &categoryUseCase{contextTimeout: contextTimeout, categoryRepository: categoryRepository, productRepository: productRepository}
}

func (c *categoryUseCase) CreateCategory(ctx context.Context, input *categorydomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	category := categorydomain.Category{
		ID:           primitive.NewObjectID(),
		CategoryName: input.CategoryName,
		Description:  input.Description,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	return c.categoryRepository.Create(ctx, category)
}

func (c *categoryUseCase) GetByIDCategory(ctx context.Context, id string) (*categorydomain.CategoryResponse, error) {
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

func (c *categoryUseCase) GetByNameCategory(ctx context.Context, name string) (*categorydomain.CategoryResponse, error) {
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

func (c *categoryUseCase) UpdateCategory(ctx context.Context, id string, input *categorydomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	categoryID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	category := categorydomain.Category{
		CategoryName: input.CategoryName,
		Description:  input.Description,
		UpdatedAt:    time.Now(),
	}

	return c.categoryRepository.Update(ctx, categoryID, category)
}

func (c *categoryUseCase) DeleteCategory(ctx context.Context, id string) error {
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

	return c.categoryRepository.Delete(ctx, categoryID)
}

func (c *categoryUseCase) GetAllCategories(ctx context.Context) ([]categorydomain.CategoryResponse, error) {
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
