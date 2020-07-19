package main

import (
	"database/sql"

	"github.com/ambi/go-web-app-patterns/adapter/controller"
	"github.com/ambi/go-web-app-patterns/adapter/sqlgateway"
	"github.com/ambi/go-web-app-patterns/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	tenantRepo model.TenantRepo
	userRepo   model.UserRepo
)

func setupRouter(e *echo.Echo) {
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

func setupDB() error {
	db, err := sql.Open("mysql", "root:@/go_web_app_patterns")
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	tenantRepo = sqlgateway.NewTenantRepo(db)
	userRepo = sqlgateway.NewUserRepo(db)

	// adapter, err := mysql.Open("root:@/go_web_app_patterns")
	// if err != nil {
	// 	return err
	// }
	// ctx := context.Background()
	// if err = adapter.Ping(ctx); err != nil {
	// 	return err
	// }

	// repo := rel.New(adapter)

	// tenantRepo = relgateway.NewTenantRepo(repo)
	// userRepo = relgateway.NewUserRepo(repo)

	return nil
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	setupRouter(e)

	err := setupDB()
	if err != nil {
		e.Logger.Fatal(err)
	}

	err = e.Start(":8080")
	if err != nil {
		e.Logger.Fatal(err)
	}
}
