package service

import "github.com/ambi/gidp/model"

// ListTenants lists tenants.
func ListTenants(tenantRepo model.TenantRepo) ([]*model.Tenant, error) {
	return tenantRepo.List()
}

// GetTenant gets a tenant.
func GetTenant(tenantRepo model.TenantRepo, id string) (*model.Tenant, error) {
	return tenantRepo.Get(id)
}

// CreateTenant creates a new tenant.
func CreateTenant(tenantRepo model.TenantRepo) (*model.Tenant, error) {
	tenant := &model.Tenant{}
	err := tenantRepo.Create(tenant)
	if err != nil {
		return nil, err
	}

	return tenant, nil
}

// DeleteTenant deletes a tenant.
func DeleteTenant(tenantRepo model.TenantRepo, id string) error {
	tenant := &model.Tenant{ID: id}
	return tenantRepo.Delete(tenant)
}
