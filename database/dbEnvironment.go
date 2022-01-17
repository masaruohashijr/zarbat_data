package database

import (
	"database/sql"
	"log"
)

func createTableEnvironment(db *sql.DB) {
	createEnvironmentTableSQL := `CREATE TABLE environment (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"name" TEXT,
		"description" TEXT
	  );`

	log.Println("Create environment table...")
	statement, err := db.Prepare(createEnvironmentTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("environment table created")
}

func insertEnvironments() {
	insertEnvironment(Db, 1, "Singapore", "")
	insertEnvironment(Db, 2, "Indonesia", "")
	insertEnvironment(Db, 3, "Europe", "")
	insertEnvironment(Db, 4, "Canada", "")
	insertEnvironment(Db, 5, "US", "")

}

// We are passing db reference connection from main to our method with other parameters
func insertEnvironment(db *sql.DB, id int, name, description string) {
	log.Println("Inserting environment record ...")
	insertEnvironmentSQL := `INSERT INTO environment(id, name, description) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertEnvironmentSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(id, name, description)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func displayEnvironments(db *sql.DB) {
	row, err := db.Query("SELECT * FROM environment ORDER BY name")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var name string
		var description string
		row.Scan(&id, &name, &description)
		log.Println("Environment: ", name, description)
	}
}
