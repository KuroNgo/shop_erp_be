package unit_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	roledomain "shop_erp_mono/domain/human_resource_management/role"
	"shop_erp_mono/infrastructor"
	role_repository "shop_erp_mono/repository/human_resource_management/role/repository"
	"testing"
)

func TestDeleteOneRole(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	positionData := &roledomain.Role{
		Title: "admin",
	}
	ro := role_repository.NewRoleRepository(database, role)
	position, err := ro.GetByTitleRole(context.Background(), positionData.Title)
	if err != nil {
		assert.Error(t, err)
	}
	mockEmptyRole := &roledomain.Role{}

	t.Run("success", func(t *testing.T) {
		ur := role_repository.NewRoleRepository(database, role)
		err = ur.DeleteOneRole(context.Background(), position.Role.ID.Hex())
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		ur := role_repository.NewRoleRepository(database, role)

		// Trying to insert an empty user, expecting an error
		err = ur.DeleteOneRole(context.Background(), mockEmptyRole.ID.Hex())
		assert.Error(t, err)
	})
}
