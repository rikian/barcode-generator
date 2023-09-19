package migrations_test

import (
	"barcode/db"
	"barcode/migrations"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMigration(t *testing.T) {
	dbConn := db.ConnectionToSql3()
	t.Run("test up", func(t *testing.T) {
		err := migrations.Up(dbConn)
		assert.Nil(t, err)
	})

	t.Run("test down", func(t *testing.T) {
		err := migrations.Down(dbConn)
		assert.Nil(t, err)
	})
}
