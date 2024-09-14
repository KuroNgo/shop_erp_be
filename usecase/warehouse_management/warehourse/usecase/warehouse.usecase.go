package warehouse_usecase

import (
	"context"
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
	//TODO implement me
	panic("implement me")
}

func (w *warehouseUseCase) UpdateWarehouse(ctx context.Context, id string, input *warehousedomain.Input) error {
	//TODO implement me
	panic("implement me")
}

func (w *warehouseUseCase) GetWarehouseByName(ctx context.Context, name string) (*warehousedomain.WarehouseResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (w *warehouseUseCase) GetWarehouseByID(ctx context.Context, id string) (*warehousedomain.WarehouseResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (w *warehouseUseCase) GetAllWarehouses(ctx context.Context) ([]warehousedomain.WarehouseResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (w *warehouseUseCase) DeleteWarehouse(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
