package unit_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	salarydomain "shop_erp_mono/domain/human_resource_management/salary"
	"shop_erp_mono/infrastructor"
	salary_repository "shop_erp_mono/repository/human_resource_management/salary/repository"
	"testing"
)

func TestGetSalaryByID(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	position := "admin"
	ur := salary_repository.NewSalaryRepository(database, salary, role)
	salaryData, err := ur.GetOneByRole(context.Background(), position)
	if err != nil {
		assert.Error(t, err)
	}

	mockEmptySalary := &salarydomain.Salary{}

	t.Run("success", func(t *testing.T) {
		_, err = ur.GetOneByID(context.Background(), salaryData.Salary.ID.Hex())
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		// Trying to insert an empty user, expecting an error
		_, err = ur.GetOneByID(context.Background(), mockEmptySalary.ID.Hex())
		assert.Error(t, err)
	})
}
