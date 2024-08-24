package unit_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"shop_erp_mono/infrastructor"
	salaryrepository "shop_erp_mono/repository/human_resource_management/salary/repository"
	"testing"
)

func TestGetAllSalary(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	t.Run("success", func(t *testing.T) {
		ur := salaryrepository.NewSalaryRepository(database, salary, role)

		_, err := ur.GetAll(context.Background())
		assert.Nil(t, err)
	})
}
