package department_seeder

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	departments_domain "shop_erp_mono/internal/domain/human_resource_management/departments"
	"sync"
	"time"
)

var departments = []departments_domain.Department{
	{
		ID:          primitive.NewObjectID(),
		Name:        "Human Resources",
		Description: "Responsible for hiring, training, and employee welfare.",
		Level:       1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          primitive.NewObjectID(),
		Name:        "Finance",
		Description: "Handles financial planning, accounting, and financial analysis.",
		Level:       1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          primitive.NewObjectID(),
		Name:        "Legal",
		Description: "Ensures company compliance with laws and manages legal issues.",
		Level:       1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          primitive.NewObjectID(),
		Name:        "IT",
		Description: "Responsible for maintaining and developing IT infrastructure.",
		Level:       1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          primitive.NewObjectID(),
		Name:        "Marketing",
		Description: "Develops and executes marketing strategies to promote the brand.",
		Level:       1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          primitive.NewObjectID(),
		Name:        "Sales",
		Description: "Generates revenue by selling products or services to customers.",
		Level:       1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          primitive.NewObjectID(),
		Name:        "Customer Service",
		Description: "Provides assistance and support to customers.",
		Level:       1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          primitive.NewObjectID(),
		Name:        "Research and Development",
		Description: "Innovates and develops new products or improves existing ones.",
		Level:       1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          primitive.NewObjectID(),
		Name:        "Quality Control",
		Description: "Ensures products or services meet quality standards.",
		Level:       1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          primitive.NewObjectID(),
		Name:        "Operations",
		Description: "Oversees the production and efficiency of business processes.",
		Level:       1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          primitive.NewObjectID(),
		Name:        "Supply Chain Management",
		Description: "Manages the logistics and supply of materials and products.",
		Level:       1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          primitive.NewObjectID(),
		Name:        "Project Management",
		Description: "Oversees and manages the execution of major projects.",
		Level:       1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
}

func SeedDepartment(ctx context.Context, client *mongo.Client) error {
	collectionDepartment := client.Database("shopERP").Collection("department")

	countDepartment, err := collectionDepartment.CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}

	errCh := make(chan error, 1)
	var wg sync.WaitGroup

	if countDepartment == 0 {
		for _, department := range departments {
			wg.Add(1)
			dept := department
			go func() {
				defer wg.Done()
				_, err = collectionDepartment.InsertOne(ctx, dept)
				if err != nil {
					select {
					case errCh <- err:
					default:
					}
				}
			}()
		}
	}

	wg.Wait()
	close(errCh)

	if err, ok := <-errCh; ok {
		return err
	}
	return nil
}
