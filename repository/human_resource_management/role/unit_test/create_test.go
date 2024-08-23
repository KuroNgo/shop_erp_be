package unit_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	roledomain "shop_erp_mono/domain/human_resource_management/role"
	"shop_erp_mono/infrastructor"
	rolerepository "shop_erp_mono/repository/human_resource_management/role/repository"
	"testing"
	"time"
)

func TestCreateOneRole(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	mockRole := &roledomain.Role{
		ID:          primitive.NewObjectID(),
		Title:       "admin",
		Description: "abc",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	mockEmptyRole := &roledomain.Role{}

	t.Run("success", func(t *testing.T) {
		ur := rolerepository.NewRoleRepository(database, role)
		err := ur.CreateOneRole(context.Background(), mockRole)
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		ur := rolerepository.NewRoleRepository(database, role)

		// Trying to insert an empty user, expecting an error
		err := ur.CreateOneRole(context.Background(), mockEmptyRole)
		assert.Error(t, err)
	})
}
