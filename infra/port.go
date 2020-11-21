package infra

import (
	"os"
	"strconv"
)

// GetPort gets the port number specified in PORT env-var or the default value.
func GetPort(defaultPort string) string {
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
