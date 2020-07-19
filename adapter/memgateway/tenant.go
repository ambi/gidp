package memgateway

import "github.com/ambi/go-web-app-patterns/model"

type tenantGateway struct {
	tenants []*model.Tenant
}

// NewTenantRepo creates a new tenant repository.
func NewTenantRepo() model.TenantRepo {
	return &tenantGateway{tenants: make([]*model.Tenant, 0)}
}

func (gw *tenantGateway) List() ([]*model.Tenant, error) {
	return gw.tenants, nil
}

func (gw *tenantGateway) Get(id string) (*model.Tenant, error) {
	for _, tenant := range gw.tenants {
		if tenant.ID == id {
			return tenant, nil
		}
	}
	return nil, model.ErrEntityNotFound
}

func (gw *tenantGateway) Create(tenant *model.Tenant) error {
	// TODO: validation check
	tenant.ID = model.NewUUID()
	gw.tenants = append(gw.tenants, tenant)
	return nil
}

func (gw *tenantGateway) Delete(tenant *model.Tenant) error {
	for i, t := range gw.tenants {
		if t.ID == tenant.ID {
			gw.tenants = append(gw.tenants[:i], gw.tenants[i+1:]...)
			return nil
		}
	}
	return model.ErrEntityNotFound
}
