package api

import "github.com/ambi/go-web-app-patterns/model"

// EncodeTenant encodes a tenant into protobuf.
func EncodeTenant(tenant *model.Tenant) *Tenant {
	return &Tenant{
		Id:     tenant.ID,
		Status: tenant.Status,
	}
}

// EncodeUser encodes a user into protobuf.
func EncodeUser(user *model.User) *User {
	return &User{
		Id:          user.ID,
		TenantId:    user.TenantID,
		DisplayName: user.DisplayName,
	}
}

// DecodeTenant decodes a tenant from protobuf.
func DecodeTenant(tenant *Tenant) *model.Tenant {
	return &model.Tenant{
		ID:     tenant.Id,
		Status: tenant.Status,
	}
}

// DecodeUser decodes a user from protobuf.
func DecodeUser(user *User) *model.User {
	return &model.User{
		ID:          user.Id,
		TenantID:    user.TenantId,
		DisplayName: user.DisplayName,
	}
}
