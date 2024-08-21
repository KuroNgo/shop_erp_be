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

func TestGetAllUser(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	t.Run("success", func(t *testing.T) {
		ur := NewUserRepository(database, individual)
		_, err := ur.FetchMany(context.Background())
		assert.Nil(t, err)
	})
}

func TestGetByIDUser(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	userData := userdomain.User{
		Email: "admin@admin.com",
	}
	ur := NewUserRepository(database, individual)
	userReq, err := ur.GetByEmail(context.Background(), userData.Email)

	mockUser := userdomain.User{
		ID: userReq.ID,
	}
	mockEmptyUser := userdomain.User{}

	t.Run("success", func(t *testing.T) {
		_, err = ur.GetByID(context.Background(), mockUser.ID.String())
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		// Trying to insert an empty user, expecting an error
		_, err = ur.GetByID(context.Background(), mockEmptyUser.ID.String())
		assert.Error(t, err)
	})
}

func TestGetByEmailUser(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	mockUser := &userdomain.User{
		Email: "admin@admin.com",
	}
	mockEmptyUser := &userdomain.User{}

	t.Run("success", func(t *testing.T) {
		ur := NewUserRepository(database, individual)
		_, err := ur.GetByEmail(context.Background(), mockUser.Email)
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		ur := NewUserRepository(database, individual)

		// Trying to insert an empty user, expecting an error
		_, err := ur.GetByEmail(context.Background(), mockEmptyUser.Email)
		assert.Error(t, err)
	})
}

func TestCreateOneUser(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	mockUser := &userdomain.User{
		ID:           primitive.NewObjectID(),
		Username:     "test",
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
		ur := NewUserRepository(database, individual)
		err := ur.Create(context.Background(), mockUser)
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		ur := NewUserRepository(database, individual)

		// Trying to insert an empty user, expecting an error
		err := ur.Create(context.Background(), mockEmptyUser)
		assert.Error(t, err)
	})
}

func TestUpdateUser(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	user := &userdomain.User{
		Email: "admin@admin.com",
	}
	ur := NewUserRepository(database, individual)
	userData, err := ur.GetByEmail(context.Background(), user.Email)
	assert.Nil(t, err)

	mockUser := &userdomain.UpdateUser{
		ID:        userData.ID,
		Username:  "test",
		AvatarURL: "",
		UpdatedAt: time.Now(),
	}
	mockEmptyUser := &userdomain.UpdateUser{}
	mockEmptyDataUser := &userdomain.UpdateUser{
		ID: userData.ID,
	}

	t.Run("success", func(t *testing.T) {
		err = ur.Update(context.Background(), mockUser)
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		// Trying to insert an empty user, expecting an error
		err = ur.Update(context.Background(), mockEmptyDataUser)
		assert.Error(t, err)
	})

	t.Run("error", func(t *testing.T) {
		// Trying to insert an empty user, expecting an error
		err = ur.Update(context.Background(), mockEmptyUser)
		assert.Error(t, err)
	})
}

func TestUpsertOneUser(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	mockUser := &userdomain.User{
		ID:           primitive.NewObjectID(),
		Username:     "test",
		PasswordHash: "123",
		Email:        "admin@admin.com",
		Phone:        "0329245971",
		Role:         "user",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	mockEmptyUser := &userdomain.User{}
	mockEmptyDataUser := &userdomain.User{
		Email: "admin@admin.com",
	}

	mockUser.PasswordHash, _ = password.HashPassword(mockUser.PasswordHash)

	t.Run("success", func(t *testing.T) {
		ur := NewUserRepository(database, individual)
		_, err := ur.UpsertOne(context.Background(), mockUser.Email, mockUser)
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		ur := NewUserRepository(database, individual)

		// Trying to insert an empty user, expecting an error
		_, err := ur.UpsertOne(context.Background(), mockEmptyUser.Email, mockEmptyUser)
		assert.Error(t, err)
	})

	t.Run("error", func(t *testing.T) {
		ur := NewUserRepository(database, individual)

		// Trying to insert an empty user, expecting an error
		_, err := ur.UpsertOne(context.Background(), mockEmptyDataUser.Email, mockEmptyDataUser)
		assert.Error(t, err)
	})
}

func TestUpdatePasswordUser(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	user := &userdomain.User{
		Email: "admin@admin.com",
	}
	ur := NewUserRepository(database, individual)
	userData, err := ur.GetByEmail(context.Background(), user.Email)
	assert.Nil(t, err)

	mockUser := &userdomain.User{
		ID:           userData.ID,
		PasswordHash: "abc",
	}
	mockEmptyPasswordUser := &userdomain.User{
		ID: userData.ID,
	}
	mockEmptyUser := &userdomain.User{}
	mockUser.PasswordHash, _ = password.HashPassword(mockUser.PasswordHash)

	t.Run("success", func(t *testing.T) {
		err = ur.UpdatePassword(context.Background(), mockUser)
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		// Trying to insert an empty user, expecting an error
		err = ur.UpdatePassword(context.Background(), mockEmptyPasswordUser)
		assert.Error(t, err)
	})

	t.Run("error", func(t *testing.T) {
		// Trying to insert an empty user, expecting an error
		err = ur.UpdatePassword(context.Background(), mockEmptyUser)
		assert.Error(t, err)
	})
}

func TestUpdateVerifiedUser(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	user := &userdomain.User{
		Email: "admin@admin.com",
	}
	ur := NewUserRepository(database, individual)
	userData, err := ur.GetByEmail(context.Background(), user.Email)
	assert.Nil(t, err)

	mockUser := &userdomain.User{
		ID:       userData.ID,
		Verified: true,
	}
	mockEmptyUser := &userdomain.User{
		ID: userData.ID,
	}

	t.Run("success", func(t *testing.T) {
		_, err = ur.UpdateVerify(context.Background(), mockUser)
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		// Trying to insert an empty user, expecting an error
		_, err = ur.UpdateVerify(context.Background(), mockEmptyUser)
		assert.Nil(t, err)
	})
}

func TestUpdateVerifiedForChangePasswordUser(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	user := &userdomain.User{
		Email: "admin@admin.com",
	}
	ur := NewUserRepository(database, individual)
	userData, err := ur.GetByEmail(context.Background(), user.Email)
	assert.Nil(t, err)

	mockUser := &userdomain.User{
		ID:       userData.ID,
		Verified: true,
	}
	mockEmptyUser := &userdomain.User{
		ID: userData.ID,
	}

	t.Run("success", func(t *testing.T) {
		_, err = ur.UpdateVerifyForChangePassword(context.Background(), mockUser)
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		// Trying to insert an empty user, expecting an error
		_, err = ur.UpdateVerifyForChangePassword(context.Background(), mockEmptyUser)
		assert.Nil(t, err)
	})
}

func TestDeleteUser(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	user := &userdomain.User{
		Email: "admin@admin.com",
	}
	ur := NewUserRepository(database, individual)
	userData, err := ur.GetByEmail(context.Background(), user.Email)
	assert.Nil(t, err)

	mockUser := &userdomain.User{
		ID: userData.ID,
	}
	mockEmptyUser := &userdomain.User{}

	t.Run("success", func(t *testing.T) {
		err = ur.DeleteOne(context.Background(), mockUser.ID.String())
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		// Trying to insert an empty user, expecting an error
		err = ur.DeleteOne(context.Background(), mockEmptyUser.ID.String())
		assert.Nil(t, err)
	})
}

func TestImageURLUser(t *testing.T) {
	client, database := infrastructor.SetupTestDatabase(t)
	defer infrastructor.TearDownTestDatabase(client, t)

	user := &userdomain.User{
		Email: "admin@admin.com",
	}
	ur := NewUserRepository(database, individual)
	userData, err := ur.GetByEmail(context.Background(), user.Email)
	assert.Nil(t, err)

	mockUser := &userdomain.User{
		ID:        userData.ID,
		AvatarURL: "abc",
	}
	mockEmptyUser := &userdomain.User{}

	t.Run("success", func(t *testing.T) {
		err = ur.UpdateImage(context.Background(), mockUser.ID.String(), mockUser.AvatarURL)
		assert.Nil(t, err)
	})

	t.Run("error", func(t *testing.T) {
		// Trying to insert an empty user, expecting an error
		err = ur.UpdateImage(context.Background(), mockEmptyUser.ID.String(), mockEmptyUser.AvatarURL)
		assert.Error(t, err)
	})
}
