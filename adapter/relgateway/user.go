package relgateway

import (
	"context"

	"github.com/ambi/gidp/model"
	"github.com/go-rel/rel"
)

type userGateway struct {
	repo rel.Repository
}

// NewUserRepo creates a new user repository (SQL DB).
func NewUserRepo(repo rel.Repository) model.UserRepo {
	return &userGateway{repo: repo}
}

func (gw *userGateway) List(tenantID string) ([]*model.User, error) {
	var users []model.User

	ctx := context.Background()

	err := gw.repo.FindAll(ctx, &users, rel.Eq("tenant_id", tenantID))
	if err != nil {
		return nil, err
	}

	// []model.User -> []*model.User
	result := make([]*model.User, len(users))
	for i := range users {
		result[i] = &users[i]
	}
	return result, nil
}

func (gw *userGateway) Get(tenantID, userID string) (*model.User, error) {
	var user model.User

	ctx := context.Background()

	err := gw.repo.Find(ctx, &user, rel.Eq("tenant_id", tenantID).AndEq("id", userID))
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (gw *userGateway) Create(user *model.User) error {
	user.ID = model.NewUUID()

	ctx := context.Background()

	err := gw.repo.Insert(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (gw *userGateway) Update(user *model.User) error {
	ctx := context.Background()

	err := gw.repo.Update(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (gw *userGateway) Delete(user *model.User) error {
	ctx := context.Background()

	err := gw.repo.Delete(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
