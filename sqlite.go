package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Pessoa struct {
	id    int32
	nome  string
	idade uint8
}

func SQLiteTest() {
	database, _ := sql.Open("sqlite3", "./sqlite3.db")

	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS Pessoa (ID INTEGER PRIMARY KEY, nome TEXT, idade INTEGER)")
	statement.Exec()

	statement2, _ := database.Prepare("INSERT INTO Pessoa (nome, idade) VALUES (?, ?)")
	statement2.Exec("luiz", 19)

	rows, _ := database.Query("SELECT * FROM Pessoa")
	var p Pessoa
	for rows.Next() {
		rows.Scan(&p.id, &p.nome, &p.idade)
	}

	fmt.Printf("%v\n", p)

}
