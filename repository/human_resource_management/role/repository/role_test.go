package repository

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	roledomain "shop_erp_mono/domain/human_resource_management/role"
	"shop_erp_mono/infrastructor"
	"testing"
	"time"
)

const (
	role = "role"
)

func TestGetAllRole(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	t.Run("success", func(t *testing.T) {
		ur := NewRoleRepository(database, role)
		_, err := ur.GetAllRole(context.Background())
		assert.Nil(t, err)
	})
}

func TestGetByID(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	positionData := &roledomain.Role{
		Title: "admin",
	}
	ro := NewRoleRepository(database, role)
	position, err := ro.GetByTitleRole(context.Background(), positionData.Title)
	if err != nil {
		assert.Error(t, err)
	}
	mockPositionNil := &roledomain.Role{}

	t.Run("success", func(t *testing.T) {
		_, err = ro.GetByIDRole(context.Background(), position.ID.Hex())
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		// Trying to insert an empty user, expecting an error
		_, err = ro.GetByIDRole(context.Background(), mockPositionNil.ID.Hex())
		assert.Error(t, err)
	})

}

func TestCreateOneRole(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	mockRole := &roledomain.Role{
		ID:          primitive.NewObjectID(),
		Title:       "admin",
		Description: "",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	mockEmptyRole := &roledomain.Role{}

	t.Run("success", func(t *testing.T) {
		ur := NewRoleRepository(database, role)
		err := ur.CreateOneRole(context.Background(), mockRole)
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		ur := NewRoleRepository(database, role)

		// Trying to insert an empty user, expecting an error
		err := ur.CreateOneRole(context.Background(), mockEmptyRole)
		assert.Error(t, err)
	})
}

func TestUpdateOneRole(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	positionData := &roledomain.Role{
		Title: "admin",
	}
	ro := NewRoleRepository(database, role)
	position, err := ro.GetByTitleRole(context.Background(), positionData.Title)
	if err != nil {
		assert.Error(t, err)
	}

	mockRole := &roledomain.Role{
		ID:          position.ID,
		Title:       "admin",
		Description: "abc",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	mockEmptyRole := &roledomain.Role{}

	t.Run("success", func(t *testing.T) {
		ur := NewRoleRepository(database, role)
		err = ur.UpdateOneRole(context.Background(), mockRole)
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		ur := NewRoleRepository(database, role)

		// Trying to insert an empty user, expecting an error
		err = ur.UpdateOneRole(context.Background(), mockEmptyRole)
		assert.Error(t, err)
	})
}

func TestDeleteOneRole(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	positionData := &roledomain.Role{
		Title: "admin",
	}
	ro := NewRoleRepository(database, role)
	position, err := ro.GetByTitleRole(context.Background(), positionData.Title)
	if err != nil {
		assert.Error(t, err)
	}
	mockEmptyRole := &roledomain.Role{}

	t.Run("success", func(t *testing.T) {
		ur := NewRoleRepository(database, role)
		err = ur.DeleteOneRole(context.Background(), position.ID.Hex())
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		ur := NewRoleRepository(database, role)

		// Trying to insert an empty user, expecting an error
		err = ur.DeleteOneRole(context.Background(), mockEmptyRole.ID.Hex())
		assert.Error(t, err)
	})
}
