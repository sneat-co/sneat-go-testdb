package testdb

import "testing"

func TestNewMockDB(t *testing.T) {
	db := NewMockDB(t)
	if db == nil {
		t.Fatal("db == nil")
	}
}
