package model

// User is a user model.
type User struct {
	TenantID    string
	ID          string
	DisplayName string
}

// UserRepo is a repository interface for User.
type UserRepo interface {
	List(tenantID string) ([]*User, error)
	Get(tenantID, userID string) (*User, error)
	Create(u *User) error
	Update(u *User) error
	Delete(u *User) error
}
