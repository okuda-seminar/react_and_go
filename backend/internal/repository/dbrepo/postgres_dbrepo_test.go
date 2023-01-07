package dbrepo

import (
	"database/sql"
	"reflect"
	"testing"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func TestConnection(t *testing.T) {
	dsn := "host=postgres port=5432 user=postgres password=postgres dbname=movies sslmode=disable timezone=UTC connect_timeout=5"
	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		t.Fatal(err)
	}

	if err = conn.Ping(); err != nil {
		t.Fatal(err)
	}

	repo := &PostgresDBRepo{DB: conn}
	defer conn.Close()

	out := repo.Connection()
	if !reflect.DeepEqual(out, conn) {
		t.Errorf("expected: %v, but got: %v", "", out)
	}
}

func TestAllMovies(t *testing.T) {
	dsn := "host=postgres port=5432 user=postgres password=postgres dbname=movies sslmode=disable timezone=UTC connect_timeout=5"
	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		t.Fatal(err)
	}

	if err = conn.Ping(); err != nil {
		t.Fatal(err)
	}

	repo := &PostgresDBRepo{DB: conn}
	defer conn.Close()

	// TODO: add mock data
	if _, err := repo.AllMovies(); err != nil {
		t.Fatal(err)
	}
}
