package database

import (
	"log"
	"microapp-fiber-kit/domains"
)

func AutoMigration(db *Database) {
	if err := db.gormDb.AutoMigrate(
		MigrationDomains...,
	); err != nil {
		log.Fatal(err)
	}
}

var MigrationDomains = []interface{}{
	&domains.User{},
	&domains.Board{},
}
