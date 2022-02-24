package testCase

import (
	"zarbat_data/features/scenario"
)

type TestCase struct {
	Id            int                 `json:"id"`
	Name          string              `json:"name"`
	Description   string              `json:"description"`
	EnvironmentId string              `json:"environmentId"`
	ContextId     string              `json:"contextId"`
	Scenarios     []scenario.Scenario `json:"scenarios"`
}
