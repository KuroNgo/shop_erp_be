package stockmovement_usecase

import (
	"context"
	"github.com/allegro/bigcache/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	userdomain "shop_erp_mono/domain/human_resource_management/user"
	productdomain "shop_erp_mono/domain/warehouse_management/product"
	stockmovementdomain "shop_erp_mono/domain/warehouse_management/stockmovement"
	warehousedomain "shop_erp_mono/domain/warehouse_management/warehouse"
	"shop_erp_mono/repository"
	"time"
)

type stockMovementUseCase struct {
	contextTimeout          time.Duration
	stockMovementRepository stockmovementdomain.IStockMovementRepository
	productRepository       productdomain.IProductRepository
	warehouseRepository     warehousedomain.IWarehouseRepository
	userRepository          userdomain.IUserRepository
	cache                   *bigcache.BigCache
}

func NewStockMovementUseCase(contextTimeout time.Duration, stockMovementRepository stockmovementdomain.IStockMovementRepository,
	productRepository productdomain.IProductRepository, userRepository userdomain.IUserRepository, warehouseRepository warehousedomain.IWarehouseRepository,
	cacheTTL time.Duration) stockmovementdomain.IStockMovementUseCase {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(cacheTTL))
	if err != nil {
		return nil
	}
	return &stockMovementUseCase{contextTimeout: contextTimeout, cache: cache, stockMovementRepository: stockMovementRepository,
		productRepository: productRepository, userRepository: userRepository, warehouseRepository: warehouseRepository}
}

func (s *stockMovementUseCase) GetByID(ctx context.Context, id string) (*stockmovementdomain.StockMovementResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	stockMovementID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	stockMovementData, err := s.stockMovementRepository.GetByID(ctx, stockMovementID)
	if err != nil {
		return nil, err
	}

	productData, err := s.productRepository.GetByID(ctx, stockMovementData.ProductID)
	if err != nil {
		return nil, err
	}

	warehouseData, err := s.warehouseRepository.GetByID(ctx, stockMovementData.WarehouseID)
	if err != nil {
		return nil, err
	}

	userData, err := s.userRepository.GetByID(ctx, stockMovementData.UserID)
	if err != nil {
		return nil, err
	}

	response := &stockmovementdomain.StockMovementResponse{
		StockMovement: *stockMovementData,
		Product:       *productData,
		Warehouse:     *warehouseData,
		User:          userData,
	}

	return response, nil
}

func (s *stockMovementUseCase) CreateOne(ctx context.Context, input *stockmovementdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	productData, err := s.productRepository.GetByName(ctx, input.Product)
	if err != nil {
		return err
	}

	warehouseData, err := s.warehouseRepository.GetByName(ctx, input.Warehouse)
	if err != nil {
		return err
	}

	userData, err := s.userRepository.GetByEmail(ctx, input.User)
	if err != nil {
		return err
	}

	stockMovement := &stockmovementdomain.StockMovement{
		ID:           primitive.NewObjectID(),
		ProductID:    productData.ID,
		WarehouseID:  warehouseData.ID,
		UserID:       userData.ID,
		MovementDate: input.MovementDate,
		MovementType: input.MovementType,
		Quantity:     input.Quantity,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	return s.stockMovementRepository.CreateOne(ctx, stockMovement)
}

func (s *stockMovementUseCase) UpdateOne(ctx context.Context, id string, input *stockmovementdomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	stockMovementID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	productData, err := s.productRepository.GetByName(ctx, input.Product)
	if err != nil {
		return err
	}

	warehouseData, err := s.warehouseRepository.GetByName(ctx, input.Warehouse)
	if err != nil {
		return err
	}

	userData, err := s.userRepository.GetByEmail(ctx, input.User)
	if err != nil {
		return err
	}

	stockMovement := &stockmovementdomain.StockMovement{
		ID:           stockMovementID,
		ProductID:    productData.ID,
		WarehouseID:  warehouseData.ID,
		UserID:       userData.ID,
		MovementDate: input.MovementDate,
		MovementType: input.MovementType,
		Quantity:     input.Quantity,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	return s.stockMovementRepository.UpdateOne(ctx, stockMovement)
}

func (s *stockMovementUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	stockMovementID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return s.stockMovementRepository.DeleteOne(ctx, stockMovementID)
}

func (s *stockMovementUseCase) GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]stockmovementdomain.StockMovementResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	stockMovementData, err := s.stockMovementRepository.GetAllWithPagination(ctx, pagination)
	if err != nil {
		return nil, err
	}

	var responses []stockmovementdomain.StockMovementResponse
	responses = make([]stockmovementdomain.StockMovementResponse, 0, len(stockMovementData))
	for _, stockMovement := range stockMovementData {
		productData, err := s.productRepository.GetByID(ctx, stockMovement.ProductID)
		if err != nil {
			return nil, err
		}

		warehouseData, err := s.warehouseRepository.GetByID(ctx, stockMovement.WarehouseID)
		if err != nil {
			return nil, err
		}

		userData, err := s.userRepository.GetByID(ctx, stockMovement.UserID)
		if err != nil {
			return nil, err
		}

		response := stockmovementdomain.StockMovementResponse{
			StockMovement: stockMovement,
			Product:       *productData,
			Warehouse:     *warehouseData,
			User:          userData,
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (s *stockMovementUseCase) GetByProductID(ctx context.Context, productID string) ([]stockmovementdomain.StockMovementResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	idProduct, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return nil, err
	}

	stockMovementData, err := s.stockMovementRepository.GetByProductID(ctx, idProduct)
	if err != nil {
		return nil, err
	}

	var responses []stockmovementdomain.StockMovementResponse
	responses = make([]stockmovementdomain.StockMovementResponse, 0, len(stockMovementData))
	for _, stockMovement := range stockMovementData {
		productData, err := s.productRepository.GetByID(ctx, stockMovement.ProductID)
		if err != nil {
			return nil, err
		}

		warehouseData, err := s.warehouseRepository.GetByID(ctx, stockMovement.WarehouseID)
		if err != nil {
			return nil, err
		}

		userData, err := s.userRepository.GetByID(ctx, stockMovement.UserID)
		if err != nil {
			return nil, err
		}

		response := stockmovementdomain.StockMovementResponse{
			StockMovement: stockMovement,
			Product:       *productData,
			Warehouse:     *warehouseData,
			User:          userData,
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (s *stockMovementUseCase) GetByWarehouseID(ctx context.Context, warehouseID string) ([]stockmovementdomain.StockMovementResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	idWarehouse, err := primitive.ObjectIDFromHex(warehouseID)
	if err != nil {
		return nil, err
	}

	stockMovementData, err := s.stockMovementRepository.GetByProductID(ctx, idWarehouse)
	if err != nil {
		return nil, err
	}

	var responses []stockmovementdomain.StockMovementResponse
	responses = make([]stockmovementdomain.StockMovementResponse, 0, len(stockMovementData))
	for _, stockMovement := range stockMovementData {
		productData, err := s.productRepository.GetByID(ctx, stockMovement.ProductID)
		if err != nil {
			return nil, err
		}

		warehouseData, err := s.warehouseRepository.GetByID(ctx, stockMovement.WarehouseID)
		if err != nil {
			return nil, err
		}

		userData, err := s.userRepository.GetByID(ctx, stockMovement.UserID)
		if err != nil {
			return nil, err
		}

		response := stockmovementdomain.StockMovementResponse{
			StockMovement: stockMovement,
			Product:       *productData,
			Warehouse:     *warehouseData,
			User:          userData,
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (s *stockMovementUseCase) GetByUserID(ctx context.Context, userID string) ([]stockmovementdomain.StockMovementResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	idUser, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	stockMovementData, err := s.stockMovementRepository.GetByProductID(ctx, idUser)
	if err != nil {
		return nil, err
	}

	var responses []stockmovementdomain.StockMovementResponse
	responses = make([]stockmovementdomain.StockMovementResponse, 0, len(stockMovementData))
	for _, stockMovement := range stockMovementData {
		productData, err := s.productRepository.GetByID(ctx, stockMovement.ProductID)
		if err != nil {
			return nil, err
		}

		warehouseData, err := s.warehouseRepository.GetByID(ctx, stockMovement.WarehouseID)
		if err != nil {
			return nil, err
		}

		userData, err := s.userRepository.GetByID(ctx, stockMovement.UserID)
		if err != nil {
			return nil, err
		}

		response := stockmovementdomain.StockMovementResponse{
			StockMovement: stockMovement,
			Product:       *productData,
			Warehouse:     *warehouseData,
			User:          userData,
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (s *stockMovementUseCase) GetByMovementDateRange(ctx context.Context, startDate, endDate string) ([]stockmovementdomain.StockMovementResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.contextTimeout)
	defer cancel()

	layout := "02/06/2002"
	start, err := time.Parse(layout, startDate)
	if err != nil {
		return nil, err
	}

	end, err := time.Parse(layout, endDate)
	if err != nil {
		return nil, err
	}

	stockMovementData, err := s.stockMovementRepository.GetByMovementDateRange(ctx, start, end)
	if err != nil {
		return nil, err
	}

	var responses []stockmovementdomain.StockMovementResponse
	responses = make([]stockmovementdomain.StockMovementResponse, 0, len(stockMovementData))
	for _, stockMovement := range stockMovementData {
		productData, err := s.productRepository.GetByID(ctx, stockMovement.ProductID)
		if err != nil {
			return nil, err
		}

		warehouseData, err := s.warehouseRepository.GetByID(ctx, stockMovement.WarehouseID)
		if err != nil {
			return nil, err
		}

		userData, err := s.userRepository.GetByID(ctx, stockMovement.UserID)
		if err != nil {
			return nil, err
		}

		response := stockmovementdomain.StockMovementResponse{
			StockMovement: stockMovement,
			Product:       *productData,
			Warehouse:     *warehouseData,
			User:          userData,
		}

		responses = append(responses, response)
	}

	return responses, nil
}
