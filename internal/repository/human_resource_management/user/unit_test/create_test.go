package unit_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	userdomain "shop_erp_mono/domain/human_resource_management/user"
	"shop_erp_mono/infrastructor"
	"shop_erp_mono/pkg/shared/password"
	user_repository "shop_erp_mono/repository/human_resource_management/user/repository"
	"testing"
	"time"
)

func TestCreateOneUser(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	mockUser := &userdomain.User{
		ID:           primitive.NewObjectID(),
		Username:     "unit_test",
		PasswordHash: "123",
		Email:        "admin@admin.com",
		Phone:        "0329245971",
		Role:         "user",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	mockEmptyUser := &userdomain.User{}
	mockUser.PasswordHash, _ = password.HashPassword(mockUser.PasswordHash)

	t.Run("success", func(t *testing.T) {
		ur := user_repository.NewUserRepository(database, individual)
		err := ur.Create(context.Background(), mockUser)
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		ur := user_repository.NewUserRepository(database, individual)

		// Trying to insert an empty user, expecting an error
		err := ur.Create(context.Background(), mockEmptyUser)
		assert.Error(t, err)
	})
}
