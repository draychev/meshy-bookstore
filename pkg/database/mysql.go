// Package database allows the bookwarehouse service to store
// total books data into MySQL persistent storage
package database

import (
	"fmt"
	"time"

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

var db *gorm.DB

// GetMySQLConnection returns a MySQL connection using default configuration
func GetMySQLConnection() (*gorm.DB, error) {
	mySQLHost := envvar.GetEnv("MYSQL_HOST", "mysql.bookwarehouse.svc.cluster.local")
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&timeout=20s",
		dbuser, dbpass, mySQLHost, dbport, dbname)
	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})

	return db, err
}

func Init() {
	log.Info().Msg("Initializig database...")
	var err error
	for {
		if db, err = GetMySQLConnection(); err != nil {
			log.Info().Msg("Booksdemo database is not ready. Wait for 10s ...")
			time.Sleep(10 * time.Second)
		} else {
			break
		}
	}

	log.Info().Msg("Booksdemo database is connected.")
	if err := db.Migrator().AutoMigrate(&Record{}); err != nil {
		log.Fatal().Msgf("Database migration failed. %v", err)
	}

	var record Record
	if result := db.Where(&Record{Key: KeyTotalBooks}).First(&record); result.RowsAffected == 0 {
		// initialize record
		record = Record{
			Key:      KeyTotalBooks,
			ValueInt: 0,
		}

		result = db.Create(&record)
		log.Info().Msgf("Initial %s record created. %v, %v, %v", KeyTotalBooks, record, result.RowsAffected, result.Error)
	}
}

func Update(numberOfBooks int) int {
	var record Record
	db.Where(&Record{Key: KeyTotalBooks}).First(&record)
	log.Info().Msgf("Fetching Books Stock Record: %+v", record)
	record.ValueInt += int64(numberOfBooks)
	totalBooks := int(record.ValueInt)
	db.Save(record)
	return totalBooks
}
