package model

// Tenant is an organization which users, groups and other resources belong to.
type Tenant struct {
	ID     string
	Status string
}

// TenantRepo is a repository interface for Tenant.
type TenantRepo interface {
	List() ([]*Tenant, error)
	Get(id string) (*Tenant, error)
	Create(tenant *Tenant) error
	// Update(tenant *Tenant) error
	Delete(tenant *Tenant) error
}
