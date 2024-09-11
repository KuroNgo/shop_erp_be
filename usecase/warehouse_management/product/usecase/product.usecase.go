package product_usecase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	productdomain "shop_erp_mono/domain/warehouse_management/product"
	categorydomain "shop_erp_mono/domain/warehouse_management/product_category"
	"time"
)

type productUseCase struct {
	contextTimeout     time.Duration
	productRepository  productdomain.IProductRepository
	categoryRepository categorydomain.ICategoryRepository
}

func NewProductUseCase(contextTimeout time.Duration, productRepository productdomain.IProductRepository, categoryRepository categorydomain.ICategoryRepository) productdomain.IProductUseCase {
	return &productUseCase{contextTimeout: contextTimeout, productRepository: productRepository, categoryRepository: categoryRepository}
}

func (p *productUseCase) CreateProduct(ctx context.Context, input *productdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

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

	return p.productRepository.CreateProduct(ctx, product)
}

func (p *productUseCase) UpdateProduct(ctx context.Context, id string, input *productdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

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

	return p.productRepository.UpdateProduct(ctx, productID, product)
}

func (p *productUseCase) GetProductByID(ctx context.Context, id string) (*productdomain.ProductResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	productID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	productData, err := p.productRepository.GetProductByID(ctx, productID)
	if err != nil {
		return nil, err
	}

	response := &productdomain.ProductResponse{
		Product: *productData,
	}

	return response, nil
}

func (p *productUseCase) GetProductByName(ctx context.Context, productName string) (*productdomain.ProductResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	productData, err := p.productRepository.GetProductByName(ctx, productName)
	if err != nil {
		return nil, err
	}

	response := &productdomain.ProductResponse{
		Product: *productData,
	}

	return response, nil
}

func (p *productUseCase) GetAllProducts(ctx context.Context) ([]productdomain.ProductResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	productData, err := p.productRepository.GetAllProducts(ctx)
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

func (p *productUseCase) DeleteProduct(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	productID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return p.productRepository.DeleteProduct(ctx, productID)
}
