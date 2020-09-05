package service

import (
	"testing"

	"github.com/ambi/go-web-app-patterns/adapter/memgateway"
	"github.com/ambi/go-web-app-patterns/model"
	"github.com/stretchr/testify/suite"
)

type UserTestSuite struct {
	suite.Suite
	users
}

type users struct {
	tenants
	userRepo model.UserRepo
	user1    *model.User
	user2    *model.User
	user3    *model.User
}

func (t *users) setup() {
	t.tenants.setup()

	t.userRepo = memgateway.NewUserRepo()
	t.user1 = &model.User{TenantID: t.tenant1.ID, DisplayName: "user1@example.com"}
	t.user2 = &model.User{TenantID: t.tenant1.ID, DisplayName: "user2@example.com"}
	t.user3 = &model.User{TenantID: t.tenant2.ID, DisplayName: "user3@example.com"}
	_ = t.userRepo.Create(t.user1)
	_ = t.userRepo.Create(t.user2)
	_ = t.userRepo.Create(t.user3)
}

func (t *UserTestSuite) SetupTest() {
	t.tenants.setup()
	t.users.setup()
}

func (t *UserTestSuite) TestListUsers() {
	users, err := ListUsers(t.tenantRepo, t.userRepo, t.tenant1.ID)

	t.Nil(err)
	t.Len(users, 2)
	for _, user := range users {
		t.Contains([]string{t.user1.ID, t.user2.ID}, user.ID)
	}
}

func (t *UserTestSuite) TestGetUser() {
	user, err := GetUser(t.tenantRepo, t.userRepo, t.user1.TenantID, t.user1.ID)

	t.Nil(err)
	t.Equal(t.user1.ID, user.ID)

	user, err = GetUser(t.tenantRepo, t.userRepo, t.tenant2.ID, t.user1.ID)

	t.Equal(model.ErrEntityNotFound, err)
	t.Nil(user)

	user, err = GetUser(t.tenantRepo, t.userRepo, "invalid", t.user1.ID)

	t.Equal(model.ErrEntityNotFound, err)
	t.Nil(user)

	user, err = GetUser(t.tenantRepo, t.userRepo, t.user1.TenantID, "invalid")

	t.Equal(model.ErrEntityNotFound, err)
	t.Nil(user)
}

func (t *UserTestSuite) TestCreateUser() {
	user := &model.User{DisplayName: "user4@example.com"}

	err := CreateUser(t.tenantRepo, t.userRepo, t.tenant1.ID, user)

	t.Nil(err)
	t.NotNil(user)

	users, _ := t.userRepo.List(t.tenant1.ID)

	t.Len(users, 3)
}

func (t *UserTestSuite) TestDeleteUser() {
	err := DeleteUser(t.tenantRepo, t.userRepo, t.user1.TenantID, t.user1.ID)
	t.Nil(err)

	users, _ := t.userRepo.List(t.user1.TenantID)
	t.Len(users, 1)

	err = DeleteUser(t.tenantRepo, t.userRepo, t.tenant2.ID, t.user1.ID)
	t.Equal(model.ErrEntityNotFound, err)

	err = DeleteUser(t.tenantRepo, t.userRepo, "invalid", t.user1.ID)
	t.Equal(model.ErrEntityNotFound, err)

	err = DeleteUser(t.tenantRepo, t.userRepo, t.user1.TenantID, "invalid")
	t.Equal(model.ErrEntityNotFound, err)
}

func TestUser(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
