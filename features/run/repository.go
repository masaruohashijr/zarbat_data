package run

import (
	"fmt"
	"log"
	"zarbat_mock/database"
)

func DbGetRuns() (runs []Run) {
	db := database.Db
	row, err := db.Query("SELECT * FROM run ORDER BY ID")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var run Run
	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&run.Id, &run.Name, &run.Description)
		runs = append(runs, run)
	}
	return runs
}
func DbGetRun(id int) (run Run) {
	db := database.Db
	stmt, err := db.Prepare("SELECT id, name, description FROM run WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := stmt.QueryRow(id)
	if err != nil {
		log.Fatal(err)
	}
	row.Scan(&run.Id, &run.Name, &run.Description)
	return run
}
func DbAddRun(run Run) Run {
	db := database.Db
	insert := "INSERT INTO run (name, description) values ( ?, ?) RETURNING id"
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatalln(err.Error())
	}
	id := 0
	err = stmt.QueryRow(run.Name, run.Description).Scan(&id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	run.Id = id
	return run
}
func DbUpdateRun(run Run) Run {
	db := database.Db
	update := "UPDATE run SET name = ?, description = ? WHERE id = ?"
	stmt, err := db.Prepare(update)
	if err != nil {
		log.Fatalln(err.Error())
	}
	res, err := stmt.Exec(run.Name, run.Description, run.Id)
	affect, err := res.RowsAffected()
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(affect)
	return DbGetRun(run.Id)
}
func DbDeleteRun(id int) Run {
	db := database.Db
	delete := "DELETE FROM run WHERE id = ?"
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
	return DbGetRun(id)
}
