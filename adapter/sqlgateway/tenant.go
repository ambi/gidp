package sqlgateway

import (
	"context"
	"database/sql"

	"github.com/ambi/go-web-app-patterns/model"
)

type tenantGateway struct {
	db *sql.DB
}

// NewTenantRepo creates a new tenant repository (SQL DB).
func NewTenantRepo(db *sql.DB) model.TenantRepo {
	return &tenantGateway{db: db}
}

func (gw *tenantGateway) List() ([]*model.Tenant, error) {
	const query = "SELECT id FROM tenants"

	ctx := context.Background()
	rows, err := gw.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tenants := make([]*model.Tenant, 0)
	for rows.Next() {
		tenant := &model.Tenant{}
		if err := rows.Scan(&tenant.ID); err != nil {
			return nil, err
		}
		tenants = append(tenants, tenant)
	}
	// TODO: check rows.Err()
	return tenants, nil
}

func (gw *tenantGateway) Get(id string) (*model.Tenant, error) {
	const query = "SELECT id FROM tenants WHERE id=?"

	ctx := context.Background()
	row := gw.db.QueryRowContext(ctx, query, id)

	tenant := &model.Tenant{}
	err := row.Scan(&tenant.ID)

	if err == sql.ErrNoRows {
		return nil, model.ErrEntityNotFound
	}
	if err != nil {
		return nil, err
	}
	return tenant, nil
}

func (gw *tenantGateway) Create(tenant *model.Tenant) error {
	const query = "INSERT INTO tenants (id) VALUES (?)"

	ctx := context.Background()
	id := model.NewUUID()
	_, err := gw.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	tenant.ID = id

	return nil
}

func (gw *tenantGateway) Delete(tenant *model.Tenant) error {
	const query = "DELETE FROM tenants WHERE id=?"

	ctx := context.Background()
	result, err := gw.db.ExecContext(ctx, query, tenant.ID)
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
