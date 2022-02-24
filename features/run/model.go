package run

type Run struct {
	Id            int       `json:"id"`
	Name          string    `json:"name"`
	ScenarioId    string    `json:"scenarioId"`
	ListOfSteps   string    `json:"listOfSteps"`
	Description   string    `json:"description"`
	Result        string    `json:"result"`
	Logs          string    `json:"logs"`
	FeatureId     string    `json:"featureId"`
	FeatureName   string    `json:"featureName"`
	EnvironmentId string    `json:"environmentId"`
	ContextId     string    `json:"contextId"`
	UserId        string    `json:"userId"`
	RunAt         string    `json:"runAt"`
	Tags          string    `json:"tags"`
	Args          Arguments `json:"args"`
}

type Arguments struct {
	Config   string `json:"config"`
	Url      string `json:"url"`
	Port     string `json:"port"`
	Log      string `json:"log"`
	LogLevel string `json:"logLevel"`
	NTries   string `json:"numberOfTries"`
	Test     string `json:"test"`
}
