package testCase

import (
	"fmt"
	"log"
	"zarbat_data/database"
	s "zarbat_data/features/scenario"
)

func DbGetTestCases() (testCases []TestCase) {
	db := database.Db
	row, err := db.Query("SELECT id, name, description, environmentId, contextId FROM testCase ORDER BY ID DESC")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var testCase TestCase
	for row.Next() {
		row.Scan(&testCase.Id, &testCase.Name, &testCase.Description, &testCase.EnvironmentId, &testCase.ContextId)
		testCases = append(testCases, testCase)
		testCase = DbGetScenarioTestCase(testCase)
	}
	return testCases
}
func DbGetTestCase(id int) (testCase TestCase) {
	db := database.Db
	stmt, err := db.Prepare("SELECT id, name, description, environmentId, contextId FROM testCase WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	row := stmt.QueryRow(id)
	if err != nil {
		log.Fatal(err)
	}
	row.Scan(&testCase.Id, &testCase.Name, &testCase.Description, &testCase.EnvironmentId, &testCase.ContextId)
	testCase = DbGetScenarioTestCase(testCase)
	return testCase
}

func DbGetScenarioTestCase(testCase TestCase) TestCase {
	db := database.Db
	var scenario s.Scenario
	var scenarios []s.Scenario
	stmt, _ := db.Prepare("SELECT b.id, b.name, b.description, b.listOfSteps, b.featureId, COALESCE(a.position,'') " +
		"FROM scenarioTestCase a INNER JOIN scenario b ON a.scenarioId = b.id WHERE a.testCaseId = ? ORDER BY position")
	rows, err := stmt.Query(testCase.Id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	for rows.Next() {
		err := rows.Scan(&scenario.Id, &scenario.Name, &scenario.Description, &scenario.ListOfSteps, &scenario.FeatureId, &scenario.Position)
		if err != nil {
			log.Fatalln(err.Error())
		}
		scenarios = append(scenarios, scenario)
	}
	testCase.Scenarios = scenarios
	return testCase
}
func DbAddTestCase(testCase TestCase) TestCase {
	db := database.Db
	insert := "INSERT INTO testCase (name, description, environmentId, contextId) values ( ?, ?, ?, ?) RETURNING id"
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatalln(err.Error())
	}
	id := 0
	err = stmt.QueryRow(testCase.Name, testCase.Description, testCase.EnvironmentId, testCase.ContextId).Scan(&id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	testCase.Id = id
	DbAddScenariosTestCase(testCase)
	return testCase
}
func DbUpdateTestCase(testCase TestCase) TestCase {
	db := database.Db
	update := "UPDATE testCase SET name = ?, description = ?, environmentId = ?, contextId = ? WHERE id = ?"
	stmt, err := db.Prepare(update)
	if err != nil {
		log.Fatalln(err.Error())
	}
	res, err := stmt.Exec(testCase.Name, testCase.Description, testCase.EnvironmentId, testCase.ContextId, testCase.Id)
	affect, err := res.RowsAffected()
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(affect)
	DbUpdateScenariosTestCase(testCase)
	return DbGetTestCase(testCase.Id)
}
func DbDeleteTestCase(id int) TestCase {
	db := database.Db
	var testCase TestCase
	testCase.Id = id
	testCase = DbGetScenarioTestCase(testCase)
	deleteScenarios(testCase, testCase.Scenarios)
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

func DbAddScenariosTestCase(testCase TestCase) {
	db := database.Db
	insert := "INSERT INTO scenarioTestcase (testCaseId, scenarioId, position) values ( ?, ?, ?) RETURNING id"
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatalln(err.Error())
	}
	id := 0
	for index, scenario := range testCase.Scenarios {
		err := stmt.QueryRow(testCase.Id, scenario.Id, index).Scan(&id)
		if err != nil {
			log.Fatalln(err.Error())
		}
		println(insert, id, testCase.Id, scenario.Id, scenario.Position)
	}
}

func DbUpdateScenariosTestCase(testCase TestCase) {
	db := database.Db
	stmt, err := db.Prepare("SELECT b.id, b.name, b.description, b.listOfSteps, b.featureId, a.position " +
		"FROM scenarioTestCase a INNER JOIN scenario b ON a.scenarioId = b.id WHERE a.testCaseId = ? ORDER BY position")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := stmt.Query(testCase.Id)
	var scenario s.Scenario
	var scenariosDB []s.Scenario
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&scenario.Id, &scenario.Name, &scenario.Description, &scenario.ListOfSteps, &scenario.FeatureId, &scenario.Position)
		scenariosDB = append(scenariosDB, scenario)
	}
	copyTestCaseScenarios := make([]s.Scenario, len(testCase.Scenarios))
	copy(copyTestCaseScenarios, testCase.Scenarios)
	scenariosDB, scenariosPage := diffScenarios(scenariosDB, copyTestCaseScenarios)
	if len(scenariosDB) > 0 {
		deleteScenarios(testCase, scenariosDB)
	}
	if len(scenariosPage) > 0 {
		addScenarios(testCase, scenariosPage)
	}
	update := "UPDATE scenarioTestCase SET scenarioId = ? WHERE testCaseId = ? AND position = ?"
	stmt, err = db.Prepare(update)
	if err != nil {
		log.Fatalln(err.Error())
	}
	for _, scenario := range testCase.Scenarios {
		res, err := stmt.Exec(scenario.Id, testCase.Id, scenario.Position)
		affect, err := res.RowsAffected()
		if err != nil {
			log.Fatalln(err.Error())
		}
		fmt.Println("Scenario ", scenario.Id, affect)
	}
}

func diffScenarios(scenariosDB, scenariosPage []s.Scenario) ([]s.Scenario, []s.Scenario) {
	for indexDB := 0; indexDB < len(scenariosDB); indexDB++ {
		sdb := scenariosDB[indexDB]
		for indexPage := 0; indexPage < len(scenariosPage); indexPage++ {
			sp := scenariosPage[indexPage]
			if sdb.Id == sp.Id && sp.Position == sdb.Position {
				scenariosDB = append(scenariosDB[0:indexDB], scenariosDB[indexDB+1:]...)
				scenariosPage = append(scenariosPage[0:indexPage], scenariosPage[indexPage+1:]...)
				indexDB = -1
				break
			}
		}
	}
	return scenariosDB, scenariosPage
}

func deleteScenarios(testCase TestCase, scenariosToBeDeleted []s.Scenario) {
	db := database.Db
	for _, scenario := range scenariosToBeDeleted {
		delete := "DELETE FROM scenarioTestcase WHERE testCaseId = ? AND scenarioId = ? AND position = ?"
		stmt, err := db.Prepare(delete)
		if err != nil {
			log.Fatalln(err.Error())
		}
		_, err = stmt.Exec(testCase.Id, scenario.Id, scenario.Position)
		if err != nil {
			log.Fatalln(err.Error())
		}
	}
}

func addScenarios(testCase TestCase, scenariosToBeAdded []s.Scenario) {
	db := database.Db
	for _, scenario := range scenariosToBeAdded {
		insert := "INSERT INTO scenarioTestcase (testCaseId, scenarioId, position) values ( ?, ?, ?)"
		stmt, err := db.Prepare(insert)
		if err != nil {
			log.Fatalln(err.Error())
		}
		_, err = stmt.Exec(testCase.Id, scenario.Id, scenario.Position)
		if err != nil {
			log.Fatalln(err.Error())
		}
	}

}
