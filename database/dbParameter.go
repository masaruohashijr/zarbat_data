package database

import (
	"database/sql"
	"log"
)

func createTableParameter(db *sql.DB) {
	createParameterTableSQL := `CREATE TABLE parameter (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"name" TEXT,
		"description" TEXT
	  );`

	log.Println("Create parameter table...")
	statement, err := db.Prepare(createParameterTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("parameter table created")
}

func insertParameters() {
	insertParameter(Db, 1, "ApiUrl", "")
	insertParameter(Db, 2, "BaseUrl", "")
	insertParameter(Db, 3, "ApiVersion", "")
	insertParameter(Db, 4, "AccountSid", "")
	insertParameter(Db, 5, "AuthToken", "")
	insertParameter(Db, 6, "From", "")
	insertParameter(Db, 7, "To", "")
	insertParameter(Db, 8, "ToSid", "")
	insertParameter(Db, 9, "ActionUrl", "")
	insertParameter(Db, 10, "StatusCallback", "")
}

// We are passing db reference connection from main to our method with other parameters
func insertParameter(db *sql.DB, id int, name, description string) {
	log.Println("Inserting parameter record ...")
	insertParameterSQL := `INSERT INTO parameter(id, name, description) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertParameterSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(id, name, description)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func displayParameters(db *sql.DB) {
	row, err := db.Query("SELECT * FROM parameter ORDER BY name")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var name string
		var description string
		row.Scan(&id, &name, &description)
		log.Println("Parameter: ", name, description)
	}
}
