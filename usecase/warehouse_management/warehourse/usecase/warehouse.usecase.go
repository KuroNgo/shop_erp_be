package warehouse_usecase

import (
	"context"
	"github.com/allegro/bigcache/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	warehousedomain "shop_erp_mono/domain/warehouse_management/warehouse"
	"shop_erp_mono/usecase/warehouse_management/warehourse/validate"
	"time"
)

type warehouseUseCase struct {
	contextTimeout      time.Duration
	warehouseRepository warehousedomain.IWarehouseRepository
	cache               *bigcache.BigCache
}

func NewWarehouseUseCase(contextTimeout time.Duration, warehouseRepository warehousedomain.IWarehouseRepository, cacheTTL time.Duration) warehousedomain.IWarehouseUseCase {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(cacheTTL))
	if err != nil {
		return nil
	}
	return &warehouseUseCase{contextTimeout: contextTimeout, cache: cache, warehouseRepository: warehouseRepository}
}

func (w *warehouseUseCase) CreateOne(ctx context.Context, input *warehousedomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, w.contextTimeout)
	defer cancel()

	if err := validate.Warehouse(input); err != nil {
		return err
	}

	warehouse := warehousedomain.Warehouse{
		ID:            primitive.NewObjectID(),
		WarehouseName: input.WarehouseName,
		Location:      input.Location,
		Capacity:      input.Capacity,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	return w.warehouseRepository.CreateOne(ctx, warehouse)
}

func (w *warehouseUseCase) UpdateOne(ctx context.Context, id string, input *warehousedomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, w.contextTimeout)
	defer cancel()

	if err := validate.Warehouse(input); err != nil {
		return err
	}

	warehouseID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	warehouse := warehousedomain.Warehouse{
		ID:            warehouseID,
		WarehouseName: input.WarehouseName,
		Location:      input.Location,
		Capacity:      input.Capacity,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	return w.warehouseRepository.UpdateOne(ctx, warehouse)
}

func (w *warehouseUseCase) GetByName(ctx context.Context, name string) (*warehousedomain.WarehouseResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, w.contextTimeout)
	defer cancel()

	warehouseData, err := w.warehouseRepository.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}

	response := &warehousedomain.WarehouseResponse{
		Warehouse: *warehouseData,
	}
	return response, nil
}

func (w *warehouseUseCase) GetByID(ctx context.Context, id string) (*warehousedomain.WarehouseResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, w.contextTimeout)
	defer cancel()

	warehouseID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	warehouseData, err := w.warehouseRepository.GetByID(ctx, warehouseID)
	if err != nil {
		return nil, err
	}

	response := &warehousedomain.WarehouseResponse{
		Warehouse: *warehouseData,
	}
	return response, nil
}

func (w *warehouseUseCase) GetAll(ctx context.Context) ([]warehousedomain.WarehouseResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, w.contextTimeout)
	defer cancel()

	warehouseData, err := w.warehouseRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []warehousedomain.WarehouseResponse
	responses = make([]warehousedomain.WarehouseResponse, 0, len(warehouseData))
	for _, warehouse := range warehouseData {
		response := warehousedomain.WarehouseResponse{
			Warehouse: warehouse,
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (w *warehouseUseCase) DeleteOne(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, w.contextTimeout)
	defer cancel()

	warehouseID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return w.warehouseRepository.DeleteOne(ctx, warehouseID)
}
