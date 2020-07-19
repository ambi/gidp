package sqlgateway

import (
	"context"
	"database/sql"

	"github.com/ambi/go-web-app-patterns/model"
)

type userGateway struct {
	db *sql.DB
}

// NewUserRepo creates a new user repository (SQL DB).
func NewUserRepo(db *sql.DB) model.UserRepo {
	return &userGateway{db: db}
}

func (gw *userGateway) List(tenantID string) ([]*model.User, error) {
	const query = "SELECT id, display_name FROM users WHERE tenant_id=?"

	ctx := context.Background()
	rows, err := gw.db.QueryContext(ctx, query, tenantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*model.User, 0)
	for rows.Next() {
		user := &model.User{}
		if err := rows.Scan(&user.ID, &user.DisplayName); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	// TODO: check rows.Err()
	return users, nil
}

func (gw *userGateway) Get(tenantID, userID string) (*model.User, error) {
	const query = "SELECT id, display_name FROM users WHERE tenant_id=? AND id=?"

	ctx := context.Background()
	row := gw.db.QueryRowContext(ctx, query, tenantID, userID)

	user := &model.User{}
	err := row.Scan(&user.ID, &user.DisplayName)

	if err == sql.ErrNoRows {
		return nil, model.ErrEntityNotFound
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (gw *userGateway) Create(user *model.User) error {
	const query = "INSERT INTO users (id, display_name) VALUES (?, ?)"

	ctx := context.Background()
	id := model.NewUUID()
	_, err := gw.db.ExecContext(ctx, query, id, user.DisplayName)
	if err != nil {
		return err
	}
	user.ID = id

	return nil
}

func (gw *userGateway) Update(user *model.User) error {
	const query = "UPDATE users SET display_name=? WHERE id=?"

	ctx := context.Background()
	result, err := gw.db.ExecContext(ctx, query, user.DisplayName, user.ID)
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

func (gw *userGateway) Delete(user *model.User) error {
	const query = "DELETE FROM users WHERE id=?"

	ctx := context.Background()
	result, err := gw.db.ExecContext(ctx, query, user.ID)
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
