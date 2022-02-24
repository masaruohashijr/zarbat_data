package number

import (
	"fmt"
	"log"
	"zarbat_data/database"
)

func DbGetNumbers() (numbers []PhoneNumber) {
	db := database.Db
	row, err := db.Query("SELECT id, phoneNumber, sid, description, environmentId FROM number ORDER BY ID")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var number PhoneNumber
	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&number.Id, &number.PhoneNumber, &number.Sid, &number.Description, &number.EnvironmentId)
		numbers = append(numbers, number)
	}
	return numbers
}
func DbGetNumber(id int) (number PhoneNumber) {
	db := database.Db
	stmt, err := db.Prepare("SELECT id, phoneNumber, sid, description, environmentId FROM number WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := stmt.QueryRow(id)
	if err != nil {
		log.Fatal(err)
	}
	row.Scan(&number.Id, &number.PhoneNumber, &number.Sid, &number.Description, &number.EnvironmentId)
	return number
}
func DbAddNumber(number PhoneNumber) PhoneNumber {
	db := database.Db
	insert := "INSERT INTO number (phoneNumber, sid, description, environmentId) values ( ?, ?, ?, ?) RETURNING id"
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatalln(err.Error())
	}
	id := 0
	err = stmt.QueryRow(number.PhoneNumber, number.Sid, number.Description, number.EnvironmentId).Scan(&id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	number.Id = id
	return number
}
func DbUpdateNumber(number PhoneNumber) PhoneNumber {
	db := database.Db
	update := "UPDATE number SET phoneNumber = ?, sid = ?, description = ?, environmentId = ? WHERE id = ?"
	stmt, err := db.Prepare(update)
	if err != nil {
		log.Fatalln(err.Error())
	}
	res, err := stmt.Exec(number.PhoneNumber, number.Sid, number.Description, number.EnvironmentId, number.Id)
	affect, err := res.RowsAffected()
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(affect)
	return DbGetNumber(number.Id)
}
func DbDeleteNumber(id int) PhoneNumber {
	db := database.Db
	delete := "DELETE FROM number WHERE id = ?"
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
	return DbGetNumber(id)
}
