package rpccontroller

import (
	"database/sql"

	"github.com/ambi/gidp/adapter/sqlgateway"
	"github.com/ambi/gidp/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errTenantNotFound       = status.Error(codes.NotFound, "Tenant not found")
	errTenantOrUserNotFound = status.Error(codes.NotFound, "Tenant or user not found")
	errInternalServerError  = status.Error(codes.Internal, "Internal server error")
)

type server struct {
	UnimplementedAPIServer
	tenantRepo model.TenantRepo
	userRepo   model.UserRepo
}

// NewServer creates a new server for GPRC.
func NewServer(db *sql.DB) APIServer {
	return &server{
		tenantRepo: sqlgateway.NewTenantRepo(db),
		userRepo:   sqlgateway.NewUserRepo(db),
	}
}
