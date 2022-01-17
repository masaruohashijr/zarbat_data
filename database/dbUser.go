package database

import (
	"database/sql"
	"log"
)

func createTableUser(db *sql.DB) {
	createUserTableSQL := `CREATE TABLE user (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"name" TEXT,
		"description" TEXT
	  );`

	log.Println("Create user table...")
	statement, err := db.Prepare(createUserTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("user table created")
}

func insertUsers() {
	insertUser(Db, 1, "Jignesh Vasoya", "Team Leader")
	insertUser(Db, 2, "Masoud Mazarei", "Tech Leader")
	insertUser(Db, 3, "Kanchan Mittal", "QA Analyst")
	insertUser(Db, 4, "Nithin", "QA Tester")
	insertUser(Db, 5, "Masaru Ohashi Jr", "GO Developer")
}

// We are passing db reference connection from main to our method with other parameters
func insertUser(db *sql.DB, id int, name, description string) {
	log.Println("Inserting user record ...")
	insertUserSQL := `INSERT INTO user(id, name, description) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertUserSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(id, name, description)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func displayUsers(db *sql.DB) {
	row, err := db.Query("SELECT * FROM user ORDER BY name")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var name string
		var description string
		row.Scan(&id, &name, &description)
		log.Println("User: ", name, description)
	}
}
