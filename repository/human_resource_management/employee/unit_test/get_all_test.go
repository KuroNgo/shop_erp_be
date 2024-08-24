package unit_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"shop_erp_mono/infrastructor"
	employeerepository "shop_erp_mono/repository/human_resource_management/employee/repository"
	"testing"
)

func TestGetAll(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	t.Run("success", func(t *testing.T) {
		ur := employeerepository.NewEmployeeRepository(database, staff, department, role, salary)

		_, err := ur.GetAll(context.Background())
		assert.Nil(t, err)
	})
}
