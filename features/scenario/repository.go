package scenario

import (
	"fmt"
	"log"
	"zarbat_mock/database"
)

func DbGetScenarios() (scenarios []Scenario) {
	db := database.Db
	row, err := db.Query("SELECT * FROM scenario ORDER BY ID")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var scenario Scenario
	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&scenario.Id, &scenario.Name, &scenario.Description, &scenario.Story, &scenario.FeatureId)
		scenarios = append(scenarios, scenario)
	}
	return scenarios
}
func DbGetScenario(id int) (scenario Scenario) {
	db := database.Db
	stmt, err := db.Prepare("SELECT id, name, description, story, featureId FROM scenario WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := stmt.QueryRow(id)
	if err != nil {
		log.Fatal(err)
	}
	row.Scan(&scenario.Id, &scenario.Name, &scenario.Description, &scenario.Story, &scenario.FeatureId)
	return scenario
}
func DbAddScenario(scenario Scenario) Scenario {
	db := database.Db
	insert := "INSERT INTO scenario (name, description, story, featureId) values ( ?, ?, ?, ?) RETURNING id"
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatalln(err.Error())
	}
	id := 0
	err = stmt.QueryRow(scenario.Name, scenario.Description, scenario.Story, scenario.FeatureId).Scan(&id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	scenario.Id = id
	return scenario
}
func DbUpdateScenario(scenario Scenario) Scenario {
	db := database.Db
	update := "UPDATE scenario SET name = ?, description = ?, story = ?, featureId = ? WHERE id = ?"
	stmt, err := db.Prepare(update)
	if err != nil {
		log.Fatalln(err.Error())
	}
	res, err := stmt.Exec(scenario.Name, scenario.Description, scenario.Story, scenario.FeatureId, scenario.Id)
	affect, err := res.RowsAffected()
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(affect)
	return DbGetScenario(scenario.Id)
}
func DbDeleteScenario(id int) Scenario {
	db := database.Db
	delete := "DELETE FROM scenario WHERE id = ?"
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
	return DbGetScenario(id)
}
