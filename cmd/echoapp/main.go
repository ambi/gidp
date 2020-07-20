package main

import (
	"os"
	"strconv"

	"github.com/ambi/go-web-app-patterns/adapter/controller"
	"github.com/ambi/go-web-app-patterns/adapter/sqlgateway"
	"github.com/ambi/go-web-app-patterns/infra"
	"github.com/ambi/go-web-app-patterns/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	defaultPort = "8080"
)

func route(e *echo.Echo, tenantRepo model.TenantRepo, userRepo model.UserRepo) {
	e.GET("/tenants", func(c echo.Context) error { return controller.ListTenants(c, tenantRepo) })
	e.GET("/tenants/:id", func(c echo.Context) error { return controller.GetTenant(c, tenantRepo) })
	e.POST("/tenants", func(c echo.Context) error { return controller.CreateTenant(c, tenantRepo) })
	e.DELETE("/tenants/:id", func(c echo.Context) error { return controller.DeleteTenant(c, tenantRepo) })
	e.GET("/:tenant_id/users", func(c echo.Context) error { return controller.ListUsers(c, tenantRepo, userRepo) })
	e.GET("/:tenant_id/users/:id", func(c echo.Context) error { return controller.GetUser(c, tenantRepo, userRepo) })
	e.POST("/:tenant_id/users", func(c echo.Context) error { return controller.CreateUser(c, tenantRepo, userRepo) })
	e.PATCH("/:tenant_id/users/:id", func(c echo.Context) error { return controller.UpdateUser(c, tenantRepo, userRepo) })
	e.DELETE("/:tenant_id/users/:id", func(c echo.Context) error { return controller.DeleteUser(c, tenantRepo, userRepo) })
}

func getPort() string {
	port := os.Getenv("PORT")
	i, err := strconv.Atoi(port)
	if err != nil {
		return defaultPort
	}
	if i < 0 || 65535 < i {
		return defaultPort
	}
	return port
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

	err = e.Start(":" + getPort())
	if err != nil {
		e.Logger.Fatal(err)
	}
}
