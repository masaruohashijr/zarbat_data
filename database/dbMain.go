package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

var Db *sql.DB

func InitDatabase() {
	resetDB()
	Db, _ = sql.Open("sqlite3", "./sqlite-database.db")
	// User
	createTableUser(Db)
	insertUsers()
	displayUsers(Db)
	// Context
	createTableContext(Db)
	insertContexts()
	displayContexts(Db)
	// Environment
	createTableEnvironment(Db)
	insertEnvironments()
	displayEnvironments(Db)
	// Feature
	createTableFeature(Db)
	insertFeatures()
	displayFeatures(Db)
	// Number
	createTableNumber(Db)
	insertNumbers()
	displayNumbers(Db)
	// Parameter
	createTableParameter(Db)
	insertParameters()
	displayParameters(Db)
	// Run
	createTableRun(Db)
	insertRuns()
	displayRuns(Db)
	// Scenario
	createTableScenario(Db)
	insertScenarios()
	displayScenarios(Db)
	// Schedule
	createTableSchedule(Db)
	insertSchedules()
	displaySchedules(Db)
	// Step
	createTableStep(Db)
	insertSteps()
	displaySteps(Db)
	// TestCase
	createTableTestCase(Db)
	insertTestCases()
	displayTestCases(Db)
}

func resetDB() {
	os.Remove("sqlite-database.db") // I delete the file to avoid duplicated records.
	// SQLite is a file based database.
	log.Println("Creating sqlite-database.db...")
	file, err := os.Create("sqlite-database.db") // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("sqlite-database.db created")
}
