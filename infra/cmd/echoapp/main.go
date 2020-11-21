package main

import (
	"github.com/ambi/gidp/adapter/restcontroller"
	"github.com/ambi/gidp/adapter/sqlgateway"
	"github.com/ambi/gidp/infra"
	"github.com/ambi/gidp/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	defaultPort = "8080"
)

func route(e *echo.Echo, tenantRepo model.TenantRepo, userRepo model.UserRepo) {
	e.GET("/tenants",
		func(c echo.Context) error { return restcontroller.ListTenants(c, tenantRepo) })
	e.GET("/tenants/:id",
		func(c echo.Context) error { return restcontroller.GetTenant(c, tenantRepo) })
	e.POST("/tenants",
		func(c echo.Context) error { return restcontroller.CreateTenant(c, tenantRepo) })
	e.DELETE("/tenants/:id",
		func(c echo.Context) error { return restcontroller.DeleteTenant(c, tenantRepo) })
	e.GET("/:tenant_id/users",
		func(c echo.Context) error { return restcontroller.ListUsers(c, tenantRepo, userRepo) })
	e.GET("/:tenant_id/users/:id",
		func(c echo.Context) error { return restcontroller.GetUser(c, tenantRepo, userRepo) })
	e.POST("/:tenant_id/users",
		func(c echo.Context) error { return restcontroller.CreateUser(c, tenantRepo, userRepo) })
	e.PATCH("/:tenant_id/users/:id",
		func(c echo.Context) error { return restcontroller.UpdateUser(c, tenantRepo, userRepo) })
	e.DELETE("/:tenant_id/users/:id",
		func(c echo.Context) error { return restcontroller.DeleteUser(c, tenantRepo, userRepo) })
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	db, err := infra.NewMySQLDB()
	if err != nil {
		e.Logger.Fatal(err)
	}
	tenantRepo := sqlgateway.NewTenantRepo(db)
	userRepo := sqlgateway.NewUserRepo(db)

	route(e, tenantRepo, userRepo)

	err = e.Start(":" + infra.GetPort(defaultPort))
	if err != nil {
		e.Logger.Fatal(err)
	}
}
