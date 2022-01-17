package run

type Run struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ScenarioId  string `json:"scenarioId"`
	Story       string `json:"story"`
	FeatureId   string `json:"featureId"`
}
