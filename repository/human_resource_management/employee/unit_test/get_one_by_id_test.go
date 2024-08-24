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

	ur := employeerepository.NewEmployeeRepository(database, staff, department, role, salary)
	employeeName := "Hoài Phong"

	t.Run("success", func(t *testing.T) {
		_, err := ur.GetOneByName(context.Background(), employeeName)
		assert.Nil(t, err)
	})

	t.Run("invalid ID format", func(t *testing.T) {
		_, err := ur.GetOneByName(context.Background(), "invalidID")

		assert.Error(t, err)
		assert.Equal(t, "invalid employee ID format", err.Error())
	})

	t.Run("non-existing ID", func(t *testing.T) {
		nonExistingID := primitive.NewObjectID().Hex()
		_, err := ur.GetOneByName(context.Background(), nonExistingID)
		assert.Error(t, err)
		assert.Equal(t, "error finding employee's information in the database", err.Error())
	})
}
