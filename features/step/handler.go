package step

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
	steps := GetSteps()
	response, _ := json.Marshal(steps)
	fmt.Println(string(response))
	helper.EnsureCors(w)
	fmt.Fprintf(w, string(response))
}

func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	step := GetStep(id)
	response, _ := json.Marshal(step)
	fmt.Println(string(response))
	helper.EnsureCors(w)
	fmt.Fprintf(w, string(response))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DELETE")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	step := DeleteStep(id)
	response, _ := json.Marshal(step)
	helper.EnsureCors(w)
	fmt.Fprintf(w, string(response))
}

func Put(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PUT")
	var step Step
	_ = json.NewDecoder(r.Body).Decode(&step)
	updatedStep := UpdateStep(step)
	response, _ := json.Marshal(updatedStep)
	helper.EnsureCors(w)
	fmt.Fprintf(w, string(response))
}

func Post(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST")
	var step Step
	_ = json.NewDecoder(r.Body).Decode(&step)
	newStep := AddStep(step)
	response, _ := json.Marshal(newStep)
	fmt.Println(string(response))
	helper.EnsureCors(w)
	fmt.Fprintf(w, string(response))
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}
