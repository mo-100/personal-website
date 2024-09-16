package db_utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"personal/internal/config"

	_ "github.com/lib/pq"
	_ "modernc.org/sqlite"
)

func GetDB(config *config.Config) *sql.DB {
	if config.Environment == "prod" {
		return mustOpen(config.DatabaseDriverName, getConnectionString(config))
	}
	conn := mustOpen("sqlite", ":memory:")
	err := runFile(conn, "./internal/store/schema.sql")
	if err != nil {
		log.Panic(err)
	}
	err = runFile(conn, "./internal/store/gen.sql")
	if err != nil {
		log.Panic(err)
	}
	return conn
}
func mustOpen(driver, dst string) *sql.DB {
	conn, err := sql.Open(driver, dst)
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
func runFile(conn *sql.DB, path string) error {
	sqlText, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	_, err = conn.Exec(string(sqlText))
	if err != nil {
		return err
	}
	return nil
}
