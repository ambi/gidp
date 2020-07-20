package infra

import (
	"database/sql"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

const (
	defaultDBHost     = "localhost"
	defaultDBPassword = ""
	defaultDBUser     = "root"
	databaseName      = "go_web_app_patterns"
)

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

// NewMySQLDB creates a new DB interface for MySQL.
func NewMySQLDB() (*sql.DB, error) {
	config := mysql.NewConfig()
	config.User = getDBUser()
	config.Passwd = getDBPassword()
	config.Net = "tcp"
	config.Addr = getDBHost()
	config.DBName = databaseName
	config.Loc = time.FixedZone("Asia/Tokyo", 9*60*60)
	config.ParseTime = true

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
