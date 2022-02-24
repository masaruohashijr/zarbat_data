package scenario

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"zarbat_data/helper"

	"github.com/gorilla/mux"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET ALL")
	scenarios := GetScenarios()
	response, _ := json.Marshal(scenarios)
	fmt.Println(string(response))
	helper.EnsureCors(w)
	fmt.Fprintf(w, string(response))
}

func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	println("id = ", id)
	scenario := GetScenario(id)
	response, _ := json.Marshal(scenario)
	fmt.Println(string(response))
	helper.EnsureCors(w)
	fmt.Fprintf(w, string(response))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DELETE")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	scenario := DeleteScenario(id)
	response, _ := json.Marshal(scenario)
	helper.EnsureCors(w)
	fmt.Fprintf(w, string(response))
}

func Put(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PUT")
	var scenario Scenario
	_ = json.NewDecoder(r.Body).Decode(&scenario)
	updatedScenario := UpdateScenario(scenario)
	response, _ := json.Marshal(updatedScenario)
	fmt.Println(string(response))
	helper.EnsureCors(w)
	fmt.Fprintf(w, string(response))
}

func Post(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST")
	var scenario Scenario
	_ = json.NewDecoder(r.Body).Decode(&scenario)
	newScenario := AddScenario(scenario)
	response, _ := json.Marshal(newScenario)
	fmt.Println(string(response))
	helper.EnsureCors(w)
	fmt.Fprintf(w, string(response))
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}

func Teste(w http.ResponseWriter, r *http.Request) {
	println("BATEU")
}
