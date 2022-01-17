package main

import (
	"net/http"
	"zarbat_mock/database"
	"zarbat_mock/features/context"
	"zarbat_mock/features/environment"
	"zarbat_mock/features/feature"
	"zarbat_mock/features/number"
	"zarbat_mock/features/parameter"
	"zarbat_mock/features/run"
	"zarbat_mock/features/scenario"
	"zarbat_mock/features/schedule"
	"zarbat_mock/features/step"
	"zarbat_mock/features/testCase"

	"github.com/gorilla/mux"
)

func main() {
	database.InitDatabase()
	router := mux.NewRouter()
	// User
	/*	router.HandleFunc("/user", user.GetAll).Methods("GET", "OPTIONS")
		router.HandleFunc("/user/{id}", user.Delete).Methods("DELETE", "OPTIONS")
		router.HandleFunc("/user/{id}", user.Get).Methods("GET", "OPTIONS")
		router.HandleFunc("/user", user.Post).Methods("POST", "OPTIONS")
		router.HandleFunc("/user", user.Put).Methods("PUT", "OPTIONS")*/
	// Scenario
	router.HandleFunc("/scenario", scenario.Put).Methods("PUT", "OPTIONS")
	router.HandleFunc("/scenario", scenario.GetAll).Methods("GET", "OPTIONS")
	router.HandleFunc("/scenario/{id}", scenario.Delete).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/scenario/{id}", scenario.Get).Methods("GET", "OPTIONS")
	router.HandleFunc("/scenario", scenario.Post).Methods("POST", "OPTIONS")
	// Context
	router.HandleFunc("/context", context.GetAll).Methods("GET", "OPTIONS")
	router.HandleFunc("/context/{id}", context.Delete).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/context/{id}", context.Get).Methods("GET", "OPTIONS")
	router.HandleFunc("/context", context.Post).Methods("POST", "OPTIONS")
	router.HandleFunc("/context", context.Put).Methods("PUT", "OPTIONS")
	// Environment
	router.HandleFunc("/environment", environment.GetAll).Methods("GET", "OPTIONS")
	router.HandleFunc("/environment/{id}", environment.Delete).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/environment/{id}/context", environment.GetByEnv).Methods("GET", "OPTIONS")
	router.HandleFunc("/environment/{id}", environment.Get).Methods("GET", "OPTIONS")
	router.HandleFunc("/environment", environment.Post).Methods("POST", "OPTIONS")
	router.HandleFunc("/environment", environment.Put).Methods("PUT", "OPTIONS")
	// Feature
	router.HandleFunc("/feature", feature.GetAll).Methods("GET", "OPTIONS")
	router.HandleFunc("/feature/{id}", feature.Delete).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/feature/{id}", feature.Get).Methods("GET", "OPTIONS")
	router.HandleFunc("/feature", feature.Post).Methods("POST", "OPTIONS")
	router.HandleFunc("/feature", feature.Put).Methods("PUT", "OPTIONS")
	// Number
	router.HandleFunc("/number", number.GetAll).Methods("GET", "OPTIONS")
	router.HandleFunc("/number/{id}", number.Delete).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/number/{id}", number.Get).Methods("GET", "OPTIONS")
	router.HandleFunc("/number", number.Post).Methods("POST", "OPTIONS")
	router.HandleFunc("/number", number.Put).Methods("PUT", "OPTIONS")
	// Parameter
	router.HandleFunc("/parameter", parameter.GetAll).Methods("GET", "OPTIONS")
	router.HandleFunc("/parameter/{id}", parameter.Delete).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/parameter/{id}", parameter.Get).Methods("GET", "OPTIONS")
	router.HandleFunc("/parameter", parameter.Post).Methods("POST", "OPTIONS")
	router.HandleFunc("/parameter", parameter.Put).Methods("PUT", "OPTIONS")
	// Run
	router.HandleFunc("/run", run.GetAll).Methods("GET", "OPTIONS")
	router.HandleFunc("/run/{id}", run.Delete).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/run/{id}", run.Get).Methods("GET", "OPTIONS")
	router.HandleFunc("/run", run.Post).Methods("POST", "OPTIONS")
	router.HandleFunc("/run", run.Put).Methods("PUT", "OPTIONS")
	// Schedule
	router.HandleFunc("/schedule", schedule.GetAll).Methods("GET", "OPTIONS")
	router.HandleFunc("/schedule/{id}", schedule.Delete).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/schedule/{id}", schedule.Get).Methods("GET", "OPTIONS")
	router.HandleFunc("/schedule", schedule.Post).Methods("POST", "OPTIONS")
	router.HandleFunc("/schedule", schedule.Put).Methods("PUT", "OPTIONS")
	// Step
	router.HandleFunc("/step", step.GetAll).Methods("GET", "OPTIONS")
	router.HandleFunc("/step/{id}", step.Delete).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/step/{id}", step.Get).Methods("GET", "OPTIONS")
	router.HandleFunc("/step", step.Post).Methods("POST", "OPTIONS")
	router.HandleFunc("/step", step.Put).Methods("PUT", "OPTIONS")
	// TestCase
	router.HandleFunc("/testCase", testCase.GetAll).Methods("GET", "OPTIONS")
	router.HandleFunc("/testCase/{id}", testCase.Delete).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/testCase/{id}", testCase.Get).Methods("GET", "OPTIONS")
	router.HandleFunc("/testCase", testCase.Post).Methods("POST", "OPTIONS")
	router.HandleFunc("/testCase", testCase.Put).Methods("PUT", "OPTIONS")
	http.Handle("/", router)
	addr := ":5000"
	println("Zarbat Mock")
	//http.ListenAndServe(addr, cors.Default().Handler(router))
	http.ListenAndServe(addr, router)
}
