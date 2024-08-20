package user_repository

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	userdomain "shop_erp_mono/domain/human_resource_management/user"
	"shop_erp_mono/infrastructor"
	"shop_erp_mono/pkg/password"
	"testing"
	"time"
)

const (
	individual = "user"
)

func TestCreateOneUser(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	mockUser := &userdomain.User{
		ID:           primitive.NewObjectID(),
		Username:     "test",
		PasswordHash: "123",
		Email:        "admin@admin.com",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	mockEmptyUser := &userdomain.User{}
	mockUser.PasswordHash, _ = password.HashPassword(mockUser.PasswordHash)

	t.Run("success", func(t *testing.T) {
		ur := NewUserRepository(database, individual)
		_ = ur.Create(context.Background(), mockUser)
	})

	t.Run("error", func(t *testing.T) {
		ur := NewUserRepository(database, individual)

		// Trying to insert an empty user, expecting an error
		err := ur.Create(context.Background(), mockEmptyUser)
		assert.Error(t, err)
	})
}
