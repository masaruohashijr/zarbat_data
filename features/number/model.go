package number

type PhoneNumber struct {
	Id            int    `json:"id"`
	PhoneNumber   string `json:"phoneNumber"`
	Sid           string `json:"sid"`
	Alias         string `json:"alias"`
	Description   string `json:"description"`
	EnvironmentId string `json:"environmentId"`
	ContextId     string `json:"contextId"`
	Position      string `json:"position"`
}
