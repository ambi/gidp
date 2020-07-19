package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/ambi/go-web-app-patterns/adapter/controller"
	"github.com/ambi/go-web-app-patterns/adapter/sqlgateway"
	"github.com/ambi/go-web-app-patterns/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	defaultDBHost     = "localhost"
	defaultDBPassword = ""
	defaultDBUser     = "root"
	databaseName      = "go_web_app_patterns"
	defaultPort       = "8080"
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

func getDBUser() string {
	user := os.Getenv("DB_USER")
	if len(user) == 0 {
		return defaultDBUser
	}
	return user
}

func getDBHost() string {
	host := os.Getenv("DB_HOST")
	if len(host) == 0 {
		return defaultDBHost
	}
	return host
}
func getDBPassword() string {
	password := os.Getenv("DB_PASSWORD")
	if len(password) == 0 {
		return defaultDBPassword
	}
	return password
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

func newDB() error {
	datasource := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", getDBUser(), getDBPassword(), getDBHost(), databaseName) // TODO: escape
	db, err := sql.Open("mysql", datasource)
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

	err := newDB()
	if err != nil {
		e.Logger.Fatal(err)
	}

	err = e.Start(":" + getPort())
	if err != nil {
		e.Logger.Fatal(err)
	}
}
