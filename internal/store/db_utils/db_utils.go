package db_utils

import (
	"database/sql"
	"fmt"
	"log"
	"personal/internal/config"

	_ "github.com/lib/pq"
)

func MustOpen(config *config.Config) *sql.DB {
	conn, err := sql.Open(config.DatabaseDriverName, getConnectionString(config))
	if err != nil {
		log.Panic(err)
	}
	return conn
}
func getConnectionString(config *config.Config) string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=verify-full",
		config.DatabaseUsername,
		config.DatabasePassword,
		config.DatabaseURL,
		config.DatabasePort,
		config.DatabaseName,
	)
}
