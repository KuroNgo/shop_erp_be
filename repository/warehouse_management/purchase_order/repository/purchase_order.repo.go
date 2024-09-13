package purchase_order_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	purchaseorderdomain "shop_erp_mono/domain/warehouse_management/purchase_order"
	"shop_erp_mono/repository"
	"time"
)

type purchaseOrderRepository struct {
	database                *mongo.Database
	purchaseOrderCollection string
}

func NewPurchaseOrderRepository(database *mongo.Database, purchaseOrderCollection string) purchaseorderdomain.IPurchaseOrderRepository {
	return &purchaseOrderRepository{database: database, purchaseOrderCollection: purchaseOrderCollection}
}

func (p *purchaseOrderRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*purchaseorderdomain.PurchaseOrder, error) {
	purchaseOrderCollection := p.database.Collection(p.purchaseOrderCollection)

	filter := bson.M{"_id": id}
	var purchaseOrder purchaseorderdomain.PurchaseOrder
	if err := purchaseOrderCollection.FindOne(ctx, filter).Decode(&purchaseOrder); err != nil {
		return nil, err
	}

	return &purchaseOrder, nil
}

func (p *purchaseOrderRepository) Create(ctx context.Context, order *purchaseorderdomain.PurchaseOrder) error {
	purchaseOrderCollection := p.database.Collection(p.purchaseOrderCollection)

	_, err := purchaseOrderCollection.InsertOne(ctx, order)
	if err != nil {
		return err
	}

	return nil
}

func (p *purchaseOrderRepository) Update(ctx context.Context, id primitive.ObjectID, order *purchaseorderdomain.PurchaseOrder) error {
	purchaseOrderCollection := p.database.Collection(p.purchaseOrderCollection)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"order_number": order.OrderNumber,
		"supplier_id":  order.SupplierID,
		"order_date":   order.OrderDate,
		"total_amount": order.TotalAmount,
		"status":       order.Status,
		"updated_at":   time.Now(),
	}}
	_, err := purchaseOrderCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (p *purchaseOrderRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	purchaseOrderCollection := p.database.Collection(p.purchaseOrderCollection)

	filter := bson.M{"_id": id}
	_, err := purchaseOrderCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (p *purchaseOrderRepository) GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]purchaseorderdomain.PurchaseOrder, error) {
	purchaseOrderCollection := p.database.Collection(p.purchaseOrderCollection)

	filter := bson.M{}
	cursor, err := repository.Paginate(ctx, purchaseOrderCollection, filter, pagination)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var purchaseOrders []purchaseorderdomain.PurchaseOrder
	purchaseOrders = make([]purchaseorderdomain.PurchaseOrder, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var purchaseOrder purchaseorderdomain.PurchaseOrder
		if err = cursor.Decode(&purchaseOrder); err != nil {
			return nil, err
		}

		purchaseOrders = append(purchaseOrders, purchaseOrder)
	}

	return purchaseOrders, nil
}

func (p *purchaseOrderRepository) GetBySupplierID(ctx context.Context, supplierID primitive.ObjectID) ([]purchaseorderdomain.PurchaseOrder, error) {
	purchaseOrderCollection := p.database.Collection(p.purchaseOrderCollection)

	filter := bson.M{"supplier_id": supplierID}
	cursor, err := purchaseOrderCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var purchaseOrders []purchaseorderdomain.PurchaseOrder
	purchaseOrders = make([]purchaseorderdomain.PurchaseOrder, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var purchaseOrder purchaseorderdomain.PurchaseOrder
		if err = cursor.Decode(&purchaseOrder); err != nil {
			return nil, err
		}

		purchaseOrders = append(purchaseOrders, purchaseOrder)
	}

	return purchaseOrders, nil
}

func (p *purchaseOrderRepository) UpdateStatus(ctx context.Context, id primitive.ObjectID, status string) error {
	purchaseOrderCollection := p.database.Collection(p.purchaseOrderCollection)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"status":     status,
		"updated_at": time.Now(),
	}}
	_, err := purchaseOrderCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}
