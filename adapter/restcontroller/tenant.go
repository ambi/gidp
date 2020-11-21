package restcontroller

import (
	"net/http"

	"github.com/ambi/gidp/model"
	"github.com/ambi/gidp/service"
	"github.com/labstack/echo/v4"
)

// ListTenants lists tenants in response to HTTP request.
func ListTenants(c echo.Context, tenantRepo model.TenantRepo) error {
	tenants, err := service.ListTenants(tenantRepo)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, tenants)
}

// GetTenant gets a tenant in response to HTTP request.
func GetTenant(c echo.Context, tenantRepo model.TenantRepo) error {
	// TODO: validation
	id := c.Param("id")
	tenant, err := service.GetTenant(tenantRepo, id)
	if err == model.ErrEntityNotFound {
		return c.JSON(http.StatusNotFound, &ErrorResponse{Error: "Tenant not found"})
	}
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, tenant)
}

// CreateTenant creates a tenant in response to HTTP request.
func CreateTenant(c echo.Context, tenantRepo model.TenantRepo) error {
	tenant, err := service.CreateTenant(tenantRepo)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, tenant)
}

// DeleteTenant deletes a tenant in response to HTTP request.
func DeleteTenant(c echo.Context, tenantRepo model.TenantRepo) error {
	// TODO: validation
	id := c.Param("id")

	err := service.DeleteTenant(tenantRepo, id)
	if err == model.ErrEntityNotFound {
		return c.JSON(http.StatusNotFound, &ErrorResponse{Error: "Tenant not found"})
	}
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
