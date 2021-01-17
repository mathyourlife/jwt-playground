package main

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/urfave/cli"

	"github.com/mathyourlife/jwt-playground/pkg/authdb"
)

func main() {
	log.Println("starting postgres testing")

	app := cli.NewApp()
	app.Name = "say"
	app.Version = "0.0.1"
	app.Authors = []cli.Author{{Name: "Dan Couture", Email: "mathyourlife@gmail.com"}}
	app.EnableBashCompletion = true
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "pg.port",
			Value: 5432,
			Usage: "Postgres port",
		},
	}
	app.Action = realMain
	app.Run(os.Args)

}
func old() {
	db, err := sqlx.Open("postgres", "postgres://postgres:mysecretpassword@localhost/postgres?sslmode=disable&port=32768")

	if err != nil {
		log.Fatal(err)
	}

	u := authdb.NewUser(db)
	log.Println(u.CreateTable())

	// err = u.CreateTable()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	return

	// res, err = db.Exec(`drop table USERS`)
	// log.Println(res, err)
	// sql := `CREATE TABLE IF NOT EXISTS USERS(
	//    user_id  INT          NOT NULL,
	//    username VARCHAR(255) NOT NULL,
	//    password VARCHAR(255) NOT NULL,
	//    salt     CHAR(50)     NOT NULL,
	//    created  timestamp    NOT NULL default current_timestamp,
	//    updated  timestamp    NOT NULL default current_timestamp,
	//    PRIMARY KEY( user_id )
	// )`

	// rows, err := db.Query(sql)

	// log.Println(rows, err)

	// rows, err = db.Query(`INSERT INTO USERS (user_id, username, password, salt)
	// VALUES
	// (3, 'user', 'password', 'salt')`)

	// rows, err = db.Query("select * from users")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for rows.Next() {
	// 	var uid int
	// 	var username, password, salt string
	// 	var created, updated time.Time
	// 	err = rows.Scan(&uid, &username, &password, &salt, &created, &updated)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Println(uid, username, password, salt, created, updated)
	// }
}

func realMain(ctx *cli.Context) error {
	config, err := authdb.NewDBConfig()
	if err != nil {
		return err
	}
	config.Options["user"] = "postgres"
	config.Options["password"] = "mysecretpassword"
	config.Options["port"] = ctx.String("pg.port")

	db, err := authdb.NewDB(config)
	if err != nil {
		log.Println(err)
		return err
	}
	setup(ctx, db)

	return nil
}

func setup(ctx *cli.Context, db *authdb.DB) {
	log.Println(db)

	u := db.User()
	err := u.CreateTable()
	if err != nil {
		log.Fatal(err)
	}
	u.Username = "dcoutur"
	u.Password = "ssssss"
	u.Salt = "1235"
	err = u.Create()
	if err != nil {
		log.Fatal(err)
	}
}

var FlagPGPort = cli.IntFlag{
	Name:  "pg.port",
	Value: 5432,
	Usage: "Postgres port",
}
