package database

import (
	"database/sql"
	"log"
)

func createTableFeature(db *sql.DB) {
	createFeatureTableSQL := `CREATE TABLE feature (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"name" TEXT,
		"description" TEXT
	  );`

	log.Println("Create feature table...")
	statement, err := db.Prepare(createFeatureTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("feature table created")
}

func insertFeatures() {
	insertFeature(Db, 1, "Conference", "Conference")
	insertFeature(Db, 2, "Dial", "Dial")
	insertFeature(Db, 3, "Gather", "Gather")
	insertFeature(Db, 4, "Hangup", "Hangup")
	insertFeature(Db, 5, "Mms", "Mms")
	insertFeature(Db, 6, "Number", "Number")
	insertFeature(Db, 7, "Pause", "Pause")
	insertFeature(Db, 8, "Ping", "Ping")
	insertFeature(Db, 9, "Play", "Play")
	insertFeature(Db, 10, "PlayLastRecording", "PlayLastRecording")
	insertFeature(Db, 11, "Record", "Record")
	insertFeature(Db, 12, "Redirect", "Redirect")
	insertFeature(Db, 13, "Reject", "Reject")
	insertFeature(Db, 14, "Say", "Say")
	insertFeature(Db, 15, "Sip", "Sip")
	insertFeature(Db, 16, "Sms", "Sms")
}

// We are passing db reference connection from main to our method with other parameters
func insertFeature(db *sql.DB, id int, name, description string) {
	log.Println("Inserting feature record ...")
	insertFeatureSQL := `INSERT INTO feature(id, name, description) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertFeatureSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(id, name, description)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func displayFeatures(db *sql.DB) {
	row, err := db.Query("SELECT * FROM feature ORDER BY name")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var name string
		var description string
		row.Scan(&id, &name, &description)
		log.Println("Feature: ", name, description)
	}
}
