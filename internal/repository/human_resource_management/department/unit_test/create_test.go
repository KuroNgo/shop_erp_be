package unit_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	departmentsdomain "shop_erp_mono/internal/domain/human_resource_management/departments"
	infrastructor "shop_erp_mono/internal/infrastructor/mongo"
	departmentrepository "shop_erp_mono/internal/repository/human_resource_management/department/repository"
	"testing"
	"time"
)

func TestCreateOneDepartment(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	// Function to clear the employee collection before each test case
	clearDepartmentCollection := func() {
		err := database.Collection(Department).Drop(context.Background())
		if err != nil {
			t.Fatalf("Failed to clear employee collection: %v", err)
		}
	}

	clearEmployeeCollection := func() {
		err := database.Collection(Staff).Drop(context.Background())
		if err != nil {
			t.Fatalf("Failed to clear employee collection: %v", err)
		}
	}

	mockDepartment := &departmentsdomain.Department{
		ID:          primitive.NewObjectID(),
		Name:        "marketing",
		Description: "abc",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	mockEmptyDepartment := &departmentsdomain.Department{
		ID:          primitive.NewObjectID(),
		Name:        "", // Name is empty
		Description: "", // Description is empty
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		clearDepartmentCollection()
		clearEmployeeCollection()
		ur := departmentrepository.NewDepartmentRepository(database, Department)
		err := ur.CreateOne(context.Background(), mockDepartment)
		assert.Nil(t, err)
	})

	t.Run("error with empty fields", func(t *testing.T) {
		clearDepartmentCollection()
		clearEmployeeCollection()
		ur := departmentrepository.NewDepartmentRepository(database, Department)
		err := ur.CreateOne(context.Background(), mockEmptyDepartment)
		assert.Error(t, err)
	})
}
