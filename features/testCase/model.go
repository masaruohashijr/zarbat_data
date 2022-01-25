package testCase

import (
	"zarbat_mock/features/scenario"
)

type TestCase struct {
	Id            int                 `json:"id"`
	Name          string              `json:"name"`
	Description   string              `json:"description"`
	EnvironmentId string              `json:"environmentId"`
	ContextId     string              `json:"contextId"`
	Scenarios     []scenario.Scenario `json:"scenarios"`
}

type ScenarioTestCase struct {
	Id          int    `json:"id"`
	TestCaseId  int    `json:"testCaseId"`
	ScenarioId  int    `json:"scenarioId"`
	Position    int    `json:"position"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ListOfSteps string `json:"listOfSteps"`
	FeatureId   int    `json:"featureId"`
}
