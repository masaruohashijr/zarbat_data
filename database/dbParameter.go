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
	insertParameterContext(Db, 1, 1, 1, "https://api.zang.io", 1)
	insertParameter(Db, 2, "BaseUrl", "")
	insertParameterContext(Db, 2, 1, 2, "http://zarbat.ngrok.io", 2)
	insertParameter(Db, 3, "ApiVersion", "")
	insertParameterContext(Db, 3, 1, 3, "v2", 3)
	insertParameter(Db, 4, "AccountSid", "")
	insertParameterContext(Db, 4, 1, 4, "AC777c3e3262352901b0e24b7092f6eef0", 4)
	insertParameter(Db, 5, "AuthToken", "")
	insertParameterContext(Db, 5, 1, 5, "c59b774e18b44769bdc8b4253a7ee0d7", 5)
	insertParameter(Db, 6, "From", "")
	insertParameter(Db, 7, "To", "")
	insertParameter(Db, 8, "ToSid", "")
	insertParameter(Db, 9, "ActionUrl", "")
	insertParameterContext(Db, 6, 1, 6, "http://zang.io/ivr/welcome/call", 6)
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

func insertParameterContext(db *sql.DB, id, contextId, parameterId int, value string, position int) {
	log.Println("Inserting parameter record ...")
	insertParameterSQL := `INSERT INTO parameterContext(id, contextId, parameterId, value, position) VALUES (?, ?, ?, ?, ?)`
	statement, err := db.Prepare(insertParameterSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(id, contextId, parameterId, value, position)
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
