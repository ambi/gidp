package service

import (
	"testing"

	"github.com/ambi/gidp/adapter/memgateway"
	"github.com/ambi/gidp/model"
	"github.com/stretchr/testify/assert"
)

type tenantFixture struct {
	tenantRepo model.TenantRepo
	tenant1    *model.Tenant
	tenant2    *model.Tenant
}

func fixTenant(f *tenantFixture) {
	f.tenantRepo = memgateway.NewTenantRepo()
	f.tenant1 = &model.Tenant{Status: "active"}
	f.tenant2 = &model.Tenant{Status: "active"}
	_ = f.tenantRepo.Create(f.tenant1)
	_ = f.tenantRepo.Create(f.tenant2)
}

func TestTenant(t *testing.T) {
	var f tenantFixture

	t.Run("ListTenants", func(t *testing.T) {
		fixTenant(&f)

		tenants, err := ListTenants(f.tenantRepo)

		assert.Nil(t, err)
		assert.Len(t, tenants, 2)
		for _, tenant := range tenants {
			assert.Contains(t, []string{f.tenant1.ID, f.tenant2.ID}, tenant.ID)
		}
	})

	t.Run("GetTenant", func(t *testing.T) {
		fixTenant(&f)

		t.Run("tenant ID is valid", func(t *testing.T) {
			tenant, err := GetTenant(f.tenantRepo, f.tenant1.ID)

			assert.Nil(t, err)
			assert.Equal(t, tenant.ID, f.tenant1.ID)
		})

		t.Run("tenant ID is invalid", func(t *testing.T) {
			tenant, err := GetTenant(f.tenantRepo, "invalid")

			assert.Equal(t, model.ErrEntityNotFound, err)
			assert.Nil(t, tenant)
		})
	})

	t.Run("CreateTenant", func(t *testing.T) {
		fixTenant(&f)

		tenant, err := CreateTenant(f.tenantRepo)

		assert.Nil(t, err)
		assert.NotNil(t, tenant)
		tenants, _ := f.tenantRepo.List()
		assert.Len(t, tenants, 3)
	})

	t.Run("DeleteTenant", func(t *testing.T) {
		fixTenant(&f)

		t.Run("tenant ID is valid", func(t *testing.T) {
			err := DeleteTenant(f.tenantRepo, f.tenant1.ID)

			assert.Nil(t, err)
			tenants, _ := f.tenantRepo.List()
			assert.Len(t, tenants, 1)
		})

		t.Run("tenant ID is invalid", func(t *testing.T) {
			err := DeleteTenant(f.tenantRepo, "invalid")

			assert.Equal(t, model.ErrEntityNotFound, err)
		})
	})
}
