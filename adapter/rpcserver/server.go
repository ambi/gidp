package rpcserver

import (
	"database/sql"

	"github.com/ambi/go-web-app-patterns/adapter/sqlgateway"
	"github.com/ambi/go-web-app-patterns/api"
	"github.com/ambi/go-web-app-patterns/model"
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

// NewServer creates a new server for GPRC.
func NewServer(db *sql.DB) api.APIServer {
	return &server{
		tenantRepo: sqlgateway.NewTenantRepo(db),
		userRepo:   sqlgateway.NewUserRepo(db),
	}
}
