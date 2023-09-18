package test

import (
	"github.com/dal-go/dalgo2buntdb"
	"github.com/sneat-co/sneat-go-testdb"
	"testing"
)

func TestNewMockDB(t *testing.T) {
	db := dalgo2buntdb.NewInMemoryMockDB(t)
	testdb.NewMockDB(t, db)
	if db == nil {
		t.Fatal("db == nil")
	}
}
