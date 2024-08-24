package unit_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	salarydomain "shop_erp_mono/domain/human_resource_management/salary"
	"shop_erp_mono/infrastructor"
	salaryrepository "shop_erp_mono/repository/human_resource_management/salary/repository"
	"testing"
)

func TestUpdateOneRole(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	ur := salaryrepository.NewSalaryRepository(database, salary, role)

	mockSalary := &salarydomain.Input{
		Role:         "admin",
		UnitCurrency: "dollar",
		BaseSalary:   1500.00,
		Bonus:        200.00,
		Deductions:   100.00,
		NetSalary:    1600.00,
	}
	mockEmptySalary := &salarydomain.Input{}

	t.Run("success", func(t *testing.T) {
		err := ur.UpdateOne(context.Background(), mockSalary)
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		// Trying to insert an empty user, expecting an error
		err := ur.UpdateOne(context.Background(), mockEmptySalary)
		assert.Error(t, err)
	})
}
