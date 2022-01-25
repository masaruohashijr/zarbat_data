package run

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ensureCors(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.Header().Add("Access-Control-Allow-Methods", "DELETE, PUT, POST, GET, OPTIONS")
	w.Header().Add("Content-Type", "application/json,text/plain")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Request-Headers", "*")
	return w
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET ALL RUN")
	runs := GetRuns()
	response, _ := json.Marshal(runs)
	fmt.Println(string(response))
	ensureCors(w)
	fmt.Fprintf(w, string(response))
}

func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET RUN")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	run := GetRun(id)
	response, _ := json.Marshal(run)
	fmt.Println(string(response))
	ensureCors(w)
	fmt.Fprintf(w, string(response))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DELETE RUN")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	run := DeleteRun(id)
	response, _ := json.Marshal(run)
	ensureCors(w)
	fmt.Fprintf(w, string(response))
}

func Put(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PUT RUN")
	var run Run
	_ = json.NewDecoder(r.Body).Decode(&run)
	fmt.Println("id", run.Id)
	fmt.Println("name", run.Name)
	fmt.Println("scenarioId", run.ScenarioId)
	fmt.Println("featureId", run.FeatureId)
	fmt.Println("environmentId", run.EnvironmentId)
	fmt.Println("contextId", run.ContextId)
	fmt.Println("userId", run.UserId)
	updatedRun := UpdateRun(run)
	response, _ := json.Marshal(updatedRun)
	ensureCors(w)
	fmt.Fprintf(w, string(response))
}

func Post(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST RUN")
	var run Run
	_ = json.NewDecoder(r.Body).Decode(&run)
	fmt.Println("id", run.Id)
	fmt.Println("name", run.Name)
	fmt.Println("scenarioId", run.ScenarioId)
	fmt.Println("featureId", run.FeatureId)
	fmt.Println("environmentId", run.EnvironmentId)
	fmt.Println("contextId", run.ContextId)
	fmt.Println("userId", run.UserId)
	newRun := AddRun(run)
	response, _ := json.Marshal(newRun)
	fmt.Println(string(response))
	ensureCors(w)
	fmt.Fprintf(w, string(response))
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}
