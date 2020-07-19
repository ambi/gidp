package main

import (
	"context"
	"database/sql"
	"log"
	"net"

	"github.com/ambi/go-web-app-patterns/adapter/sqlgateway"
	"github.com/ambi/go-web-app-patterns/api"
	"github.com/ambi/go-web-app-patterns/model"
	"github.com/ambi/go-web-app-patterns/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errTenantNotFound       = status.Error(codes.NotFound, "Tenant not found")
	errTenantOrUserNotFound = status.Error(codes.NotFound, "Tenant or user not found")
	errInternalServerError  = status.Error(codes.Internal, "Internal server error")
)

type server struct {
	tenantRepo model.TenantRepo
	userRepo   model.UserRepo
}

func newServer(db *sql.DB) api.APIServer {
	return &server{
		tenantRepo: sqlgateway.NewTenantRepo(db),
		userRepo:   sqlgateway.NewUserRepo(db),
	}
}

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

func newDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@/go_web_app_patterns")
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	db, err := newDB()
	if err != nil {
		log.Fatal(err)
	}

	srv := newServer(db)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	api.RegisterAPIServer(grpcServer, srv)
	grpcServer.Serve(lis)
}
