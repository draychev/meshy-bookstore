// Package database allows the bookwarehouse service to store
// total books data into MySQL persistent storage
package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/draychev/go-toolbox/pkg/envvar"
)

// Refer to https://github.com/openservicemesh/osm-docs/blob/main/manifests/apps/mysql.yaml for database setup
const (
	dbuser = "root"
	dbpass = "mypassword"
	dbport = 3306
	dbname = "booksdemo"
)

// GetMySQLConnection returns a MySQL connection using default configuration
func GetMySQLConnection() (*gorm.DB, error) {
	mySQLHost := envvar.GetEnv("MYSQL_HOST", "mysql.bookwarehouse.svc.cluster.local")
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&timeout=20s", dbuser, dbpass, mySQLHost, dbport, dbname)
	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})

	return db, err
}
