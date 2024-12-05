package company_repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	companydomain "shop_erp_mono/internal/domain/company"
	"shop_erp_mono/internal/repository"
)

type companyRepository struct {
	companyCollection string
	database          *mongo.Database
}

func NewCompanyRepository(companyCollection string, database *mongo.Database) companydomain.ICompanyRepository {
	return &companyRepository{database: database, companyCollection: companyCollection}
}

func (c *companyRepository) CreateOne(ctx context.Context, company *companydomain.Company) error {
	companyCollection := c.database.Collection(c.companyCollection)

	_, err := companyCollection.InsertOne(ctx, company)
	if err != nil {
		return err
	}

	return nil
}

func (c *companyRepository) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	companyCollection := c.database.Collection(c.companyCollection)

	filter := bson.M{"_id": id}
	_, err := companyCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (c *companyRepository) DeleteSoft(ctx context.Context, id primitive.ObjectID) error {
	companyCollection := c.database.Collection(c.companyCollection)

	filter := bson.M{"_id": id}
	update := bson.D{{"status", "Inactive"}}
	_, err := companyCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (c *companyRepository) UpdateOne(ctx context.Context, id primitive.ObjectID, company *companydomain.Company) error {
	companyCollection := c.database.Collection(c.companyCollection)

	filter := bson.M{"_id": id}
	update := bson.M{
		"tax_id":         company.TaxID,
		"represent":      company.Represent,
		"name":           company.Name,
		"description":    company.Description,
		"address":        company.Address,
		"verify":         false,
		"status":         "inactive",
		"payment_status": "free",
		"level_payment":  0,
	}
	_, err := companyCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (c *companyRepository) UpdateStatus(ctx context.Context, id primitive.ObjectID, status string) error {
	companyCollection := c.database.Collection(c.companyCollection)

	filter := bson.M{"_id": id}
	update := bson.M{
		"verify":         true,
		"status":         "active",
		"payment_status": "free",
		"level_payment":  0,
	}
	_, err := companyCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (c *companyRepository) GetByID(ctx context.Context, id primitive.ObjectID) (companydomain.Company, error) {
	companyCollection := c.database.Collection(c.companyCollection)

	filter := bson.M{"_id": id}
	var company companydomain.Company
	err := companyCollection.FindOne(ctx, filter).Decode(&company)
	if err != nil {
		return companydomain.Company{}, err
	}

	return company, nil
}

func (c *companyRepository) GetByName(ctx context.Context, name string) (companydomain.Company, error) {
	companyCollection := c.database.Collection(c.companyCollection)

	filter := bson.M{"name": name}
	var company companydomain.Company
	err := companyCollection.FindOne(ctx, filter).Decode(&company)
	if err != nil {
		return companydomain.Company{}, err
	}

	return company, nil
}

func (c *companyRepository) GetAll(ctx context.Context) ([]companydomain.Company, error) {
	companyCollection := c.database.Collection(c.companyCollection)

	filter := bson.M{}
	cursor, err := companyCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	var companies []companydomain.Company
	companies = make([]companydomain.Company, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var company companydomain.Company
		if err = cursor.Decode(&company); err != nil {
			return nil, err
		}

		companies = append(companies, company)
	}

	// Check for any errors encountered during iteration
	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return companies, nil
}

func (c *companyRepository) GetAllWithPagination(ctx context.Context, pagination repository.Pagination) ([]companydomain.Company, error) {
	companyCollection := c.database.Collection(c.companyCollection)

	filter := bson.M{}
	cursor, err := repository.Paginate(ctx, companyCollection, filter, pagination)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err = cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	var companies []companydomain.Company
	companies = make([]companydomain.Company, 0, cursor.RemainingBatchLength())
	for cursor.Next(ctx) {
		var company companydomain.Company
		if err = cursor.Decode(&company); err != nil {
			return nil, err
		}

		companies = append(companies, company)
	}

	// Check for any errors encountered during iteration
	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return companies, nil
}
