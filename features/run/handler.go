package run

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"zarbat_data/helper"

	"github.com/gorilla/mux"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET ALL RUN")
	runs := GetRuns()
	response, _ := json.Marshal(runs)
	fmt.Println(string(response))
	helper.EnsureCors(w)
	fmt.Fprintf(w, string(response))
}

func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET RUN")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	run := GetRun(id)
	fmt.Println("id:", run.Id)
	fmt.Println("scenario:", run.ScenarioId)
	fmt.Println("feature:", run.FeatureId)
	fmt.Println("context:", run.ContextId)
	fmt.Println("tags:", run.Tags)
	fmt.Println("name:", run.Name)
	response, _ := json.Marshal(run)
	//fmt.Println(string(response))
	helper.EnsureCors(w)
	fmt.Fprintf(w, string(response))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DELETE RUN")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	run := DeleteRun(id)
	response, _ := json.Marshal(run)
	helper.EnsureCors(w)
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
	fmt.Println("tags", run.Tags)
	fmt.Println("userId", run.UserId)
	updatedRun := UpdateRun(run)
	response, _ := json.Marshal(updatedRun)
	helper.EnsureCors(w)
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
	fmt.Println("tags", run.Tags)
	fmt.Println("userId", run.UserId)
	newRun := AddRun(run)
	response, _ := json.Marshal(newRun)
	//fmt.Println(string(response))
	helper.EnsureCors(w)
	fmt.Fprintf(w, string(response))
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}
