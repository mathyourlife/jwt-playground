package main

import (
	"fmt"
	"log"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func main() {
	log.Println("hi")

	example1()
	example2()
}

func example2() {
	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Println("failed to open sqlmock database:", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "title"}).
		AddRow(1, "one").
		AddRow(2, "two")
	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	db.Query("SELECT * from here")
}

func example1() {
	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Println("failed to open sqlmock database:", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "title"}).
		AddRow(1, "one").
		AddRow(2, "two")

	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	rs, _ := db.Query("SELECT")
	defer rs.Close()

	for rs.Next() {
		var id int
		var title string
		rs.Scan(&id, &title)
		fmt.Println("scanned id:", id, "and title:", title)
	}

	if rs.Err() != nil {
		fmt.Println("got rows error:", rs.Err())
	}

}
