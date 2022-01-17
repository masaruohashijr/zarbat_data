package database

import (
	"database/sql"
	"log"
)

func createTableContext(db *sql.DB) {
	createContextTableSQL := `CREATE TABLE context (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"name" TEXT,
		"description" TEXT,
		"environmentId" integer
	  );`

	log.Println("Create context table...")
	statement, err := db.Prepare(createContextTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("context table created")
}

func insertContexts() {
	insertContext(Db, 1, "ALFA", "Context ALFA", 1)
	insertContext(Db, 2, "BETA", "Context BETA", 1)
	insertContext(Db, 3, "GAMA", "Context GAMA", 1)
	insertContext(Db, 4, "DELTA", "Context DELTA", 1)
	insertContext(Db, 5, "EPSILON", "Context EPSILON", 1)
	insertContext(Db, 6, "ZETA", "Context ZETA", 1)
}

// We are passing db reference connection from main to our method with other parameters
func insertContext(db *sql.DB, id int, name, description string, environmentId int) {
	log.Println("Inserting context record ...")
	insertContextSQL := `INSERT INTO context(id, name, description, environmentId) VALUES (?, ?, ?, ?)`
	statement, err := db.Prepare(insertContextSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(id, name, description, environmentId)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func displayContexts(db *sql.DB) {
	row, err := db.Query("SELECT * FROM context ORDER BY name")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var name string
		var description string
		var envioronmentId int
		row.Scan(&id, &name, &description, &envioronmentId)
		log.Println("Context: ", name, description, envioronmentId)
	}
}
