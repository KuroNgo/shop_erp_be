package stock_adjustment_usecase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	productdomain "shop_erp_mono/domain/warehouse_management/product"
	stockadjustmentdomain "shop_erp_mono/domain/warehouse_management/stock_adjustment"
	warehousedomain "shop_erp_mono/domain/warehouse_management/warehouse"
	"shop_erp_mono/repository"
	"shop_erp_mono/usecase/warehouse_management/stock_adjustment/validate"
	"time"
)

type stockAdjustmentUseCase struct {
	contextTimeout            time.Duration
	stockAdjustmentRepository stockadjustmentdomain.IStockAdjustmentRepository
	productRepository         productdomain.IProductRepository
	warehouseRepository       warehousedomain.IWarehouseRepository
}

func NewStockAdjustmentUseCase(contextTimeout time.Duration, stockAdjustmentRepository stockadjustmentdomain.IStockAdjustmentRepository, productRepository productdomain.IProductRepository, warehouseRepository warehousedomain.IWarehouseRepository) stockadjustmentdomain.IStockAdjustmentUseCase {
	return &stockAdjustmentUseCase{contextTimeout: contextTimeout, stockAdjustmentRepository: stockAdjustmentRepository, productRepository: productRepository, warehouseRepository: warehouseRepository}
}

func (s *stockAdjustmentUseCase) GetByID(ctx context.Context, id string) (*stockadjustmentdomain.StockAdjustmentResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	stockAdjustmentID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	stockAdjustmentData, err := s.stockAdjustmentRepository.GetByID(ctx, stockAdjustmentID)
	if err != nil {
		return nil, err
	}

	response := &stockadjustmentdomain.StockAdjustmentResponse{
		StockAdjustment: *stockAdjustmentData,
	}

	return response, nil
}

func (s *stockAdjustmentUseCase) CreateOne(ctx context.Context, input *stockadjustmentdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	if err := validate.ValidateStockAdjustment(input); err != nil {
		return err
	}

	productData, err := s.productRepository.GetProductByName(ctx, input.Product)
	if err != nil {
		return err
	}

	warehouseData, err := s.warehouseRepository.GetByName(ctx, input.Warehouse)
	if err != nil {
		return err
	}

	stockAdjustment := &stockadjustmentdomain.StockAdjustment{
		ID:             primitive.NewObjectID(),
		ProductID:      productData.ID,
		WarehouseID:    warehouseData.ID,
		AdjustmentDate: input.AdjustmentDate,
		AdjustmentType: input.AdjustmentType,
		Quantity:       input.Quantity,
		Reason:         input.Reason,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	return s.stockAdjustmentRepository.CreateOne(ctx, stockAdjustment)
}

func (s *stockAdjustmentUseCase) UpdateOne(ctx context.Context, id string, input *stockadjustmentdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	if err := validate.ValidateStockAdjustment(input); err != nil {
		return err
	}

	stockAdjustmentID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	productData, err := s.productRepository.GetProductByName(ctx, input.Product)
	if err != nil {
		return err
	}

	warehouseData, err := s.warehouseRepository.GetByName(ctx, input.Warehouse)
	if err != nil {
		return err
	}

	stockAdjustment := &stockadjustmentdomain.StockAdjustment{
		ID:             stockAdjustmentID,
		ProductID:      productData.ID,
		WarehouseID:    warehouseData.ID,
		AdjustmentDate: input.AdjustmentDate,
		AdjustmentType: input.AdjustmentType,
		Quantity:       input.Quantity,
		Reason:         input.Reason,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	return s.stockAdjustmentRepository.UpdateOne(ctx, stockAdjustment)
}

func (s *stockAdjustmentUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	stockAdjustmentID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return s.stockAdjustmentRepository.DeleteOne(ctx, stockAdjustmentID)
}

func (s *stockAdjustmentUseCase) GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]stockadjustmentdomain.StockAdjustmentResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	stockAdjustmentData, err := s.stockAdjustmentRepository.GetAllWithPagination(ctx, pagination)
	if err != nil {
		return nil, err
	}

	var responses []stockadjustmentdomain.StockAdjustmentResponse
	responses = make([]stockadjustmentdomain.StockAdjustmentResponse, 0, len(stockAdjustmentData))
	for _, stockAdjustment := range stockAdjustmentData {
		productData, err := s.productRepository.GetProductByID(ctx, stockAdjustment.ProductID)
		if err != nil {
			return nil, err
		}

		warehouseData, err := s.warehouseRepository.GetByID(ctx, stockAdjustment.WarehouseID)
		if err != nil {
			return nil, err
		}

		response := stockadjustmentdomain.StockAdjustmentResponse{
			StockAdjustment: stockAdjustment,
			Product:         *productData,
			Warehouse:       *warehouseData,
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (s *stockAdjustmentUseCase) GetByProductID(ctx context.Context, productID string) ([]stockadjustmentdomain.StockAdjustmentResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	idProduct, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return nil, err
	}

	stockAdjustmentData, err := s.stockAdjustmentRepository.GetByProductID(ctx, idProduct)
	if err != nil {
		return nil, err
	}

	var responses []stockadjustmentdomain.StockAdjustmentResponse
	responses = make([]stockadjustmentdomain.StockAdjustmentResponse, 0, len(stockAdjustmentData))
	for _, stockAdjustment := range stockAdjustmentData {
		productData, err := s.productRepository.GetProductByID(ctx, stockAdjustment.ProductID)
		if err != nil {
			return nil, err
		}

		warehouseData, err := s.warehouseRepository.GetByID(ctx, stockAdjustment.WarehouseID)
		if err != nil {
			return nil, err
		}

		response := stockadjustmentdomain.StockAdjustmentResponse{
			StockAdjustment: stockAdjustment,
			Product:         *productData,
			Warehouse:       *warehouseData,
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (s *stockAdjustmentUseCase) GetByWarehouseID(ctx context.Context, warehouseID string) ([]stockadjustmentdomain.StockAdjustmentResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	idWarehouse, err := primitive.ObjectIDFromHex(warehouseID)
	if err != nil {
		return nil, err
	}

	stockAdjustmentData, err := s.stockAdjustmentRepository.GetByWarehouseID(ctx, idWarehouse)
	if err != nil {
		return nil, err
	}

	var responses []stockadjustmentdomain.StockAdjustmentResponse
	responses = make([]stockadjustmentdomain.StockAdjustmentResponse, 0, len(stockAdjustmentData))
	for _, stockAdjustment := range stockAdjustmentData {
		productData, err := s.productRepository.GetProductByID(ctx, stockAdjustment.ProductID)
		if err != nil {
			return nil, err
		}

		warehouseData, err := s.warehouseRepository.GetByID(ctx, stockAdjustment.WarehouseID)
		if err != nil {
			return nil, err
		}

		response := stockadjustmentdomain.StockAdjustmentResponse{
			StockAdjustment: stockAdjustment,
			Product:         *productData,
			Warehouse:       *warehouseData,
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (s *stockAdjustmentUseCase) GetByAdjustmentDateRange(ctx context.Context, startDate, endDate string) ([]stockadjustmentdomain.StockAdjustmentResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	layout := "06/02/2002"
	start, err := time.Parse(layout, startDate)
	if err != nil {
		return nil, err
	}

	end, err := time.Parse(layout, endDate)
	if err != nil {
		return nil, err
	}

	stockAdjustmentData, err := s.stockAdjustmentRepository.GetByAdjustmentDateRange(ctx, start, end)
	if err != nil {
		return nil, err
	}

	var responses []stockadjustmentdomain.StockAdjustmentResponse
	responses = make([]stockadjustmentdomain.StockAdjustmentResponse, 0, len(stockAdjustmentData))
	for _, stockAdjustment := range stockAdjustmentData {
		productData, err := s.productRepository.GetProductByID(ctx, stockAdjustment.ProductID)
		if err != nil {
			return nil, err
		}

		warehouseData, err := s.warehouseRepository.GetByID(ctx, stockAdjustment.WarehouseID)
		if err != nil {
			return nil, err
		}

		response := stockadjustmentdomain.StockAdjustmentResponse{
			StockAdjustment: stockAdjustment,
			Product:         *productData,
			Warehouse:       *warehouseData,
		}

		responses = append(responses, response)
	}

	return responses, nil
}
