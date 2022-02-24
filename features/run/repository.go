package run

import (
	"fmt"
	"log"
	"zarbat_data/database"
)

func DbGetRuns() (runs []Run) {
	db := database.Db
	row, err := db.Query("SELECT id, name, description, listOfSteps, logs, featureId, environmentId, contextId, userId, tags, runAt FROM run ORDER BY ID")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var run Run
	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&run.Id, &run.Name, &run.Description, &run.ListOfSteps, &run.Logs, &run.FeatureId, &run.EnvironmentId, &run.ContextId, &run.Tags, &run.RunAt, &run.UserId)
		runs = append(runs, run)
	}
	return runs
}
func DbGetRun(id int) (run Run) {
	db := database.Db
	stmt, err := db.Prepare("SELECT id, name, description, scenarioId, listOfSteps, logs, featureId, environmentId, contextId, userId, tags, runAt FROM run WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := stmt.QueryRow(id)
	if err != nil {
		log.Fatal(err)
	}
	row.Scan(&run.Id, &run.Name, &run.Description, &run.ScenarioId, &run.ListOfSteps, &run.Logs, &run.FeatureId, &run.EnvironmentId, &run.ContextId, &run.UserId, &run.Tags, &run.RunAt)
	return run
}
func DbAddRun(run Run) Run {
	db := database.Db
	insert := "INSERT INTO run (name, description, listOfSteps, logs, scenarioId, featureId, environmentId, contextId, userId, tags, runAt) values ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) RETURNING id"
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatalln(err.Error())
	}
	id := 0
	err = stmt.QueryRow(run.Name, run.Description, run.ListOfSteps, run.Logs, run.ScenarioId, run.FeatureId, run.EnvironmentId, run.ContextId, run.UserId, run.Tags, run.RunAt).Scan(&id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	run.Id = id
	return run
}
func DbUpdateRun(run Run) Run {
	db := database.Db
	update := "UPDATE run SET name = ?, description = ?, scenarioId = ?, listOfSteps = ?, logs = ?, featureId = ?, environmentId = ?, contextId = ?, userId = ?, tags = ?, runAt =?  WHERE id = ?"
	stmt, err := db.Prepare(update)
	if err != nil {
		log.Fatalln(err.Error())
	}
	res, err := stmt.Exec(run.Name, run.Description, run.ScenarioId, run.ListOfSteps, run.Logs, run.FeatureId, run.EnvironmentId, run.ContextId, run.UserId, run.Tags, run.RunAt, run.Id)
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
