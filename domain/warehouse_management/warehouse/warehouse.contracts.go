package warehouse_domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IWarehouseRepository interface {
	CreateOne(ctx context.Context, warehouse Warehouse) error
	UpdateOne(ctx context.Context, warehouse Warehouse) error
	GetByID(ctx context.Context, id primitive.ObjectID) (*Warehouse, error)
	GetByName(ctx context.Context, name string) (*Warehouse, error)
	GetAll(ctx context.Context) ([]Warehouse, error)
	DeleteOne(ctx context.Context, id primitive.ObjectID) error
}

type IWarehouseUseCase interface {
	CreateOne(ctx context.Context, input *Input) error
	UpdateOne(ctx context.Context, id string, input *Input) error
	GetByName(ctx context.Context, name string) (*WarehouseResponse, error)
	GetByID(ctx context.Context, id string) (*WarehouseResponse, error)
	GetAll(ctx context.Context) ([]WarehouseResponse, error)
	DeleteOne(ctx context.Context, id string) error
}
