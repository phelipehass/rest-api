package main

import (
	"../rest-api/task"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	task.CreateData()

	router := mux.NewRouter()
	router.HandleFunc("/tasks", task.GetTasks).Methods("GET")
	router.HandleFunc("/task/{id}", task.GetTaskById).Methods("GET")
	router.HandleFunc("/task/{id}", task.CreateTask).Methods("POST")
	router.HandleFunc("/task/{id}", task.DeleteTask).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
