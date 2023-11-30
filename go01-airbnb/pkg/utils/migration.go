package utils

import (
	"fmt"
	"go01-airbnb/config"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunDBMigration(c *config.Config) {
	dbURL := fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s",
		c.MySQL.User,
		c.MySQL.Password,
		c.MySQL.Host,
		c.MySQL.Port,
		c.MySQL.DBName,
	)

	migration, err := migrate.New(c.App.MigrationURL, dbURL)
	if err != nil {
		log.Fatal("Cannot create migration instance", err)
	}

	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Cannot run migrate up", err)
	}
}
