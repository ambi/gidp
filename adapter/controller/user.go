package controller

import (
	"net/http"

	"github.com/ambi/go-web-app-patterns/model"
	"github.com/ambi/go-web-app-patterns/service"
	"github.com/labstack/echo/v4"
)

// ListUsers lists users in response to HTTP request.
func ListUsers(c echo.Context, tenantRepo model.TenantRepo, userRepo model.UserRepo) error {
	// TODO: validation
	tenantID := c.Param("tenant_id")

	users, err := service.ListUsers(tenantRepo, userRepo, tenantID)
	if err == model.ErrEntityNotFound {
		return c.JSON(http.StatusNotFound, &ErrorResponse{Error: "Tenant not found"})
	}
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

// GetUser gets a user in response to HTTP request.
func GetUser(c echo.Context, tenantRepo model.TenantRepo, userRepo model.UserRepo) error {
	// TODO: validation
	tenantID := c.Param("tenant_id")
	id := c.Param("id")

	user, err := service.GetUser(tenantRepo, userRepo, tenantID, id)
	if err == model.ErrEntityNotFound {
		return c.JSON(http.StatusNotFound, &ErrorResponse{Error: "Tenant or user not found"})
	}
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

// CreateUser creates a user in response to HTTP request.
func CreateUser(c echo.Context, tenantRepo model.TenantRepo, userRepo model.UserRepo) error {
	// TODO: validation
	tenantID := c.Param("tenant_id")
	displayName := c.FormValue("display_name")

	user := &model.User{
		DisplayName: displayName,
	}

	err := service.CreateUser(tenantRepo, userRepo, tenantID, user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, user)
}

// UpdateUser updates a user in response to HTTP request.
func UpdateUser(c echo.Context, tenantRepo model.TenantRepo, userRepo model.UserRepo) error {
	// TODO: validation
	tenantID := c.Param("tenant_id")
	id := c.Param("id")
	displayName := c.FormValue("display_name")

	user := &model.User{ID: id, DisplayName: displayName}
	err := service.UpdateUser(tenantRepo, userRepo, tenantID, user)
	if err == model.ErrEntityNotFound {
		return c.JSON(http.StatusNotFound, &ErrorResponse{Error: "User not found"})
	}
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

// DeleteUser deletes a user in response to HTTP request.
func DeleteUser(c echo.Context, tenantRepo model.TenantRepo, userRepo model.UserRepo) error {
	// TODO: validation
	tenantID := c.Param("tenant_id")
	id := c.Param("id")

	err := service.DeleteUser(tenantRepo, userRepo, tenantID, id)
	if err == model.ErrEntityNotFound {
		return c.JSON(http.StatusNotFound, &ErrorResponse{Error: "Tenant or user not found"})
	}
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
