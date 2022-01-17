package database

import (
	"database/sql"
	"log"
)

func createTableStep(db *sql.DB) {
	createStepTableSQL := `CREATE TABLE step (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"name" TEXT,
		"description" TEXT
	  );`

	log.Println("Create step table...")
	statement, err := db.Prepare(createStepTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("step table created")
}

func insertSteps() {
	insertStep(Db, 1, "Given my test setup runs", "")
	insertStep(Db, 2, "Given \"NumberB\" configured to dial \"NumberC\"", "")
	insertStep(Db, 3, "When I make a call from \"NumberA\" to \"NumberB\"", "")
	insertStep(Db, 4, "Then \"NumberC\" should get the incoming call from \"NumberB\"", "")
}

// We are passing db reference connection from main to our method with other parameters
func insertStep(db *sql.DB, id int, name, description string) {
	log.Println("Inserting step record ...")
	insertStepSQL := `INSERT INTO step(id, name, description) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertStepSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(id, name, description)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func displaySteps(db *sql.DB) {
	row, err := db.Query("SELECT * FROM step ORDER BY name")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var name string
		var description string
		row.Scan(&id, &name, &description)
		log.Println("Step: ", name, description)
	}
}
