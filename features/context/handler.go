package context

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ensureCors(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Request-Headers", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Add("Access-Control-Allow-Methods", "DELETE, PUT, POST, GET, OPTIONS")
	w.Header().Add("Content-Type", "application/json,text/plain")
	return w
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET ALL")
	contexts := GetContexts()
	response, _ := json.Marshal(contexts)
	fmt.Println(string(response))
	ensureCors(w)
	fmt.Fprintf(w, string(response))
}

func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	context := GetContext(id)
	response, _ := json.Marshal(context)
	fmt.Println(string(response))
	ensureCors(w)
	fmt.Fprintf(w, string(response))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DELETE")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	context := DeleteContext(id)
	response, _ := json.Marshal(context)
	ensureCors(w)
	fmt.Fprintf(w, string(response))
}

func Put(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PUT")
	var context Context
	_ = json.NewDecoder(r.Body).Decode(&context)
	updatedContext := UpdateContext(context)
	response, _ := json.Marshal(updatedContext)
	fmt.Println(string(response))
	ensureCors(w)
	fmt.Fprintf(w, string(response))
}

func Post(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST")
	var context Context
	_ = json.NewDecoder(r.Body).Decode(&context)
	newContext := AddContext(context)
	response, _ := json.Marshal(newContext)
	fmt.Println(string(response))
	ensureCors(w)
	fmt.Fprintf(w, string(response))
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}
