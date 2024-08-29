package unit_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	userdomain "shop_erp_mono/domain/human_resource_management/user"
	"shop_erp_mono/infrastructor"
	userrepository "shop_erp_mono/repository/human_resource_management/user/repository"
	"testing"
)

func TestGetByIDUser(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	userData := userdomain.User{
		Email: "admin@admin.com",
	}

	ur := userrepository.NewUserRepository(database, individual)
	userReq, err := ur.GetByEmail(context.Background(), userData.Email)
	if err != nil {
		assert.Error(t, err)
	}

	mockUser := userdomain.User{
		ID: userReq.ID,
	}
	mockEmptyUser := userdomain.User{}

	t.Run("success", func(t *testing.T) {
		_, err = ur.GetByID(context.Background(), mockUser.ID.Hex())
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		// Trying to insert an empty user, expecting an error
		_, err = ur.GetByID(context.Background(), mockEmptyUser.ID.Hex())
		assert.Error(t, err)
	})
}
