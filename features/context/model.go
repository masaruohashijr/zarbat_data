package context

type Context struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	EnvironmentId int    `json:"environmentId"`
	Environment   string `json:"environmentName"`
}
