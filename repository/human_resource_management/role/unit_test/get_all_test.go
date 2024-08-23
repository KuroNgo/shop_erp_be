package unit_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"shop_erp_mono/infrastructor"
	role_repository "shop_erp_mono/repository/human_resource_management/role/repository"
	"testing"
)

func TestGetAllRole(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	t.Run("success", func(t *testing.T) {
		ur := role_repository.NewRoleRepository(database, role)
		_, err := ur.GetAllRole(context.Background())
		assert.Nil(t, err)
	})
}
