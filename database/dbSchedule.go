package database

import (
	"database/sql"
	"log"
)

func createTableSchedule(db *sql.DB) {
	createScheduleTableSQL := `CREATE TABLE schedule (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"name" TEXT,
		"description" TEXT
	  );`

	log.Println("Create schedule table...")
	statement, err := db.Prepare(createScheduleTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("schedule table created")
}

func insertSchedules() {
	insertSchedule(Db, 1, "US Every Saturday 12:00 AM", "")
	insertSchedule(Db, 2, "SI Every Saturday 12:00 AM", "")
	insertSchedule(Db, 3, "CA Every Saturday 12:00 AM", "")
	insertSchedule(Db, 4, "IN Every Saturday 12:00 AM", "")
	insertSchedule(Db, 5, "EU Every Saturday 12:00 AM", "")
	insertSchedule(Db, 6, "SI Every Saturday 12:00 PM", "")
}

// We are passing db reference connection from main to our method with other parameters
func insertSchedule(db *sql.DB, id int, name, description string) {
	log.Println("Inserting schedule record ...")
	insertScheduleSQL := `INSERT INTO schedule(id, name, description) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertScheduleSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(id, name, description)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func displaySchedules(db *sql.DB) {
	row, err := db.Query("SELECT * FROM schedule ORDER BY name")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var name string
		var description string
		row.Scan(&id, &name, &description)
		log.Println("Schedule: ", name, description)
	}
}
