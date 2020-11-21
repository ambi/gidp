package sqlgateway

import (
	"context"
	"database/sql"

	"github.com/ambi/gidp/model"
)

const (
	sqlTenantList   = "SELECT id, status FROM tenants"
	sqlTenantGet    = "SELECT id, status FROM tenants WHERE id=?"
	sqlTenantCreate = "INSERT INTO tenants (id) VALUES (?)"
	sqlTenantDelete = "DELETE FROM tenants WHERE id=?"
)

type tenantGateway struct {
	db *sql.DB
}

// NewTenantRepo creates a new tenant repository (SQL DB).
func NewTenantRepo(db *sql.DB) model.TenantRepo {
	return &tenantGateway{db: db}
}

func (gw *tenantGateway) List() ([]*model.Tenant, error) {
	ctx := context.Background()
	rows, err := gw.db.QueryContext(ctx, sqlTenantList)
	if err != nil {
		return nil, err
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	defer rows.Close()

	tenants := make([]*model.Tenant, 0)
	for rows.Next() {
		tenant := &model.Tenant{}
		if err := rows.Scan(&tenant.ID, &tenant.Status); err != nil {
			return nil, err
		}
		tenants = append(tenants, tenant)
	}
	// TODO: check rows.Err()
	return tenants, nil
}

func (gw *tenantGateway) Get(id string) (*model.Tenant, error) {
	ctx := context.Background()
	row := gw.db.QueryRowContext(ctx, sqlTenantGet, id)

	tenant := &model.Tenant{}
	err := row.Scan(&tenant.ID, &tenant.Status)

	if err == sql.ErrNoRows {
		return nil, model.ErrEntityNotFound
	}
	if err != nil {
		return nil, err
	}
	return tenant, nil
}

func (gw *tenantGateway) Create(tenant *model.Tenant) error {
	ctx := context.Background()
	id := model.NewUUID()
	status := "active"
	_, err := gw.db.ExecContext(ctx, sqlTenantCreate, id, status)
	if err != nil {
		return err
	}

	tenant.ID = id

	return nil
}

func (gw *tenantGateway) Delete(tenant *model.Tenant) error {
	ctx := context.Background()
	result, err := gw.db.ExecContext(ctx, sqlTenantDelete, tenant.ID)
	if err != nil {
		return err
	}

	// Note: Not every database or database driver may support this (RowsAffected).
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return model.ErrEntityNotFound
	}
	return nil
}
