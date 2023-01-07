package main

import "testing"

// TODO: Implement TestConnectToDB

func TestOpenDB(t *testing.T) {
	dsn := "host=postgres port=5432 user=postgres password=postgres dbname=movies sslmode=disable timezone=UTC connect_timeout=5"
	if _, err := openDB(dsn); err != nil {
		t.Fatal(err)
	}
}
