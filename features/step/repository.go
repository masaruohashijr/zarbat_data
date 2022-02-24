package step

import (
	"fmt"
	"log"
	"zarbat_data/database"
)

func DbGetSteps() (steps []Step) {
	db := database.Db
	row, err := db.Query("SELECT * FROM step ORDER BY ID")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var step Step
	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&step.Id, &step.Name, &step.Description)
		steps = append(steps, step)
	}
	return steps
}
func DbGetStep(id int) (step Step) {
	db := database.Db
	stmt, err := db.Prepare("SELECT id, name, description FROM step WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := stmt.QueryRow(id)
	if err != nil {
		log.Fatal(err)
	}
	row.Scan(&step.Id, &step.Name, &step.Description)
	return step
}
func DbAddStep(step Step) Step {
	db := database.Db
	insert := "INSERT INTO step (name, description) values ( ?, ?) RETURNING id"
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatalln(err.Error())
	}
	id := 0
	err = stmt.QueryRow(step.Name, step.Description).Scan(&id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	step.Id = id
	return step
}
func DbUpdateStep(step Step) Step {
	db := database.Db
	update := "UPDATE step SET name = ?, description = ? WHERE id = ?"
	stmt, err := db.Prepare(update)
	if err != nil {
		log.Fatalln(err.Error())
	}
	res, err := stmt.Exec(step.Name, step.Description, step.Id)
	affect, err := res.RowsAffected()
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(affect)
	return DbGetStep(step.Id)
}
func DbDeleteStep(id int) Step {
	db := database.Db
	delete := "DELETE FROM step WHERE id = ?"
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
	return DbGetStep(id)
}
