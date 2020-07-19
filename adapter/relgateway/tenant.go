package relgateway

import (
	"context"

	"github.com/Fs02/rel"

	"github.com/ambi/go-web-app-patterns/model"
)

type tenantGateway struct {
	repo rel.Repository
}

// NewTenantRepo creates a new tenant repository (SQL DB).
func NewTenantRepo(repo rel.Repository) model.TenantRepo {
	return &tenantGateway{repo: repo}
}

func (gw *tenantGateway) List() ([]*model.Tenant, error) {
	var tenants []model.Tenant

	ctx := context.Background()

	err := gw.repo.FindAll(ctx, &tenants)
	if err != nil {
		return nil, err
	}

	// []model.Tenant -> []*model.Tenant
	result := make([]*model.Tenant, len(tenants))
	for i, tenant := range tenants {
		result[i] = &tenant
	}
	return result, nil
}

func (gw *tenantGateway) Get(id string) (*model.Tenant, error) {
	var tenant model.Tenant

	ctx := context.Background()

	err := gw.repo.Find(ctx, &tenant, rel.Eq("id", id))
	if err != nil {
		return nil, err
	}
	return &tenant, nil
}

func (gw *tenantGateway) Create(tenant *model.Tenant) error {
	tenant.ID = model.NewUUID()

	ctx := context.Background()

	err := gw.repo.Insert(ctx, &tenant)
	if err != nil {
		return err
	}

	return nil
}

func (gw *tenantGateway) Delete(tenant *model.Tenant) error {
	ctx := context.Background()

	err := gw.repo.Delete(ctx, tenant)
	if err != nil {
		return err
	}

	return nil
}
