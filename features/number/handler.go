package number

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
	numbers := GetNumbers()
	response, _ := json.Marshal(numbers)
	fmt.Println(string(response))
	helper.EnsureCors(w)
	fmt.Fprintf(w, string(response))
}

func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	number := GetNumber(id)
	response, _ := json.Marshal(number)
	fmt.Println(string(response))
	helper.EnsureCors(w)
	fmt.Fprintf(w, string(response))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DELETE")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	number := DeleteNumber(id)
	response, _ := json.Marshal(number)
	helper.EnsureCors(w)
	fmt.Fprintf(w, string(response))
}

func Put(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PUT NUMBER")
	var number PhoneNumber
	_ = json.NewDecoder(r.Body).Decode(&number)
	updatedNumber := UpdateNumber(number)
	response, _ := json.Marshal(updatedNumber)
	helper.EnsureCors(w)
	fmt.Fprintf(w, string(response))
}

func Post(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST")
	var number PhoneNumber
	_ = json.NewDecoder(r.Body).Decode(&number)
	newNumber := AddNumber(number)
	response, _ := json.Marshal(newNumber)
	fmt.Println(string(response))
	helper.EnsureCors(w)
	fmt.Fprintf(w, string(response))
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}
