package main

import (
	"database/sql"
	"errors"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/k0kubun/pp/v3"
	"gopkg.in/guregu/null.v4"
	// _ "github.com/lib/pq"
)

type Person struct {
	FirstName      string    `db:"first_name"`
	LastName       string    `db:"last_name"`
	Email          string    `db:"email"`
	CreatedAt      time.Time `db:"created_at"`
	BirthDate      null.Time `db:"birth_date"`
	LastReservedAt PgTime    `db:"last_reserved_at"`
}

var _ sql.Scanner = (*PgTime)(nil)

type PgTime null.Time

func (t *PgTime) Scan(src interface{}) error {
	switch v := src.(type) {
	case time.Time:
		*t = PgTime(null.TimeFrom(v.In(time.Local)))
		return nil
	case nil:
		*t = PgTime(null.Time{})
		return nil
	default:
		return errors.New("incompatible type for PgTime")
	}
}

func main() {
	// connStr := "user=postgres password=password dbname=postgres sslmode=disable port=5432 timezone=Asia/Tokyo"
	connStr := "postgres://postgres:password@127.0.0.1:5432/postgres?sslmode=disable&timezone=Asia/Tokyo"
	// db, err := sqlx.Connect("postgres", connStr)
	// if err != nil {
	//   log.Fatal(err)
	// }
	// defer db.Close()

	db, err := sqlx.Connect("pgx", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Queryx("SELECT first_name, last_name, email, birth_date, created_at, last_reserved_at FROM persons")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var persons []Person

	for rows.Next() {
		var person Person
		if err := rows.StructScan(&person); err != nil {
			panic(err)
		}
		persons = append(persons, person)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}

	pp.Println(persons)
}
