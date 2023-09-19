package db

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const dbName = "./barcode.db"

func ConnectionToSql3() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Print("connection to sqlite open ...")
	return db
}

func CreateDatabase() {
	os.Remove(dbName)
	log.Printf("creating databse %v ...", dbName)
	file, err := os.Create(dbName)
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Printf("success creating database %v ...", dbName)
}
