package context

import (
	"zarbat_data/features/number"
	"zarbat_data/features/parameter"
)

type Context struct {
	Id              int                   `json:"id"`
	Name            string                `json:"name"`
	Description     string                `json:"description"`
	EnvironmentId   string                `json:"environmentId"`
	EnvironmentName string                `json:"environmentName"`
	PhoneNumbers    []number.PhoneNumber  `json:"phoneNumbers"`
	Parameters      []parameter.Parameter `json:"parameters"`
}
