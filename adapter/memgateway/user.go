package memgateway

import (
	"github.com/ambi/gidp/model"
)

type userGateway struct {
	users []*model.User
}

// NewUserRepo creates a new user repository.
func NewUserRepo() model.UserRepo {
	return &userGateway{users: make([]*model.User, 0)}
}

func (gw *userGateway) List(tenantID string) ([]*model.User, error) {
	users := []*model.User{}
	for _, user := range gw.users {
		if user.TenantID == tenantID {
			users = append(users, user)
		}
	}
	return users, nil
}

func (gw *userGateway) Get(tenantID, userID string) (*model.User, error) {
	for _, user := range gw.users {
		if user.TenantID == tenantID && user.ID == userID {
			return user, nil
		}
	}
	return nil, model.ErrEntityNotFound
}

func (gw *userGateway) Create(user *model.User) error {
	// TODO: validation check
	user.ID = model.NewUUID()
	gw.users = append(gw.users, user)
	return nil
}

func (gw *userGateway) Update(user *model.User) error {
	for i, u := range gw.users {
		if u.TenantID == user.TenantID && u.ID == user.ID {
			gw.users[i] = user
			return nil
		}
	}
	return model.ErrEntityNotFound
}

func (gw *userGateway) Delete(user *model.User) error {
	for i, u := range gw.users {
		if u.TenantID == user.TenantID && u.ID == user.ID {
			gw.users = append(gw.users[:i], gw.users[i+1:]...)
			return nil
		}
	}
	return model.ErrEntityNotFound
}
