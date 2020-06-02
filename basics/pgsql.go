package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "123456"
	dbname   = "root"
)

type Tmp struct {
	id   int
	name string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	sqlStatement := `SELECT * FROM tmp WHERE id=$1;`
	var tmp Tmp
	row := db.QueryRow(sqlStatement, 3)
	err = row.Scan(&tmp.id, &tmp.name)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return
	case nil:
		fmt.Println(tmp.id)
		fmt.Println(tmp.name)
		fmt.Println(tmp)
	default:
		panic(err)
	}
}
