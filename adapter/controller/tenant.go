package controller

import (
	"net/http"

	"github.com/ambi/go-web-app-patterns/model"
	"github.com/ambi/go-web-app-patterns/service"
	"github.com/labstack/echo/v4"
)

// ListTenants lists tenants in response to HTTP request.
func ListTenants(c echo.Context, tenantRepo model.TenantRepo) error {
	tenants, err := service.ListTenants(tenantRepo)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, tenants)
	return nil
}

// GetTenant gets a tenant in response to HTTP request.
func GetTenant(c echo.Context, tenantRepo model.TenantRepo) error {
	// TODO: validation
	id := c.Param("id")
	tenant, err := service.GetTenant(tenantRepo, id)
	if err == model.ErrEntityNotFound {
		c.JSON(http.StatusNotFound, &ErrorResponse{Error: "Tenant not found"})
		return nil
	}
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, tenant)
	return nil
}

// CreateTenant creates a tenant in response to HTTP request.
func CreateTenant(c echo.Context, tenantRepo model.TenantRepo) error {
	tenant, err := service.CreateTenant(tenantRepo)
	if err != nil {
		return err
	}

	c.JSON(http.StatusCreated, tenant)
	return nil
}

// DeleteTenant deletes a tenant in response to HTTP request.
func DeleteTenant(c echo.Context, tenantRepo model.TenantRepo) error {
	// TODO: validation
	id := c.Param("id")

	err := service.DeleteTenant(tenantRepo, id)
	if err == model.ErrEntityNotFound {
		c.JSON(http.StatusNotFound, &ErrorResponse{Error: "Tenant not found"})
		return nil
	}
	if err != nil {
		return err
	}

	c.NoContent(http.StatusNoContent)
	return nil
}
