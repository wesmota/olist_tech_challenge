package reader

import (
	"testing"
)

const csvFile = "testdata/authors.csv"
const dbFile = "testdata/olist_test.db"

func TestLoader(t *testing.T) {
	// load data
	loader := NewCSVLoader(csvFile, dbFile)
	err := loader.Load()
	if err != nil {
		t.Error(err)
		t.Fail()
	}

}
