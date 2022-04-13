package database

import (
	"database/sql"
	"log"
)

func createTableNumber(db *sql.DB) {
	createNumberTableSQL := `CREATE TABLE number (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"phoneNumber" TEXT,
		"sid" TEXT,
		"description" TEXT,
		"environmentId" integer
	  );`

	log.Println("Create number table...")
	statement, err := db.Prepare(createNumberTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("number table created")
}

func insertNumbers() {
	insertNumber(Db, 1, "+1 204-800-3082", "DI777c3e320f2f657cfb9d418ca591a6a6", "")
	insertPhoneNumber(Db, 1, 1, 1, "NumberA", 1)
	insertNumber(Db, 2, "+1 732-401-9498", "DI777c3e3217428380ead64b358f23a6ce", "")
	insertPhoneNumber(Db, 2, 1, 2, "NumberB", 2)
	insertNumber(Db, 3, "+1 204-800-3047", "DI777c3e32697ba80d5a4f4082bd4dcff5", "")
	insertPhoneNumber(Db, 3, 1, 3, "NumberC", 3)
	insertNumber(Db, 4, "+1 204-800-3029", "DI777c3e3270908f97680344b5a3fea522", "")
	insertPhoneNumber(Db, 4, 1, 4, "NumberD", 4)
	insertNumber(Db, 5, "+1 204-800-3030", "DI777c3e32d3e6623dbbe740ea93fdebb4", "")
	insertPhoneNumber(Db, 5, 1, 5, "NumberE", 5)
}

// We are passing db reference connection from main to our method with other parameters
func insertNumber(db *sql.DB, id int, phoneNumber, sid string, description string) {
	log.Println("Inserting number record ...")
	insertNumberSQL := `INSERT INTO number(id, phoneNumber, sid, description) VALUES (?, ?, ?, ?)`

	statement, err := db.Prepare(insertNumberSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(id, phoneNumber, sid, description)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func insertPhoneNumber(db *sql.DB, id, contextId, phoneNumberId int, alias string, position int) {
	insertNumberContextSQL := "INSERT INTO phoneNumberContext(id, contextId, phoneNumberId, alias, position) VALUES (?,?,?,?,?)"
	statement, err := db.Prepare(insertNumberContextSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(id, contextId, phoneNumberId, alias, position)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func displayNumbers(db *sql.DB) {
	row, err := db.Query("SELECT id, phoneNumber, sid, description FROM number ORDER BY phoneNumber")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var phoneNumber string
		var sid string
		var description string
		row.Scan(&id, &phoneNumber, &sid, &description)
		log.Println("Number: ", phoneNumber, sid, description)
	}
}
