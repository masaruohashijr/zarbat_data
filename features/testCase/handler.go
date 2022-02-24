package testCase

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
	testCases := GetTestCases()
	response, _ := json.Marshal(testCases)
	fmt.Println(string(response))
	helper.EnsureCors(w)
	fmt.Fprintf(w, string(response))
}

func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET Test Case")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	testCase := GetTestCase(id)
	response, _ := json.Marshal(testCase)
	fmt.Println(string(response))
	helper.EnsureCors(w)
	fmt.Fprintf(w, string(response))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DELETE")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	testCase := DeleteTestCase(id)
	response, _ := json.Marshal(testCase)
	helper.EnsureCors(w)
	fmt.Fprintf(w, string(response))
}

func Put(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PUT")
	var testCase TestCase
	_ = json.NewDecoder(r.Body).Decode(&testCase)
	updatedTestCase := UpdateTestCase(testCase)
	response, _ := json.Marshal(updatedTestCase)
	helper.EnsureCors(w)
	fmt.Fprintf(w, string(response))
}

func Post(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST")
	var testCase TestCase
	_ = json.NewDecoder(r.Body).Decode(&testCase)
	newTestCase := AddTestCase(testCase)
	response, _ := json.Marshal(newTestCase)
	fmt.Println(string(response))
	helper.EnsureCors(w)
	fmt.Fprintf(w, string(response))
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}
