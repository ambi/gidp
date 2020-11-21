package infra

import (
	"database/sql"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

const (
	defaultDBHost     = "localhost"
	defaultDBName     = "gidp"
	defaultDBPassword = ""
	defaultDBUser     = "root"
)

func getDBName() string {
	name := os.Getenv("DB_NAME")
	if len(name) == 0 {
		return defaultDBName
	}
	return name
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

func getLocalLocation() *time.Location {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		panic(err)
	}

	return loc
}

// NewMySQLDB creates a new DB interface for MySQL.
func NewMySQLDB() (*sql.DB, error) {
	config := mysql.NewConfig()
	config.Addr = getDBHost()
	config.DBName = getDBName()
	config.Loc = getLocalLocation()
	config.Net = "tcp"
	config.ParseTime = true
	config.Passwd = getDBPassword()
	config.User = getDBUser()

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
