package testCase

import (
	"fmt"
	"log"
	"zarbat_mock/database"
)

func DbGetTestCases() (testCases []TestCase) {
	db := database.Db
	row, err := db.Query("SELECT * FROM testCase ORDER BY ID")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var testCase TestCase
	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&testCase.Id, &testCase.Name, &testCase.Description)
		testCases = append(testCases, testCase)
	}
	return testCases
}
func DbGetTestCase(id int) (testCase TestCase) {
	db := database.Db
	stmt, err := db.Prepare("SELECT id, name, description FROM testCase WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := stmt.QueryRow(id)
	if err != nil {
		log.Fatal(err)
	}
	row.Scan(&testCase.Id, &testCase.Name, &testCase.Description)
	return testCase
}
func DbAddTestCase(testCase TestCase) TestCase {
	db := database.Db
	insert := "INSERT INTO testCase (name, description) values ( ?, ?) RETURNING id"
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatalln(err.Error())
	}
	id := 0
	err = stmt.QueryRow(testCase.Name, testCase.Description).Scan(&id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	testCase.Id = id
	return testCase
}
func DbUpdateTestCase(testCase TestCase) TestCase {
	db := database.Db
	update := "UPDATE testCase SET name = ?, description = ? WHERE id = ?"
	stmt, err := db.Prepare(update)
	if err != nil {
		log.Fatalln(err.Error())
	}
	res, err := stmt.Exec(testCase.Name, testCase.Description, testCase.Id)
	affect, err := res.RowsAffected()
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(affect)
	return DbGetTestCase(testCase.Id)
}
func DbDeleteTestCase(id int) TestCase {
	db := database.Db
	delete := "DELETE FROM testCase WHERE id = ?"
	stmt, err := db.Prepare(delete)
	if err != nil {
		log.Fatalln(err.Error())
	}
	res, err := stmt.Exec(id)
	affect, err := res.RowsAffected()
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(affect)
	return DbGetTestCase(id)
}
