package unit_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	roledomain "shop_erp_mono/domain/human_resource_management/role"
	"shop_erp_mono/infrastructor"
	rolerepository "shop_erp_mono/repository/human_resource_management/role/repository"
	"testing"
)

func TestGetByID(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	positionData := &roledomain.Role{
		Title: "admin",
	}
	ro := rolerepository.NewRoleRepository(database, role)
	position, err := ro.GetByTitleRole(context.Background(), positionData.Title)
	if err != nil {
		assert.Error(t, err)
	}
	mockPositionNil := &roledomain.Role{}

	t.Run("success", func(t *testing.T) {
		_, err = ro.GetByIDRole(context.Background(), position.Role.ID.Hex())
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		// Trying to insert an empty user, expecting an error
		_, err = ro.GetByIDRole(context.Background(), mockPositionNil.ID.Hex())
		assert.Error(t, err)
	})

}
