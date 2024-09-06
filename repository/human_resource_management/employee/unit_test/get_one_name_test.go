package unit_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"shop_erp_mono/infrastructor"
	employeerepository "shop_erp_mono/repository/human_resource_management/employee/repository"
	"testing"
)

func TestGetOneByID(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	ur := employeerepository.NewEmployeeRepository(database, staff)
	email := "admin@admin.com"
	employeeData, err := ur.GetOneByEmail(context.Background(), email)
	if err != nil {
		t.Fatalf("error setting up test data: %v", err)
	}

	t.Run("success", func(t *testing.T) {
		_, err = ur.GetOneByID(context.Background(), employeeData.ID)
		assert.Nil(t, err)
	})

	t.Run("non-existing ID", func(t *testing.T) {
		nonExistingID := primitive.NewObjectID()
		_, err := ur.GetOneByID(context.Background(), nonExistingID)
		assert.Error(t, err)
		assert.Equal(t, "error finding employee's information in the database", err.Error())
	})
}
