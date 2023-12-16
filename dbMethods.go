package main

import (
	"database/sql"
	"fmt"
	"log"
)

type user struct {
	id       int
	username string
	password string
}

func inicializeDB() *sql.DB {
	db, err := sql.Open("mysql", "user:password@(localhost)/datanase?parseTime=true")

	checkErr(err)
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Servidor conectado!")
	}

	return db

}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
