package database

import (
	"database/sql"
	"log"
)

func createTableRun(db *sql.DB) {
	createRunTableSQL := `CREATE TABLE run (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"name" TEXT,
		"description" TEXT,
		"scenarioId" integer,
		"listOfSteps" TEXT,
		"logs" TEXT,
		"featureId" integer,
		"environmentId" integer,
		"contextId" integer,
		"userId" integer,
		"tags" TEXT,
		"runAt" TEXT
	  );`

	log.Println("Create run table...")
	statement, err := db.Prepare(createRunTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("run table created")
}

func insertRuns() {
	insertRun(Db, 1, "Run 1 Test 1", "")
}

// We are passing db reference connection from main to our method with other parameters
func insertRun(db *sql.DB, id int, name, description string) {
	log.Println("Inserting run record ...")
	insertRunSQL := `INSERT INTO run(id, name, description) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertRunSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(id, name, description)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func displayRuns(db *sql.DB) {
	row, err := db.Query("SELECT id, name, description FROM run ORDER BY name")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var name string
		var description string
		row.Scan(&id, &name, &description)
		log.Println("Run: ", name, description)
	}
}
