package main

import "testing"

// TODO: Implement TestConnectToDB
func TestconnectToDB(t *testing.T) {
	app := &application{
		DSN: "user=username password=password host=hostname port=5432 dbname=dbname sslmode=disable",
	}
	if _, err := connectToDB(dsn); err != nil {
		t.Fatal(err)
	}
}

func TestOpenDB(t *testing.T) {
	dsn := "host=postgres port=5432 user=postgres password=postgres dbname=movies sslmode=disable timezone=UTC connect_timeout=5"
	if _, err := openDB(dsn); err != nil {
		t.Fatal(err)
	}
}
