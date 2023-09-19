package db_test

import (
	"barcode/db"
	"testing"
)

func TestDatabase(t *testing.T) {
	t.Run("test create database", func(t *testing.T) {
		db.CreateDatabase()
	})

	t.Run("test connection to sqlite", func(t *testing.T) {
		db.ConnectionToSql3()
	})
}
