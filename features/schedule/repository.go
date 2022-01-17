package schedule

import (
	"fmt"
	"log"
	"zarbat_mock/database"
)

func DbGetSchedules() (schedules []Schedule) {
	db := database.Db
	row, err := db.Query("SELECT * FROM schedule ORDER BY ID")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var schedule Schedule
	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&schedule.Id, &schedule.Name, &schedule.Description)
		schedules = append(schedules, schedule)
	}
	return schedules
}
func DbGetSchedule(id int) (schedule Schedule) {
	db := database.Db
	stmt, err := db.Prepare("SELECT id, name, description FROM schedule WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := stmt.QueryRow(id)
	if err != nil {
		log.Fatal(err)
	}
	row.Scan(&schedule.Id, &schedule.Name, &schedule.Description)
	return schedule
}
func DbAddSchedule(schedule Schedule) Schedule {
	db := database.Db
	insert := "INSERT INTO schedule (name, description) values ( ?, ?) RETURNING id"
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatalln(err.Error())
	}
	id := 0
	err = stmt.QueryRow(schedule.Name, schedule.Description).Scan(&id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	schedule.Id = id
	return schedule
}
func DbUpdateSchedule(schedule Schedule) Schedule {
	db := database.Db
	update := "UPDATE schedule SET name = ?, description = ? WHERE id = ?"
	stmt, err := db.Prepare(update)
	if err != nil {
		log.Fatalln(err.Error())
	}
	res, err := stmt.Exec(schedule.Name, schedule.Description, schedule.Id)
	affect, err := res.RowsAffected()
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(affect)
	return DbGetSchedule(schedule.Id)
}
func DbDeleteSchedule(id int) Schedule {
	db := database.Db
	delete := "DELETE FROM schedule WHERE id = ?"
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
	return DbGetSchedule(id)
}
