package rpcserver

import (
	"context"

	"github.com/ambi/go-web-app-patterns/api"
	"github.com/ambi/go-web-app-patterns/model"
	"github.com/ambi/go-web-app-patterns/service"
	"github.com/golang/protobuf/ptypes/empty"
)

func (srv *server) ListTenants(ctx context.Context, _ *empty.Empty) (*api.ListTenantsResponse, error) {
	tenants, err := service.ListTenants(srv.tenantRepo)
	if err != nil {
		return nil, errInternalServerError
	}

	res := &api.ListTenantsResponse{}
	res.Tenants = make([]*api.Tenant, len(tenants))
	for i, tenant := range tenants {
		res.Tenants[i] = api.EncodeTenant(tenant)
	}

	return res, nil
}

func (srv *server) GetTenant(ctx context.Context, req *api.GetTenantRequest) (*api.Tenant, error) {
	// TODO: validation
	id := req.TenantId

	tenant, err := service.GetTenant(srv.tenantRepo, id)
	if err == model.ErrEntityNotFound {
		return nil, errTenantNotFound
	}
	if err != nil {
		return nil, errInternalServerError
	}

	return api.EncodeTenant(tenant), nil
}

func (srv *server) CreateTenant(ctx context.Context, _ *empty.Empty) (*api.Tenant, error) {
	tenant, err := service.CreateTenant(srv.tenantRepo)
	if err != nil {
		return nil, errInternalServerError
	}

	return api.EncodeTenant(tenant), nil
}

func (srv *server) DeleteTenant(ctx context.Context, req *api.DeleteTenantRequest) (*empty.Empty, error) {
	// TODO: validation
	id := req.TenantId

	err := service.DeleteTenant(srv.tenantRepo, id)
	if err == model.ErrEntityNotFound {
		return nil, errTenantNotFound
	}
	if err != nil {
		return nil, errInternalServerError
	}

	return &empty.Empty{}, nil
}
