package parameter

import (
	"fmt"
	"log"
	"zarbat_data/database"
)

func DbGetParameters() (parameters []Parameter) {
	db := database.Db
	row, err := db.Query("SELECT * FROM parameter ORDER BY ID")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var parameter Parameter
	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&parameter.Id, &parameter.Name, &parameter.Description)
		parameters = append(parameters, parameter)
	}
	return parameters
}
func DbGetParameter(id int) (parameter Parameter) {
	db := database.Db
	stmt, err := db.Prepare("SELECT id, name, description FROM parameter WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := stmt.QueryRow(id)
	if err != nil {
		log.Fatal(err)
	}
	row.Scan(&parameter.Id, &parameter.Name, &parameter.Description)
	return parameter
}
func DbAddParameter(parameter Parameter) Parameter {
	db := database.Db
	insert := "INSERT INTO parameter (name, description) values ( ?, ?) RETURNING id"
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatalln(err.Error())
	}
	id := 0
	err = stmt.QueryRow(parameter.Name, parameter.Description).Scan(&id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	parameter.Id = id
	return parameter
}
func DbUpdateParameter(parameter Parameter) Parameter {
	db := database.Db
	update := "UPDATE parameter SET name = ?, description = ? WHERE id = ?"
	stmt, err := db.Prepare(update)
	if err != nil {
		log.Fatalln(err.Error())
	}
	res, err := stmt.Exec(parameter.Name, parameter.Description, parameter.Id)
	affect, err := res.RowsAffected()
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(affect)
	return DbGetParameter(parameter.Id)
}
func DbDeleteParameter(id int) Parameter {
	db := database.Db
	delete := "DELETE FROM parameter WHERE id = ?"
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
	return DbGetParameter(id)
}
