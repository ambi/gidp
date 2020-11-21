package relgateway

import (
	"testing"

	"github.com/ambi/gidp/model"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/reltest"
	"github.com/stretchr/testify/assert"
)

func fixTenants() []*model.Tenant {
	return []*model.Tenant{
		{ID: "edd2f7e8-21a0-4c56-a94d-35f3b6773f12", Status: "active"},
		{ID: "a463c6a2-86f9-44d6-9c36-024ec766b0e1", Status: "active"},
	}
}

func setupTenantGateway() (*reltest.Repository, model.TenantRepo, []*model.Tenant) {
	repo := reltest.New()
	gw := NewTenantRepo(repo)
	tenants := fixTenants()
	return repo, gw, tenants
}

func TestTenantGateway(t *testing.T) {
	t.Run("NewTenantRepo", func(t *testing.T) {
		_, gw, _ := setupTenantGateway()

		assert.NotNil(t, gw)
	})

	t.Run("List", func(t *testing.T) {
		repo, gw, tenants := setupTenantGateway()
		relTenants := []model.Tenant{*tenants[0], *tenants[1]}
		repo.ExpectFindAll().Result(relTenants)

		result, err := gw.List()

		assert.Nil(t, err)
		expected := make([]*model.Tenant, 2)
		for i := 0; i < len(tenants); i++ {
			expected[i] = tenants[i]
		}
		assert.Equal(t, expected, result)
		repo.AssertExpectations(t)
	})

	t.Run("Get", func(t *testing.T) {
		repo, gw, tenants := setupTenantGateway()

		t.Run("tenant ID is valid", func(t *testing.T) {
			repo.ExpectFind(rel.Eq("id", tenants[0].ID)).Result(*tenants[0])

			result, err := gw.Get(tenants[0].ID)

			assert.Nil(t, err)
			assert.Equal(t, tenants[0], result)
			repo.AssertExpectations(t)
		})

		t.Run("tenant ID is invalid", func(t *testing.T) {
			repo.ExpectFind(rel.Eq("id", tenants[0].ID)).NotFound()

			result, err := gw.Get(tenants[0].ID)

			assert.NotNil(t, err)
			assert.Nil(t, result)
			repo.AssertExpectations(t)
		})
	})

	t.Run("Create", func(t *testing.T) {
		repo, gw, _ := setupTenantGateway()
		repo.ExpectInsert()

		tenant := model.Tenant{
			Status: "inactive",
		}
		err := gw.Create(&tenant)

		assert.Nil(t, err)
		assert.Equal(t, "inactive", tenant.Status)
		assert.Regexp(t, "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$", tenant.ID)
		repo.AssertExpectations(t)
	})

	t.Run("Delete", func(t *testing.T) {
		repo, gw, tenants := setupTenantGateway()
		repo.ExpectDelete()

		err := gw.Delete(tenants[0])

		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
}
