package db

import (
	"barcode/config"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const DbName = "/barcode.db"

func ConnectionToSql3() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(config.GetDbName(DbName)), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Print("connection to sqlite open ...")
	return db
}

func CreateDatabase() {
	db := config.GetDbName(DbName)
	os.Remove(db)
	log.Printf("creating databse %v ...", db)
	file, err := os.Create(db)
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Printf("success creating database %v ...", db)
}
