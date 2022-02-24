package parameter

type Parameter struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ContextId   string `json:"contextId"`
	Value       string `json:"value"`
	Position    string `json:"position"`
}
