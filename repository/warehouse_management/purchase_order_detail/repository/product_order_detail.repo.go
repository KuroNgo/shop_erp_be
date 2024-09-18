package purchase_order_detail_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	purchaseorderdetaildomain "shop_erp_mono/domain/warehouse_management/purchase_order_detail"
	"shop_erp_mono/repository"
)

type purchaseOrderDetailRepository struct {
	database                *mongo.Database
	purchaseOrderCollection string
}

func NewPurchaseOrderDetailRepository(database *mongo.Database, purchaseOrderCollection string) purchaseorderdetaildomain.IPurchaseOrderDetailRepository {
	return &purchaseOrderDetailRepository{database: database, purchaseOrderCollection: purchaseOrderCollection}
}

func (p *purchaseOrderDetailRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*purchaseorderdetaildomain.PurchaseOrderDetail, error) {
	purchaseOrderCollection := p.database.Collection(p.purchaseOrderCollection)

	filter := bson.M{"_id": id}
	var purchaseOrderDetail purchaseorderdetaildomain.PurchaseOrderDetail
	if err := purchaseOrderCollection.FindOne(ctx, filter).Decode(&purchaseOrderDetail); err != nil {
		return nil, err
	}

	return &purchaseOrderDetail, nil
}

func (p *purchaseOrderDetailRepository) GetByPurchaseOrderID(ctx context.Context, purchaseOrderID primitive.ObjectID) ([]purchaseorderdetaildomain.PurchaseOrderDetail, error) {
	purchaseOrderCollection := p.database.Collection(p.purchaseOrderCollection)

	filter := bson.M{"purchase_order_id": purchaseOrderID}
	cursor, err := purchaseOrderCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var purchaseOrderDetails []purchaseorderdetaildomain.PurchaseOrderDetail
	purchaseOrderDetails = make([]purchaseorderdetaildomain.PurchaseOrderDetail, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var purchaseOrderDetail purchaseorderdetaildomain.PurchaseOrderDetail
		if err = cursor.Decode(&purchaseOrderDetail); err != nil {
			return nil, err
		}

		purchaseOrderDetails = append(purchaseOrderDetails, purchaseOrderDetail)
	}

	return purchaseOrderDetails, nil
}

func (p *purchaseOrderDetailRepository) Create(ctx context.Context, detail *purchaseorderdetaildomain.PurchaseOrderDetail) error {
	purchaseOrderCollection := p.database.Collection(p.purchaseOrderCollection)

	_, err := purchaseOrderCollection.InsertOne(ctx, detail)
	if err != nil {
		return err
	}

	return nil
}

func (p *purchaseOrderDetailRepository) Update(ctx context.Context, detail *purchaseorderdetaildomain.PurchaseOrderDetail) error {
	//TODO implement me
	panic("implement me")
}

func (p *purchaseOrderDetailRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	purchaseOrderCollection := p.database.Collection(p.purchaseOrderCollection)

	filter := bson.M{"_id": id}
	_, err := purchaseOrderCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (p *purchaseOrderDetailRepository) GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]purchaseorderdetaildomain.PurchaseOrderDetail, error) {
	purchaseOrderCollection := p.database.Collection(p.purchaseOrderCollection)

	filter := bson.M{}
	cursor, err := repository.Paginate(ctx, purchaseOrderCollection, filter, pagination)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	var purchaseOrderDetails []purchaseorderdetaildomain.PurchaseOrderDetail
	purchaseOrderDetails = make([]purchaseorderdetaildomain.PurchaseOrderDetail, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var purchaseOrderDetail purchaseorderdetaildomain.PurchaseOrderDetail
		if err = cursor.Decode(&purchaseOrderDetail); err != nil {
			return nil, err
		}

		purchaseOrderDetails = append(purchaseOrderDetails, purchaseOrderDetail)
	}

	return purchaseOrderDetails, nil
}

func (p *purchaseOrderDetailRepository) GetAll(ctx context.Context) ([]purchaseorderdetaildomain.PurchaseOrderDetail, error) {
	purchaseOrderCollection := p.database.Collection(p.purchaseOrderCollection)

	filter := bson.M{}
	cursor, err := purchaseOrderCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	var purchaseOrderDetails []purchaseorderdetaildomain.PurchaseOrderDetail
	purchaseOrderDetails = make([]purchaseorderdetaildomain.PurchaseOrderDetail, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var purchaseOrderDetail purchaseorderdetaildomain.PurchaseOrderDetail
		if err = cursor.Decode(&purchaseOrderDetail); err != nil {
			return nil, err
		}

		purchaseOrderDetails = append(purchaseOrderDetails, purchaseOrderDetail)
	}

	return purchaseOrderDetails, nil
}
