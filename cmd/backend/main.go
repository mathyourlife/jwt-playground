package main

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/mathyourlife/jwt-playground/pkg/backend"
)

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile | log.LUTC)
	c, err := backend.NewPostgresConfig()
	for err != nil {
		log.Fatal(err)
	}
	c.Options["host"] = "db"
	c.Options["user"] = "postgres"
	c.Options["password"] = "mysecretpassword2"

	db, err := backend.Postgres(c)
	for err != nil {
		log.Fatal(err)
	}

	setup(db)
}

func setup(db *sqlx.DB) {
	statements := []string{
		"DROP DATABASE IF EXISTS app",
		"DROP USER IF EXISTS adminuser",
		"DROP ROLE IF EXISTS appuser",
		"CREATE DATABASE app ENCODING 'UTF8'",
		"CREATE ROLE appuser LOGIN CREATEROLE",
		"CREATE USER adminuser PASSWORD 'aoeu'",
		"GRANT appuser to adminuser",
		"GRANT ALL PRIVILEGES ON DATABASE app to appuser",
	}

	for _, q := range statements {
		_, err := db.Exec(q)
		if err != nil {
			log.Printf("%#v\n", err)
		}
	}
	time.Sleep(1 * time.Minute)
}
