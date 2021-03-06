package environment

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"zarbat_data/features/context"
	"zarbat_data/features/number"
	"zarbat_data/helper"

	"github.com/gorilla/mux"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET ALL")
	environments := GetEnvironments()
	response, _ := json.Marshal(environments)
	fmt.Println(string(response))
	helper.EnsureCors(w)
	fmt.Fprintf(w, string(response))
}

func GetContextsByEnvironmentId(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET CONTEXTS BY ENVIRONMENT")
	params := mux.Vars(r)
	environmentId, _ := strconv.Atoi(params["id"])
	fmt.Println("environmentId", environmentId)
	contexts := context.GetContextsByEnvironmentId(environmentId)
	response, _ := json.Marshal(contexts)
	fmt.Println(string(response))
	helper.EnsureCors(w)
	fmt.Fprintf(w, string(response))
}

func GetNumbersByEnvironmentId(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET NUMBERS BY ENVIRONMENT")
	params := mux.Vars(r)
	environmentId, _ := strconv.Atoi(params["id"])
	fmt.Println("environmentId", environmentId)
	numbers := number.GetNumbersByEnvironmentId(environmentId)
	response, _ := json.Marshal(numbers)
	fmt.Println(string(response))
	helper.EnsureCors(w)
	fmt.Fprintf(w, string(response))
}

func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	environment := GetEnvironment(id)
	response, _ := json.Marshal(environment)
	fmt.Println(string(response))
	helper.EnsureCors(w)
	fmt.Fprintf(w, string(response))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DELETE")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	environment := DeleteEnvironment(id)
	response, _ := json.Marshal(environment)
	helper.EnsureCors(w)
	fmt.Fprintf(w, string(response))
}

func Put(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PUT")
	var environment Environment
	_ = json.NewDecoder(r.Body).Decode(&environment)
	updatedEnvironment := UpdateEnvironment(environment)
	response, _ := json.Marshal(updatedEnvironment)
	helper.EnsureCors(w)
	fmt.Fprintf(w, string(response))
}

func Post(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST")
	var environment Environment
	_ = json.NewDecoder(r.Body).Decode(&environment)
	newEnvironment := AddEnvironment(environment)
	response, _ := json.Marshal(newEnvironment)
	fmt.Println(string(response))
	helper.EnsureCors(w)
	fmt.Fprintf(w, string(response))
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}
