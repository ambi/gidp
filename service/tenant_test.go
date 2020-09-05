package service

import (
	"testing"

	"github.com/ambi/go-web-app-patterns/adapter/memgateway"
	"github.com/ambi/go-web-app-patterns/model"
	"github.com/stretchr/testify/suite"
)

type TenantTestSuite struct {
	suite.Suite
	tenants
}

type tenants struct {
	tenantRepo model.TenantRepo
	tenant1    *model.Tenant
	tenant2    *model.Tenant
}

func (t *tenants) setup() {
	t.tenantRepo = memgateway.NewTenantRepo()
	t.tenant1 = &model.Tenant{Status: "active"}
	t.tenant2 = &model.Tenant{Status: "active"}
	_ = t.tenantRepo.Create(t.tenant1)
	_ = t.tenantRepo.Create(t.tenant2)
}

func (t *TenantTestSuite) SetupTest() {
	t.tenants.setup()
}

func (t *TenantTestSuite) TestListTenants() {
	tenants, err := ListTenants(t.tenantRepo)

	t.Nil(err)
	t.Len(tenants, 2)
	for _, tenant := range tenants {
		t.Contains([]string{t.tenant1.ID, t.tenant2.ID}, tenant.ID)
	}
}

func (t *TenantTestSuite) TestGetTenant() {
	tenant, err := GetTenant(t.tenantRepo, t.tenant1.ID)

	t.Nil(err)
	t.Equal(tenant.ID, t.tenant1.ID)

	tenant, err = GetTenant(t.tenantRepo, "invalid")
	t.Equal(model.ErrEntityNotFound, err)
	t.Nil(tenant)
}

func (t *TenantTestSuite) TestCreateTenant() {
	tenant, err := CreateTenant(t.tenantRepo)

	t.Nil(err)
	t.NotNil(tenant)

	tenants, _ := t.tenantRepo.List()

	t.Len(tenants, 3)
}

func (t *TenantTestSuite) TestDeleteTenant() {
	err := DeleteTenant(t.tenantRepo, t.tenant1.ID)

	t.Nil(err)

	tenants, _ := t.tenantRepo.List()
	t.Len(tenants, 1)

	err = DeleteTenant(t.tenantRepo, "invalid")

	t.Equal(model.ErrEntityNotFound, err)
}

func TestTenant(t *testing.T) {
	suite.Run(t, new(TenantTestSuite))
}
