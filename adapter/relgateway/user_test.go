package relgateway

import (
	"testing"

	"github.com/ambi/gidp/model"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/reltest"
	"github.com/stretchr/testify/assert"
)

func fixUsers() []*model.User {
	tenants := fixTenants()

	return []*model.User{
		{ID: "f2bf01e9-be5f-47a3-b8f4-6dba90d7d497", DisplayName: "x", TenantID: tenants[0].ID},
		{ID: "d2b02687-8d0d-4fae-8eb8-1771a3d26a78", DisplayName: "y", TenantID: tenants[0].ID},
		{ID: "535fa165-89ce-47f8-be5c-153d6f8a4323", DisplayName: "z", TenantID: tenants[1].ID},
	}
}

func setupUserGateway() (*reltest.Repository, model.UserRepo, []*model.User) {
	repo := reltest.New()
	gw := NewUserRepo(repo)
	users := fixUsers()
	return repo, gw, users
}

func TestUserGateway(t *testing.T) {
	t.Run("NewUserRepo", func(t *testing.T) {
		_, gw, _ := setupUserGateway()

		assert.NotNil(t, gw)
	})

	t.Run("List", func(t *testing.T) {
		repo, gw, users := setupUserGateway()
		relUsers := []model.User{*users[0], *users[1]}
		repo.ExpectFindAll(rel.Eq("tenant_id", users[0].TenantID)).Result(relUsers)

		result, err := gw.List(users[0].TenantID)

		assert.Nil(t, err)
		expected := make([]*model.User, 2)
		for i := 0; i < 2; i++ {
			expected[i] = users[i]
		}
		assert.Equal(t, expected, result)
		repo.AssertExpectations(t)
	})

	t.Run("Get", func(t *testing.T) {
		repo, gw, users := setupUserGateway()

		t.Run("user ID is valid", func(t *testing.T) {
			repo.ExpectFind(rel.Eq("tenant_id", users[0].TenantID).AndEq("id", users[0].ID)).Result(*users[0])

			result, err := gw.Get(users[0].TenantID, users[0].ID)

			assert.Nil(t, err)
			assert.Equal(t, users[0], result)
			repo.AssertExpectations(t)
		})

		t.Run("user ID is invalid", func(t *testing.T) {
			repo.ExpectFind(rel.Eq("tenant_id", users[0].TenantID).AndEq("id", users[0].ID)).NotFound()

			result, err := gw.Get(users[0].TenantID, users[0].ID)

			assert.NotNil(t, err)
			assert.Nil(t, result)
			repo.AssertExpectations(t)
		})
	})

	t.Run("Create", func(t *testing.T) {
		repo, gw, users := setupUserGateway()
		repo.ExpectInsert()

		user := model.User{
			DisplayName: "a",
			TenantID:    users[2].TenantID,
		}
		err := gw.Create(&user)

		assert.Nil(t, err)
		assert.Equal(t, "a", user.DisplayName)
		assert.Equal(t, users[2].TenantID, user.TenantID)
		assert.Regexp(t, "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$", user.ID)
		repo.AssertExpectations(t)
	})

	t.Run("Update", func(t *testing.T) {
		repo, gw, users := setupUserGateway()
		repo.ExpectUpdate()

		err := gw.Update(users[0])

		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Delete", func(t *testing.T) {
		repo, gw, users := setupUserGateway()
		repo.ExpectDelete()

		err := gw.Delete(users[0])

		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
}
