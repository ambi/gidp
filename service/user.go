package service

import (
	"github.com/ambi/go-web-app-patterns/model"
)

// ListUsers lists users.
func ListUsers(tenantRepo model.TenantRepo, userRepo model.UserRepo, tenantID string) ([]*model.User, error) {
	tenant, err := tenantRepo.Get(tenantID)
	if err != nil {
		return nil, err
	}

	return userRepo.List(tenant.ID)
}

// GetUser gets a User.
func GetUser(tenantRepo model.TenantRepo, userRepo model.UserRepo, tenantID, userID string) (*model.User, error) {
	tenant, err := tenantRepo.Get(tenantID)
	if err != nil {
		return nil, err
	}

	return userRepo.Get(tenant.ID, userID)
}

// CreateUser creates a new user.
func CreateUser(tenantRepo model.TenantRepo, userRepo model.UserRepo, tenantID string, user *model.User) error {
	tenant, err := tenantRepo.Get(tenantID)
	if err != nil {
		return err
	}
	user.TenantID = tenant.ID

	return userRepo.Create(user)
}

// UpdateUser Updates a User.
func UpdateUser(tenantRepo model.TenantRepo, userRepo model.UserRepo, tenantID string, user *model.User) error {
	tenant, err := tenantRepo.Get(tenantID)
	if err != nil {
		return err
	}
	dbUser, err := userRepo.Get(tenant.ID, user.ID)
	if err != nil {
		return err
	}
	user.ID = dbUser.ID

	return userRepo.Update(user)
}

// DeleteUser deletes a User.
func DeleteUser(tenantRepo model.TenantRepo, userRepo model.UserRepo, tenantID, userID string) error {
	tenant, err := tenantRepo.Get(tenantID)
	if err != nil {
		return err
	}

	user, err := userRepo.Get(tenant.ID, userID)
	if err != nil {
		return err
	}

	return userRepo.Delete(user)
}
