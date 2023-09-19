package migrations

import (
	"barcode/models/entities"
	"log"

	"gorm.io/gorm"
)

const (
	success = "success"
	failed  = "failed"
)

func RegisteredTable() []interface{} {
	tables := []interface{}{
		&entities.Barcode{},
	}
	return tables
}

func Down(pgConn *gorm.DB) error {
	tables := RegisteredTable()
	for i := 0; i < len(tables); i++ {
		err := pgConn.Migrator().DropTable(tables[i])
		if err != nil {
			log.Print(failed + " down")
			return err
		}
	}
	log.Print(success + " down")
	return nil
}

func Up(pgConn *gorm.DB) error {
	Down(pgConn)
	tables := RegisteredTable()
	for _, v := range tables {
		err := pgConn.AutoMigrate(v)
		if err != nil {
			Down(pgConn)
			log.Print(failed + " up")
			return err
		}
	}
	log.Print(success + " up")
	return nil
}
