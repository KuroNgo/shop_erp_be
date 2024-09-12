package warehouse_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IWarehouseRepository interface {
	CreateOne(ctx context.Context, warehouse Warehouse) (*Warehouse, error)
	UpdateOne(ctx context.Context, id primitive.ObjectID, warehouse Warehouse) (*Warehouse, error)
	GetByID(ctx context.Context, id primitive.ObjectID) (*WarehouseResponse, error)
	GetByName(ctx context.Context, name string) (*WarehouseResponse, error)
	GetAll(ctx context.Context) ([]WarehouseResponse, error)
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
}

type IWarehouseUseCase interface {
	CreateWarehouse(ctx context.Context, input *Input) (*WarehouseResponse, error)
	UpdateWarehouse(ctx context.Context, id string, input *Input) (*WarehouseResponse, error)
	GetWarehouseByName(ctx context.Context, name string) (*WarehouseResponse, error)
	GetWarehouseByID(ctx context.Context, id string) (*WarehouseResponse, error)
	GetAllWarehouses(ctx context.Context) ([]WarehouseResponse, error)
	DeleteWarehouse(ctx context.Context, id string) error
}
