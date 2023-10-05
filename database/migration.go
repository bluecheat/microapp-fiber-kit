package database

import "log"

func AutoMigration(db *Database) {
	if err := db.gormDb.AutoMigrate(
		MigrationDomains...,
	); err != nil {
		log.Fatal(err)
	}
}

var MigrationDomains = []interface{}{}
