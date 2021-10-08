package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Dept string `json:"department"`
}

var data []Tag

func main() {
	db, err := sql.Open("mysql", "user:pass@tcp(server.example.com:3306)/databasename")
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Method to fetch rows returned
	results, err := db.Query("SELECT * FROM people")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var id int
		var name, department string
		// for each row, scan the result into our tag composite object

		err = results.Scan(&id, &name, &department)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		data = append(data, Tag{id, name, department})
	}

	fmt.Printf("%T", data)
	dataBytes, _ := json.Marshal(&data)
	fmt.Println(string(dataBytes))
}
