package unit_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	employeesdomain "shop_erp_mono/domain/human_resource_management/employees"
	"shop_erp_mono/infrastructor"
	employeerepository "shop_erp_mono/repository/human_resource_management/employee/repository"
	"testing"
)

func TestDeleteOneEmployee(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	employeeName := "Hoài Phong"
	ur := employeerepository.NewEmployeeRepository(database, staff, department, role, salary)
	employeeData, err := ur.GetOneByName(context.Background(), employeeName)
	if err != nil {
		assert.Error(t, err)
	}

	mockNilEmployee := employeesdomain.Employee{}
	t.Run("success", func(t *testing.T) {
		err = ur.DeleteOne(context.Background(), employeeData.Employee.ID.Hex())
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		// Trying to insert an empty user, expecting an error
		err = ur.DeleteOne(context.Background(), mockNilEmployee.ID.Hex())
		assert.Error(t, err)
	})
}
