package unit_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	salarydomain "shop_erp_mono/domain/human_resource_management/salary"
	"shop_erp_mono/infrastructor"
	salaryrepository "shop_erp_mono/repository/human_resource_management/salary/repository"
	"testing"
)

func TestCreateOneSalary(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

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
		ur := salaryrepository.NewSalaryRepository(database, salary, role)
		err := ur.CreateOne(context.Background(), mockSalary)
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		ur := salaryrepository.NewSalaryRepository(database, salary, role)
		// Trying to insert an empty user, expecting an error
		err := ur.CreateOne(context.Background(), mockEmptySalary)
		assert.Error(t, err)
	})
}
