package database

import (
	"database/sql"
	"log"
)

func createTableNumber(db *sql.DB) {
	createNumberTableSQL := `CREATE TABLE number (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"name" TEXT,
		"description" TEXT
	  );`

	log.Println("Create number table...")
	statement, err := db.Prepare(createNumberTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("number table created")
}

func insertNumbers() {
	insertNumber(Db, 1, "+1 204-800-3082", "")
	insertNumber(Db, 2, "+1 732-401-9498", "")
	insertNumber(Db, 3, "+1 204-800-3047", "")
	insertNumber(Db, 4, "+1 204-800-3029", "")
	insertNumber(Db, 5, "+1 204-800-3030", "")
}

// We are passing db reference connection from main to our method with other parameters
func insertNumber(db *sql.DB, id int, name, description string) {
	log.Println("Inserting number record ...")
	insertNumberSQL := `INSERT INTO number(id, name, description) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertNumberSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(id, name, description)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func displayNumbers(db *sql.DB) {
	row, err := db.Query("SELECT * FROM number ORDER BY name")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var name string
		var description string
		row.Scan(&id, &name, &description)
		log.Println("Number: ", name, description)
	}
}
