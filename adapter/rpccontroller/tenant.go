package rpccontroller

import (
	"context"

	"github.com/ambi/gidp/model"
	"github.com/ambi/gidp/service"
	"github.com/golang/protobuf/ptypes/empty"
)

func (srv *server) ListTenants(ctx context.Context, _ *empty.Empty) (*ListTenantsResponse, error) {
	tenants, err := service.ListTenants(srv.tenantRepo)
	if err != nil {
		return nil, errInternalServerError
	}

	res := &ListTenantsResponse{}
	res.Tenants = make([]*Tenant, len(tenants))
	for i, tenant := range tenants {
		res.Tenants[i] = EncodeTenant(tenant)
	}

	return res, nil
}

func (srv *server) GetTenant(ctx context.Context, req *GetTenantRequest) (*Tenant, error) {
	// TODO: validation
	id := req.TenantId

	tenant, err := service.GetTenant(srv.tenantRepo, id)
	if err == model.ErrEntityNotFound {
		return nil, errTenantNotFound
	}
	if err != nil {
		return nil, errInternalServerError
	}

	return EncodeTenant(tenant), nil
}

func (srv *server) CreateTenant(ctx context.Context, _ *empty.Empty) (*Tenant, error) {
	tenant, err := service.CreateTenant(srv.tenantRepo)
	if err != nil {
		return nil, errInternalServerError
	}

	return EncodeTenant(tenant), nil
}

func (srv *server) DeleteTenant(ctx context.Context, req *DeleteTenantRequest) (*empty.Empty, error) {
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
