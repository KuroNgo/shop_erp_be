package warehouse_usecase

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	warehousedomain "shop_erp_mono/domain/warehouse_management/warehouse"
	"time"
)

type warehouseUseCase struct {
	contextTimeout      time.Duration
	warehouseRepository warehousedomain.IWarehouseRepository
}

func NewWarehouseUseCase(contextTimeout time.Duration, warehouseRepository warehousedomain.IWarehouseRepository) warehousedomain.IWarehouseUseCase {
	return &warehouseUseCase{contextTimeout: contextTimeout, warehouseRepository: warehouseRepository}
}

func (w *warehouseUseCase) CreateWarehouse(ctx context.Context, input *warehousedomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, w.contextTimeout)
	defer cancel()

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

func (w *warehouseUseCase) UpdateWarehouse(ctx context.Context, id string, input *warehousedomain.Input) error {
	ctx, cancel := context.WithTimeout(ctx, w.contextTimeout)
	defer cancel()

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

func (w *warehouseUseCase) GetWarehouseByName(ctx context.Context, name string) (*warehousedomain.WarehouseResponse, error) {
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

func (w *warehouseUseCase) GetWarehouseByID(ctx context.Context, id string) (*warehousedomain.WarehouseResponse, error) {
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

func (w *warehouseUseCase) GetAllWarehouses(ctx context.Context) ([]warehousedomain.WarehouseResponse, error) {
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

func (w *warehouseUseCase) DeleteWarehouse(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, w.contextTimeout)
	defer cancel()

	warehouseID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return w.warehouseRepository.DeleteOne(ctx, warehouseID)
}
