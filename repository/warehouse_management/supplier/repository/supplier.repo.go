package supplier_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	supplierdomain "shop_erp_mono/domain/warehouse_management/supplier"
	"shop_erp_mono/repository"
)

type supplierRepository struct {
	database           *mongo.Database
	supplierCollection string
}

func NewSupplierRepository(database *mongo.Database, supplierCollection string) supplierdomain.ISupplierRepository {
	return &supplierRepository{database: database, supplierCollection: supplierCollection}
}

func (s *supplierRepository) CreateOne(ctx context.Context, supplier supplierdomain.Supplier) error {
	supplierCollection := s.database.Collection(s.supplierCollection)

	_, err := supplierCollection.InsertOne(ctx, supplier)
	if err != nil {
		return err
	}

	return nil
}

func (s *supplierRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*supplierdomain.Supplier, error) {
	supplierCollection := s.database.Collection(s.supplierCollection)

	filter := bson.M{"_id": id}
	var supplier supplierdomain.Supplier
	err := supplierCollection.FindOne(ctx, filter).Decode(&supplier)
	if err != nil {
		return nil, err
	}

	return &supplier, nil
}

func (s *supplierRepository) GetByName(ctx context.Context, name string) (*supplierdomain.Supplier, error) {
	supplierCollection := s.database.Collection(s.supplierCollection)

	filter := bson.M{"name": name}
	var supplier supplierdomain.Supplier
	err := supplierCollection.FindOne(ctx, filter).Decode(&supplier)
	if err != nil {
		return nil, err
	}

	return &supplier, nil
}

func (s *supplierRepository) GetAll(ctx context.Context) ([]supplierdomain.Supplier, error) {
	supplierCollection := s.database.Collection(s.supplierCollection)

	filter := bson.M{}
	cursor, err := supplierCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var suppliers []supplierdomain.Supplier
	suppliers = make([]supplierdomain.Supplier, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var supplier supplierdomain.Supplier
		if err = cursor.Decode(&supplier); err != nil {
			return nil, err
		}

		suppliers = append(suppliers, supplier)
	}

	return suppliers, nil
}

func (s *supplierRepository) GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]supplierdomain.Supplier, error) {
	supplierCollection := s.database.Collection(s.supplierCollection)

	filter := bson.M{}
	cursor, err := repository.Paginate(ctx, supplierCollection, filter, pagination)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var suppliers []supplierdomain.Supplier
	suppliers = make([]supplierdomain.Supplier, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var supplier supplierdomain.Supplier
		if err = cursor.Decode(&supplier); err != nil {
			return nil, err
		}

		suppliers = append(suppliers, supplier)
	}

	return suppliers, nil
}

func (s *supplierRepository) UpdateOne(ctx context.Context, supplier *supplierdomain.Supplier) error {
	supplierCollection := s.database.Collection(s.supplierCollection)

	filter := bson.M{"_id": supplier.ID}
	update := bson.M{"$set": bson.M{
		"supplier_name":  supplier.SupplierName,
		"contact_person": supplier.ContactPerson,
		"phone_number":   supplier.PhoneNumber,
		"email":          supplier.Email,
		"address":        supplier.Address,
		"updated_at":     supplier.UpdatedAt,
	}}
	_, err := supplierCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (s *supplierRepository) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	supplierCollection := s.database.Collection(s.supplierCollection)

	filter := bson.M{"_id": id}
	_, err := supplierCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
