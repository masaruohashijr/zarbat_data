package scenario

type Scenario struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ListOfSteps string `json:"listOfSteps"`
	FeatureId   string `json:"featureId"`
	Position    string `json:"position"`
}
