package repository

import (
	"os"
	"testing"

	"github.com/Daniel-Vinicius/golang-live/adapter/repository/fixture"
	"github.com/stretchr/testify/assert"
)

func TestTransactionRepositoryDb_Insert(t *testing.T) {
	migrationsDir := os.DirFS("fixture/sql")
	db := fixture.Up(migrationsDir)
	defer fixture.Down(db, migrationsDir)

	repository := NewTransactionRepositoryDb(db)
	err := repository.Insert("1", "1", 100, "approved", "")
	assert.Nil(t, err)
}
