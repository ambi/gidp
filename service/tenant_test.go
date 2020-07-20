package service

import (
	"testing"

	"github.com/ambi/go-web-app-patterns/adapter/memgateway"
	"github.com/ambi/go-web-app-patterns/model"
)

var (
	tenantRepo model.TenantRepo
	tenant1    *model.Tenant
	tenant2    *model.Tenant
)

func setup() {
	tenantRepo = memgateway.NewTenantRepo()
	tenant1 = &model.Tenant{Status: "active"}
	tenant2 = &model.Tenant{Status: "active"}
	_ = tenantRepo.Create(tenant1)
	_ = tenantRepo.Create(tenant2)
}

func TestListTenants(t *testing.T) {
	setup()

	tenants, err := ListTenants(tenantRepo)

	if err != nil {
		t.Errorf("ListTenants() should not return error, but got %v", err)
	}
	if len(tenants) != 2 {
		t.Errorf("ListTenants() should return 2 tenants, but got %d tenants", len(tenants))
	}
	for _, tenant := range tenants {
		if tenant.ID != tenant1.ID && tenant.ID != tenant2.ID {
			t.Errorf("ListTenants() should return saved tenants, but got %v", tenant)
		}
	}
}

func TestGetTenant(t *testing.T) {
	setup()

	tenant, err := GetTenant(tenantRepo, tenant1.ID)
	if err != nil {
		t.Errorf("GetTenant() should not return error, but got %v", err)
	}
	if tenant.ID != tenant1.ID {
		t.Errorf("GetTenant() should return id=%s tenant, but got id=%s tenant", tenant1.ID, tenant.ID)
	}

	tenant, err = GetTenant(tenantRepo, "invalid")
	if err != model.ErrEntityNotFound {
		t.Errorf("GetTenant() should return error (ErrEntityNotFound), but got %v", err)
	}
	if tenant != nil {
		t.Errorf("GetTenant() should not return tenant, but got %v", tenant)
	}
}

func TestCreateTenant(t *testing.T) {
	setup()

	tenant, err := CreateTenant(tenantRepo)
	if err != nil {
		t.Errorf("CreateTenant() should not return error, but got %v", err)
	}
	if tenant == nil {
		t.Errorf("CreateTenant() should return tenant, but got nil")
	}

	tenants, _ := tenantRepo.List()
	if len(tenants) != 3 {
		t.Errorf("CreateTenant() should create a new tenant, but the size of tenant is %d", len(tenants))
	}
}

func TestDeleteTenant(t *testing.T) {
	setup()

	err := DeleteTenant(tenantRepo, tenant1.ID)
	if err != nil {
		t.Errorf("DeleteTenant() should not return error, but got %v", err)
	}
	tenants, _ := tenantRepo.List()
	if len(tenants) != 1 {
		t.Errorf("DeleteTenant() should delete a tenant, but the size of tenant is %d", len(tenants))
	}

	err = DeleteTenant(tenantRepo, "invalid")
	if err != model.ErrEntityNotFound {
		t.Errorf("DeleteTenant() should return error (ErrEntityNotFound), but got %v", err)
	}
}
