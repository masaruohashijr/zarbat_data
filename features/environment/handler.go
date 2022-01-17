package environment

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"zarbat_mock/features/context"

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
	fmt.Println("GET ALL")
	environments := GetEnvironments()
	response, _ := json.Marshal(environments)
	fmt.Println(string(response))
	ensureCors(w)
	fmt.Fprintf(w, string(response))
}

func GetByEnv(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET BY ENV")
	params := mux.Vars(r)
	environmentId, _ := strconv.Atoi(params["id"])
	fmt.Println("environmentId", environmentId)
	contexts := context.GetContextsByEnv(environmentId)
	response, _ := json.Marshal(contexts)
	fmt.Println(string(response))
	ensureCors(w)
	fmt.Fprintf(w, string(response))
}

func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	environment := GetEnvironment(id)
	response, _ := json.Marshal(environment)
	fmt.Println(string(response))
	ensureCors(w)
	fmt.Fprintf(w, string(response))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DELETE")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	environment := DeleteEnvironment(id)
	response, _ := json.Marshal(environment)
	ensureCors(w)
	fmt.Fprintf(w, string(response))
}

func Put(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PUT")
	var environment Environment
	_ = json.NewDecoder(r.Body).Decode(&environment)
	updatedEnvironment := UpdateEnvironment(environment)
	response, _ := json.Marshal(updatedEnvironment)
	ensureCors(w)
	fmt.Fprintf(w, string(response))
}

func Post(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST")
	var environment Environment
	_ = json.NewDecoder(r.Body).Decode(&environment)
	newEnvironment := AddEnvironment(environment)
	response, _ := json.Marshal(newEnvironment)
	fmt.Println(string(response))
	ensureCors(w)
	fmt.Fprintf(w, string(response))
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}
