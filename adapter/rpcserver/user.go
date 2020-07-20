package rpcserver

import (
	"context"

	"github.com/ambi/go-web-app-patterns/api"
	"github.com/ambi/go-web-app-patterns/model"
	"github.com/ambi/go-web-app-patterns/service"
	"github.com/golang/protobuf/ptypes/empty"
)

func (srv *server) ListUsers(ctx context.Context, req *api.ListUsersRequest) (*api.ListUsersResponse, error) {
	// TODO: validation
	tenantID := req.TenantId

	users, err := service.ListUsers(srv.tenantRepo, srv.userRepo, tenantID)
	if err != nil {
		return nil, errInternalServerError
	}

	res := &api.ListUsersResponse{}
	res.Users = make([]*api.User, len(users))
	for i, user := range users {
		res.Users[i] = api.EncodeUser(user)
	}

	return res, nil
}

func (srv *server) GetUser(ctx context.Context, req *api.GetUserRequest) (*api.User, error) {
	// TODO: validation
	tenantID := req.TenantId
	id := req.UserId

	user, err := service.GetUser(srv.tenantRepo, srv.userRepo, tenantID, id)
	if err == model.ErrEntityNotFound {
		return nil, errTenantOrUserNotFound
	}
	if err != nil {
		return nil, errInternalServerError
	}

	return api.EncodeUser(user), nil
}

func (srv *server) CreateUser(ctx context.Context, req *api.CreateUserRequest) (*api.User, error) {
	// TODO: validation
	tenantID := req.TenantId
	displayName := req.User.DisplayName

	user := &model.User{
		DisplayName: displayName,
	}

	err := service.CreateUser(srv.tenantRepo, srv.userRepo, tenantID, user)
	if err != nil {
		return nil, errInternalServerError
	}

	return api.EncodeUser(user), nil
}

func (srv *server) UpdateUser(ctx context.Context, req *api.UpdateUserRequest) (*api.User, error) {
	// TODO: validation
	tenantID := req.TenantId
	id := req.User.Id
	displayName := req.User.DisplayName

	user := &model.User{ID: id, DisplayName: displayName}
	err := service.UpdateUser(srv.tenantRepo, srv.userRepo, tenantID, user)
	if err == model.ErrEntityNotFound {
		return nil, errTenantOrUserNotFound
	}
	if err != nil {
		return nil, errInternalServerError
	}

	return api.EncodeUser(user), nil
}

func (srv *server) DeleteUser(ctx context.Context, req *api.DeleteUserRequest) (*empty.Empty, error) {
	// TODO: validation
	tenantID := req.TenantId
	id := req.UserId

	err := service.DeleteUser(srv.tenantRepo, srv.userRepo, tenantID, id)
	if err == model.ErrEntityNotFound {
		return nil, errTenantOrUserNotFound
	}
	if err != nil {
		return nil, errInternalServerError
	}

	return &empty.Empty{}, nil
}
