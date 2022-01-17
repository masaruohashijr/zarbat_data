package scenario

func GetScenarios() []Scenario {
	return DbGetScenarios()
}
func GetScenario(id int) Scenario {
	return DbGetScenario(id)
}
func AddScenario(scenario Scenario) Scenario {
	return DbAddScenario(scenario)
}
func UpdateScenario(scenario Scenario) Scenario {
	return DbUpdateScenario(scenario)
}
func DeleteScenario(id int) Scenario {
	return DbDeleteScenario(id)
}
