package relgateway

import (
	"testing"

	"github.com/Fs02/rel"
	"github.com/Fs02/rel/reltest"
	"github.com/ambi/go-web-app-patterns/model"
	"github.com/stretchr/testify/assert"
)

type tenantTest struct {
	repo    *reltest.Repository
	gw      model.TenantRepo
	tenants []model.Tenant
}

func newTenantTest(_ *testing.T) *tenantTest {
	repo := reltest.New()
	gw := NewTenantRepo(repo)
	tenants := []model.Tenant{
		{ID: "ID1", Status: "active"},
		{ID: "ID2", Status: "active"},
	}

	return &tenantTest{
		repo:    repo,
		gw:      gw,
		tenants: tenants,
	}
}

func TestNewTenantRepo(t *testing.T) {
	s := newTenantTest(t)

	assert.NotNil(t, s.gw)
}

func TestTenantGateway_List(t *testing.T) {
	s := newTenantTest(t)
	expected := make([]*model.Tenant, len(s.tenants))
	for i := 0; i < len(s.tenants); i++ {
		expected[i] = &s.tenants[i]
	}

	s.repo.ExpectFindAll().Result(s.tenants)

	result, err := s.gw.List()

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
	s.repo.AssertExpectations(t)
}

func TestTenantGateway_Get(t *testing.T) {
	s := newTenantTest(t)

	s.repo.ExpectFind(rel.Eq("id", s.tenants[0].ID)).Result(s.tenants[0])
	s.repo.ExpectFind(rel.Eq("id", s.tenants[1].ID)).NotFound()

	result, err := s.gw.Get(s.tenants[0].ID)

	assert.Nil(t, err)
	assert.Equal(t, &s.tenants[0], result)

	result, err = s.gw.Get(s.tenants[1].ID)

	assert.NotNil(t, err)
	assert.Nil(t, result)

	s.repo.AssertExpectations(t)
}

func TestTenantGateway_Create(t *testing.T) {
	s := newTenantTest(t)
	tenant := s.tenants[0]

	s.repo.ExpectInsert()

	err := s.gw.Create(&tenant)

	assert.Nil(t, err)
	assert.Equal(t, s.tenants[0].Status, tenant.Status)
	assert.Regexp(t, "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$", tenant.ID)
	s.repo.AssertExpectations(t)
}
