package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Todo struct {
	Name      string `json:"nome_aqui"`
	Completed bool
	Due       time.Time
}

type Todos []Todo

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/todos/{todoId}", todoShow).Methods("GET", "OPTIONS")

	router.HandleFunc("/todos", todoIndex).Methods("GET")
	router.HandleFunc("/todos", postHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func todoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	var t Todo
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &t)
	j, _ := json.Marshal(t)
	w.Write(j)
}

func todoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	json.NewEncoder(w).Encode(todos)
}
