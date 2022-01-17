package database

import (
	"database/sql"
	"log"
)

func createTableScenario(db *sql.DB) {
	createScenarioTableSQL := `CREATE TABLE scenario (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"name" TEXT,
		"description" TEXT,
		"story" TEXT,
		"featureId" integer
	  );`

	log.Println("Create scenario table...")
	statement, err := db.Prepare(createScenarioTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("scenario table created")
}

func insertScenarios() {
	insertScenario(Db, 1, "Make Call", "", "", 0)
	insertScenario(Db, 2, "View Call", "", "", 0)
	insertScenario(Db, 3, "List Calls", "", "", 0)
	insertScenario(Db, 4, "Interrupt Live Call", "", "", 0)
	insertScenario(Db, 5, "Send Digits to Live Call", "", "", 0)
	insertScenario(Db, 6, "Play Audio to Live Call", "", "", 0)
	insertScenario(Db, 7, "Apply Voice Effect", "", "", 0)
	insertScenario(Db, 8, "View Sms", "", "", 0)
	insertScenario(Db, 9, "List Sms", "", "", 0)
	insertScenario(Db, 10, "Send Sms", "", "", 0)
}

// We are passing db reference connection from main to our method with other parameters
func insertScenario(db *sql.DB, id int, name, description, story string, featureId int) {
	log.Println("Inserting scenario record ...")
	insertScenarioSQL := `INSERT INTO scenario(id, name, description, story, featureId) VALUES (?, ?, ?, ?, ?)`
	statement, err := db.Prepare(insertScenarioSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(id, name, description, story, featureId)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func displayScenarios(db *sql.DB) {
	row, err := db.Query("SELECT * FROM scenario ORDER BY name")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var name string
		var description string
		var story string
		var featureId int
		row.Scan(&id, &name, &description, &story, &featureId)
		log.Println("Scenario: ", name, description, story, featureId)
	}
}
