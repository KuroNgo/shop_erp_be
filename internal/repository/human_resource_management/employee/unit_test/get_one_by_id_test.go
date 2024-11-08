package unit_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"shop_erp_mono/infrastructor"
	employeerepository "shop_erp_mono/repository/human_resource_management/employee/repository"
	"testing"
)

func TestGetOneByEmail(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	ur := employeerepository.NewEmployeeRepository(database, staff)
	email := "admin@admin.com"

	t.Run("success", func(t *testing.T) {
		_, err := ur.GetOneByEmail(context.Background(), email)
		assert.Nil(t, err)
	})

	t.Run("invalid ID format", func(t *testing.T) {
		_, err := ur.GetOneByEmail(context.Background(), "invalidID")

		assert.Error(t, err)
		assert.Equal(t, "invalid employee ID format", err.Error())
	})

	t.Run("non-existing ID", func(t *testing.T) {
		nonExistingID := ""
		_, err := ur.GetOneByEmail(context.Background(), nonExistingID)
		assert.Error(t, err)
		assert.Equal(t, "error finding employee's information in the database", err.Error())
	})
}
