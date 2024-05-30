package migration

import (
	mysql "github.com/spurtcms/team-roles/migration/mysql"
	postgres "github.com/spurtcms/team-roles/migration/postgres"
	"gorm.io/gorm"
)

func AutoMigration(DB *gorm.DB, dbtype any) {

	if dbtype == "postgres" {

		postgres.MigrationTables(DB) //auto migrate table

	} else if dbtype == "mysql" {

		mysql.MigrationTables(DB) //auto migrate table
	}

}
