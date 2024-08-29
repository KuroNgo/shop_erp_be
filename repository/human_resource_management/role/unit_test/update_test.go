package unit_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	roledomain "shop_erp_mono/domain/human_resource_management/role"
	"shop_erp_mono/infrastructor"
	role_repository "shop_erp_mono/repository/human_resource_management/role/repository"
	"testing"
)

func TestUpdateOneRole(t *testing.T) {
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

	mockRole := &roledomain.Input{
		Title:       "admin",
		Description: "abc",
	}
	mockEmptyRole := &roledomain.Input{}
	nilID := primitive.NilObjectID

	t.Run("success", func(t *testing.T) {
		ur := role_repository.NewRoleRepository(database, role)
		err = ur.UpdateOneRole(context.Background(), position.Role.ID.Hex(), mockRole)
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		ur := role_repository.NewRoleRepository(database, role)

		// Trying to insert an empty user, expecting an error
		err = ur.UpdateOneRole(context.Background(), nilID.Hex(), mockEmptyRole)
		assert.Error(t, err)
	})
}
