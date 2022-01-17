package scenario

type Scenario struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Story       string `json:"story"`
	FeatureId   string `json:"featureId"`
}
