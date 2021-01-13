package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Task struct {
	ID          int64  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Isconcluded bool   `json:"isconcluded,omitempty"`
}

var listTasks []Task

func main() {
	listTasks = append(listTasks, Task{1, "Correr", "Corrida de 5km", false})
	listTasks = append(listTasks, Task{2, "Fazer café", "Preparar o café da manhã", false})
	listTasks = append(listTasks, Task{3, "Ler", "Fazer a leitura de um livro", false})

	router := mux.NewRouter()
	router.HandleFunc("/tasks", GetTasks).Methods("GET")
	router.HandleFunc("task/{id}", GetTaskById).Methods("GET")
	router.HandleFunc("task/{id}", CreateTask).Methods("POST")
	router.HandleFunc("task/{id}", DeleteTask).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(listTasks)
}

func GetTaskById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idRequest, err := strconv.ParseInt(params["ID"], 10, 64)

	if err != nil {
		log.Fatalf("Error in ID parameter: %s", err.Error())
		return
	}
	for _, item := range listTasks {
		if item.ID == idRequest {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Task{})
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idRequest, err := strconv.ParseInt(params["ID"], 10, 64)

	if err != nil {
		log.Fatalf("Error in ID parameter: %s", err.Error())
		return
	}

	var task Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	task.ID = idRequest
	listTasks = append(listTasks, task)
	json.NewEncoder(w).Encode(listTasks)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	idRequest, err := strconv.ParseInt(params["ID"], 10, 64)

	if err != nil {
		log.Fatalf("Error in ID parameter: %s", err.Error())
		return
	}
	for index, item := range listTasks {
		if item.ID == idRequest {
			listTasks = append(listTasks[:index], listTasks[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(listTasks)
	}
}
