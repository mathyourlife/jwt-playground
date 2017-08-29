// https://github.com/ripple-cloud/cloud/blob/af5b8e09cb5e742f22ea394daff517af6fc2e830/data/hub.go
package main

import (
	"fmt"
	"log"

	"github.com/mathyourlife/jwt-playground/pkg/authdb"
)

func main() {
	log.Println("starting test")

	c, err := authdb.NewPostgresConfig()
	if err != nil {
		log.Fatal(err)
	}

	c.Options["user"] = "postgres"
	c.Options["password"] = "e5d59d860efb42f63a1d04d31b5c590e"

	pg, err := authdb.NewPostgres(c)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(pg)

	err = pg.CreateAppDB()
	if err != nil {
		log.Fatal(err)
	}

	c.Options["user"] = "app_user"
	c.Options["password"] = "app_user_password"
	c.Options["dbname"] = "app_database"

	pg, err = authdb.NewPostgres(c)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(pg)

	err = pg.SetupAppDB()
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 4; i++ {
		u1 := &authdb.User{Username: fmt.Sprintf("%s-%d", "dcouture", i), Password: "pass", Salt: "salt"}
		_, err = pg.User(u1, authdb.OpCreate)
		if err != nil {
			log.Fatal(err)
		}
	}

}
