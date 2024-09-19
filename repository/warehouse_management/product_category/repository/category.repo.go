package category_repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"shop_erp_mono/domain/warehouse_management/product_category"
	"time"
)

type categoryRepository struct {
	database           *mongo.Database
	categoryCollection string
}

func NewCategoryRepository(database *mongo.Database, categoryCollection string) category_domain.ICategoryRepository {
	return &categoryRepository{database: database, categoryCollection: categoryCollection}
}

func (c *categoryRepository) CreateOne(ctx context.Context, category category_domain.Category) error {
	categoryCollection := c.database.Collection(c.categoryCollection)

	_, err := categoryCollection.InsertOne(ctx, category)
	if err != nil {
		return err
	}

	return nil
}

func (c *categoryRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*category_domain.Category, error) {
	categoryCollection := c.database.Collection(c.categoryCollection)

	filter := bson.M{"_id": id}
	var category *category_domain.Category
	err := categoryCollection.FindOne(ctx, filter).Decode(&category)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return category, nil
}

func (c *categoryRepository) GetByName(ctx context.Context, name string) (*category_domain.Category, error) {
	categoryCollection := c.database.Collection(c.categoryCollection)

	filter := bson.M{"name": name}
	var category *category_domain.Category
	err := categoryCollection.FindOne(ctx, filter).Decode(&category)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return category, nil
}

func (c *categoryRepository) UpdateOne(ctx context.Context, id primitive.ObjectID, input category_domain.Category) error {
	categoryCollection := c.database.Collection(c.categoryCollection)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"category_name": input.CategoryName,
		"description":   input.Description,
		"updated_at":    time.Now(),
	}}

	_, err := categoryCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (c *categoryRepository) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	categoryCollection := c.database.Collection(c.categoryCollection)

	filter := bson.M{"_id": id}
	_, err := categoryCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (c *categoryRepository) GetAll(ctx context.Context) ([]category_domain.Category, error) {
	categoryCollection := c.database.Collection(c.categoryCollection)

	filter := bson.M{}
	cursor, err := categoryCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var categories []category_domain.Category
	categories = make([]category_domain.Category, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var category category_domain.Category
		if err = cursor.Decode(&category); err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}
	return categories, nil
}
