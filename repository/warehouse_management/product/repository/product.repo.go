package product_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	productdomain "shop_erp_mono/domain/warehouse_management/product"
	"time"
)

type productRepository struct {
	database          *mongo.Database
	productCollection string
}

func NewProductRepository(database *mongo.Database, productCollection string) productdomain.IProductRepository {
	return &productRepository{database: database, productCollection: productCollection}
}

func (p *productRepository) CreateProduct(ctx context.Context, product productdomain.Product) error {
	productCollection := p.database.Collection(p.productCollection)

	_, err := productCollection.InsertOne(ctx, product)
	if err != nil {
		return err
	}

	return nil
}

func (p *productRepository) UpdateProduct(ctx context.Context, id primitive.ObjectID, product productdomain.Product) error {
	productCollection := p.database.Collection(p.productCollection)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"product_name":      product.ProductName,
		"description":       product.Description,
		"price":             product.Price,
		"quantity_in_stock": product.QuantityInStock,
		"category_id":       product.CategoryID,
		"updated_at":        time.Now(),
	}}

	_, err := productCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (p *productRepository) GetProductByID(ctx context.Context, id primitive.ObjectID) (*productdomain.Product, error) {
	productCollection := p.database.Collection(p.productCollection)

	filter := bson.M{"_id": id}

	var product *productdomain.Product
	if err := productCollection.FindOne(ctx, filter).Decode(&product); err != nil {
		return nil, err
	}

	return product, nil
}

func (p *productRepository) GetProductByName(ctx context.Context, productName string) (*productdomain.Product, error) {
	productCollection := p.database.Collection(p.productCollection)

	filter := bson.M{"name": productName}
	var product *productdomain.Product
	if err := productCollection.FindOne(ctx, filter).Decode(&product); err != nil {
		return nil, err
	}

	return product, nil
}

func (p *productRepository) GetAllProducts(ctx context.Context) ([]productdomain.Product, error) {
	productCollection := p.database.Collection(p.productCollection)

	filter := bson.M{}
	cursor, err := productCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Next(ctx)

	var products []productdomain.Product
	products = make([]productdomain.Product, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var product productdomain.Product
		if err = cursor.Decode(&product); err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (p *productRepository) DeleteProduct(ctx context.Context, id primitive.ObjectID) error {
	productCollection := p.database.Collection(p.productCollection)

	filter := bson.M{"_id": id}
	_, err := productCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (p *productRepository) CountCategory(ctx context.Context, categoryID primitive.ObjectID) (int64, error) {
	productCollection := p.database.Collection(p.productCollection)

	filter := bson.M{"category_id": categoryID}
	count, err := productCollection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}

	return count, nil
}
