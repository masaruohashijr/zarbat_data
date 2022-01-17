package number

import (
	"fmt"
	"log"
	"zarbat_mock/database"
)

func DbGetNumbers() (numbers []Number) {
	db := database.Db
	row, err := db.Query("SELECT * FROM number ORDER BY ID")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var number Number
	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&number.Id, &number.Name, &number.Description)
		numbers = append(numbers, number)
	}
	return numbers
}
func DbGetNumber(id int) (number Number) {
	db := database.Db
	stmt, err := db.Prepare("SELECT id, name, description FROM number WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := stmt.QueryRow(id)
	if err != nil {
		log.Fatal(err)
	}
	row.Scan(&number.Id, &number.Name, &number.Description)
	return number
}
func DbAddNumber(number Number) Number {
	db := database.Db
	insert := "INSERT INTO number (name, description) values ( ?, ?) RETURNING id"
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatalln(err.Error())
	}
	id := 0
	err = stmt.QueryRow(number.Name, number.Description).Scan(&id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	number.Id = id
	return number
}
func DbUpdateNumber(number Number) Number {
	db := database.Db
	update := "UPDATE number SET name = ?, description = ? WHERE id = ?"
	stmt, err := db.Prepare(update)
	if err != nil {
		log.Fatalln(err.Error())
	}
	res, err := stmt.Exec(number.Name, number.Description, number.Id)
	affect, err := res.RowsAffected()
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(affect)
	return DbGetNumber(number.Id)
}
func DbDeleteNumber(id int) Number {
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
