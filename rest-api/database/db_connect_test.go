package database

import (
	"testing"
)

const DBPath = "olist.db"

func TestConnectDB(t *testing.T) {
	db := ConnectDB(DBPath)
	if db == nil {
		t.Error("Database connection failed")
	}
}
