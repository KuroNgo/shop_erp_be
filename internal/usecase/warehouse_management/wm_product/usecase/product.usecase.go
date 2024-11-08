package product_usecase

import (
	"context"
	"github.com/allegro/bigcache/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	productdomain "shop_erp_mono/internal/domain/warehouse_management/product"
	categorydomain "shop_erp_mono/internal/domain/warehouse_management/product_category"
	"shop_erp_mono/internal/usecase/warehouse_management/wm_product/validate"
	"time"
)

type productUseCase struct {
	contextTimeout     time.Duration
	productRepository  productdomain.IProductRepository
	categoryRepository categorydomain.ICategoryRepository
	cache              *bigcache.BigCache
}

func NewProductUseCase(contextTimeout time.Duration, productRepository productdomain.IProductRepository,
	categoryRepository categorydomain.ICategoryRepository, cacheTTL time.Duration) productdomain.IProductUseCase {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(cacheTTL))
	if err != nil {
		return nil
	}
	return &productUseCase{contextTimeout: contextTimeout, cache: cache, productRepository: productRepository, categoryRepository: categoryRepository}
}

func (p *productUseCase) CreateOne(ctx context.Context, input *productdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	if err := validate.Product(input); err != nil {
		return err
	}

	categoryData, err := p.categoryRepository.GetByName(ctx, input.Category)
	if err != nil {
		return err
	}

	product := productdomain.Product{
		ID:              primitive.NewObjectID(),
		CategoryID:      categoryData.ID,
		ProductName:     input.ProductName,
		Description:     input.Description,
		Price:           input.Price,
		QuantityInStock: input.QuantityInStock,
		UpdatedAt:       time.Now(),
		CreatedAt:       time.Now(),
	}

	return p.productRepository.CreateOne(ctx, product)
}

func (p *productUseCase) UpdateOne(ctx context.Context, id string, input *productdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	if err := validate.Product(input); err != nil {
		return err
	}

	productID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	categoryData, err := p.categoryRepository.GetByName(ctx, input.Category)
	if err != nil {
		return err
	}

	product := productdomain.Product{
		ID:              primitive.NewObjectID(),
		CategoryID:      categoryData.ID,
		ProductName:     input.ProductName,
		Description:     input.Description,
		Price:           input.Price,
		QuantityInStock: input.QuantityInStock,
		UpdatedAt:       time.Now(),
		CreatedAt:       time.Now(),
	}

	return p.productRepository.UpdateOne(ctx, productID, product)
}

func (p *productUseCase) GetByID(ctx context.Context, id string) (*productdomain.ProductResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	productID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	productData, err := p.productRepository.GetByID(ctx, productID)
	if err != nil {
		return nil, err
	}

	response := &productdomain.ProductResponse{
		Product: *productData,
	}

	return response, nil
}

func (p *productUseCase) GetByName(ctx context.Context, productName string) (*productdomain.ProductResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	productData, err := p.productRepository.GetByName(ctx, productName)
	if err != nil {
		return nil, err
	}

	response := &productdomain.ProductResponse{
		Product: *productData,
	}

	return response, nil
}

func (p *productUseCase) GetAll(ctx context.Context) ([]productdomain.ProductResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	productData, err := p.productRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []productdomain.ProductResponse
	responses = make([]productdomain.ProductResponse, 0, len(productData))
	for _, product := range productData {
		response := productdomain.ProductResponse{
			Product: product,
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (p *productUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	productID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return p.productRepository.DeleteOne(ctx, productID)
}
