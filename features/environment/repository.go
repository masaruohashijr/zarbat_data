package environment

import (
	"fmt"
	"log"
	"zarbat_data/database"
)

func DbGetEnvironments() (environments []Environment) {
	db := database.Db
	row, err := db.Query("SELECT * FROM environment ORDER BY ID")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var environment Environment
	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&environment.Id, &environment.Name, &environment.Description)
		environments = append(environments, environment)
	}
	return environments
}
func DbGetEnvironment(id int) (environment Environment) {
	db := database.Db
	stmt, err := db.Prepare("SELECT id, name, description FROM environment WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := stmt.QueryRow(id)
	if err != nil {
		log.Fatal(err)
	}
	row.Scan(&environment.Id, &environment.Name, &environment.Description)
	return environment
}
func DbAddEnvironment(environment Environment) Environment {
	db := database.Db
	insert := "INSERT INTO environment (name, description) values ( ?, ?) RETURNING id"
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatalln(err.Error())
	}
	id := 0
	err = stmt.QueryRow(environment.Name, environment.Description).Scan(&id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	environment.Id = id
	return environment
}
func DbUpdateEnvironment(environment Environment) Environment {
	db := database.Db
	update := "UPDATE environment SET name = ?, description = ? WHERE id = ?"
	stmt, err := db.Prepare(update)
	if err != nil {
		log.Fatalln(err.Error())
	}
	res, err := stmt.Exec(environment.Name, environment.Description, environment.Id)
	affect, err := res.RowsAffected()
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(affect)
	return DbGetEnvironment(environment.Id)
}
func DbDeleteEnvironment(id int) Environment {
	db := database.Db
	delete := "DELETE FROM environment WHERE id = ?"
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
	return DbGetEnvironment(id)
}
