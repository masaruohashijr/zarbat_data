package database

import (
	"database/sql"
	"log"
)

func createTableTestCase(db *sql.DB) {
	createTestCaseTableSQL := `CREATE TABLE testCase (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"name" TEXT,
		"description" TEXT
	  );`

	log.Println("Create testCase table...")
	statement, err := db.Prepare(createTestCaseTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("testCase table created")
}

func insertTestCases() {
	insertTestCase(Db, 1, "Test 1", "")
	insertTestCase(Db, 2, "Test 2", "")
	insertTestCase(Db, 3, "Test 3", "")
}

// We are passing db reference connection from main to our method with other parameters
func insertTestCase(db *sql.DB, id int, name, description string) {
	log.Println("Inserting testCase record ...")
	insertTestCaseSQL := `INSERT INTO testCase(id, name, description) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertTestCaseSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(id, name, description)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func displayTestCases(db *sql.DB) {
	row, err := db.Query("SELECT * FROM testCase ORDER BY name")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var name string
		var description string
		row.Scan(&id, &name, &description)
		log.Println("TestCase: ", name, description)
	}
}
