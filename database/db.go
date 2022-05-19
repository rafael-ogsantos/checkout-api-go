package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
)

var (
	DB *bun.DB
)

func Connect() {
	dsn := "user=root dbname=root password=root sslmode=disable"
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal(err)
	}

	// defer db.Close()

	DB = bun.NewDB(db, pgdialect.New())

	DB.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))
}
