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
		"listOfSteps" TEXT,
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
	insertScenario(Db, 1, "Make Call", "", " Given my test setup runs\n"+
		" And \"NumberB\" configured to hang up after 2 seconds\n"+
		" When I make a call from \"NumberA\" to \"NumberB\"\n"+
		" And After waiting for 10 seconds\n"+
		" Then I should get to view a call from \"NumberA\" to \"NumberB\" with status \"completed\"\n", 0)
	insertScenario(Db, 2, "View Call", "", "Step 1\nStep 2\nStep 3\nStep 4\nStep 5", 0)
	insertScenario(Db, 3, "List Calls", "", "Step 1\nStep 2\nStep 3\nStep 4\nStep 5", 0)
	insertScenario(Db, 4, "Interrupt Live Call", "", "Step 1\nStep 2\nStep 3\nStep 4\nStep 5", 0)
	insertScenario(Db, 5, "Send Digits to Live Call", "", "Step 1\nStep 2\nStep 3\nStep 4\nStep 5", 0)
	insertScenario(Db, 6, "Play Audio to Live Call", "", "Step 1\nStep 2\nStep 3\nStep 4\nStep 5", 0)
	insertScenario(Db, 7, "Apply Voice Effect", "", "Step 1\nStep 2\nStep 3\nStep 4\nStep 5", 0)
	insertScenario(Db, 8, "View Sms", "", "Step 1\nStep 2\nStep 3\nStep 4\nStep 5", 0)
	insertScenario(Db, 9, "List Sms", "", "Step 1\nStep 2\nStep 3\nStep 4\nStep 5", 0)
	insertScenario(Db, 10, "Send Sms", "", "Step 1\nStep 2\nStep 3\nStep 4\nStep 5", 0)
}

// We are passing db reference connection from main to our method with other parameters
func insertScenario(db *sql.DB, id int, name, description, listOfSteps string, featureId int) {
	log.Println("Inserting scenario record ...")
	insertScenarioSQL := `INSERT INTO scenario(id, name, description, listOfSteps, featureId) VALUES (?, ?, ?, ?, ?)`
	statement, err := db.Prepare(insertScenarioSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(id, name, description, listOfSteps, featureId)
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
		var listOfSteps string
		var featureId int
		row.Scan(&id, &name, &description, &listOfSteps, &featureId)
		log.Println("Scenario: ", name, description, listOfSteps, featureId)
	}
}
