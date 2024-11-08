package unit_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"shop_erp_mono/infrastructor"
	salary_repository "shop_erp_mono/repository/human_resource_management/salary/repository"
	"testing"
)

func TestGetSalaryByName(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	position := "admin"
	ur := salary_repository.NewSalaryRepository(database, salary, role)
	mockEmptyPosition := ""

	t.Run("success", func(t *testing.T) {
		_, err := ur.GetOneByRole(context.Background(), position)
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		// Trying to insert an empty user, expecting an error
		_, err := ur.GetOneByRole(context.Background(), mockEmptyPosition)
		assert.Error(t, err)
	})
}
