package migration

import (
	"github.com/spurtcms/ecommerce/migration/mysql"
	"github.com/spurtcms/ecommerce/migration/postgres"
	"gorm.io/gorm"
)

func AutoMigration(DB *gorm.DB, dbtype string) {

	if dbtype == "postgres" {

		postgres.MigrationTables(DB) //auto migrate table

	} else if dbtype == "mysql" {

		mysql.MigrationTables(DB) //auto migrate table
	}

}
