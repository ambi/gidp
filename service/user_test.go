package service

import (
	"testing"

	"github.com/ambi/gidp/adapter/memgateway"
	"github.com/ambi/gidp/model"
	"github.com/stretchr/testify/assert"
)

type userFixture struct {
	tenantFixture
	userRepo model.UserRepo
	user1    *model.User
	user2    *model.User
	user3    *model.User
}

func fixUser(f *userFixture) {
	fixTenant(&f.tenantFixture)

	f.userRepo = memgateway.NewUserRepo()
	f.user1 = &model.User{TenantID: f.tenant1.ID, DisplayName: "user1@example.com"}
	f.user2 = &model.User{TenantID: f.tenant1.ID, DisplayName: "user2@example.com"}
	f.user3 = &model.User{TenantID: f.tenant2.ID, DisplayName: "user3@example.com"}
	_ = f.userRepo.Create(f.user1)
	_ = f.userRepo.Create(f.user2)
	_ = f.userRepo.Create(f.user3)
}

func TestUser(t *testing.T) {
	var f userFixture

	t.Run("ListUsers", func(t *testing.T) {
		fixUser(&f)

		users, err := ListUsers(f.tenantRepo, f.userRepo, f.tenant1.ID)

		assert.Nil(t, err)
		assert.Len(t, users, 2)
		for _, user := range users {
			assert.Contains(t, []string{f.user1.ID, f.user2.ID}, user.ID)
		}
	})

	t.Run("GetUser", func(t *testing.T) {
		fixUser(&f)

		t.Run("tenant and user ID are valid", func(t *testing.T) {
			user, err := GetUser(f.tenantRepo, f.userRepo, f.user1.TenantID, f.user1.ID)

			assert.Nil(t, err)
			assert.Equal(t, f.user1.ID, user.ID)
		})

		t.Run("user ID is in another tenant", func(t *testing.T) {
			user, err := GetUser(f.tenantRepo, f.userRepo, f.tenant2.ID, f.user1.ID)

			assert.Equal(t, model.ErrEntityNotFound, err)
			assert.Nil(t, user)
		})

		t.Run("tenant ID is invalid", func(t *testing.T) {
			user, err := GetUser(f.tenantRepo, f.userRepo, "invalid", f.user1.ID)

			assert.Equal(t, model.ErrEntityNotFound, err)
			assert.Nil(t, user)
		})

		t.Run("user ID is invalid", func(t *testing.T) {
			user, err := GetUser(f.tenantRepo, f.userRepo, f.user1.TenantID, "invalid")

			assert.Equal(t, model.ErrEntityNotFound, err)
			assert.Nil(t, user)
		})
	})

	t.Run("CreateUser", func(t *testing.T) {
		fixUser(&f)

		t.Run("tenant ID is valid", func(t *testing.T) {
			user := &model.User{DisplayName: "user4@example.com"}

			err := CreateUser(f.tenantRepo, f.userRepo, f.tenant1.ID, user)

			assert.Nil(t, err)
			assert.NotNil(t, user)
			users, _ := f.userRepo.List(f.tenant1.ID)
			assert.Len(t, users, 3)
		})

		t.Run("tenant ID is invalid", func(t *testing.T) {
			user := &model.User{DisplayName: "user4@example.com"}

			err := CreateUser(f.tenantRepo, f.userRepo, "invalid", user)

			assert.Equal(t, model.ErrEntityNotFound, err)
		})
	})

	t.Run("DeleteUser", func(t *testing.T) {
		fixUser(&f)

		t.Run("tenant and user ID are valid", func(t *testing.T) {
			err := DeleteUser(f.tenantRepo, f.userRepo, f.user1.TenantID, f.user1.ID)

			assert.Nil(t, err)
			users, _ := f.userRepo.List(f.user1.TenantID)
			assert.Len(t, users, 1)
		})

		t.Run("user ID is in another tenant", func(t *testing.T) {
			err := DeleteUser(f.tenantRepo, f.userRepo, f.tenant1.ID, f.user3.ID)
			assert.Equal(t, model.ErrEntityNotFound, err)
		})

		t.Run("tenant ID is invalid", func(t *testing.T) {
			err := DeleteUser(f.tenantRepo, f.userRepo, "invalid", f.user1.ID)

			assert.Equal(t, model.ErrEntityNotFound, err)
		})

		t.Run("user ID is invalid", func(t *testing.T) {
			err := DeleteUser(f.tenantRepo, f.userRepo, f.user1.TenantID, "invalid")

			assert.Equal(t, model.ErrEntityNotFound, err)
		})
	})
}
