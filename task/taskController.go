package task

import (
	"../model"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

var listTasks []model.Task

func CreateData() {
	listTasks = append(listTasks, model.Task{1, "Correr", "Corrida de 5km", false})
	listTasks = append(listTasks, model.Task{2, "Fazer café", "Preparar o café da manhã", false})
	listTasks = append(listTasks, model.Task{3, "Ler", "Fazer a leitura de um livro", false})
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(listTasks)
}

func GetTaskById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idRequest, err := IdParameterTreatment(params["id"])
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
		return
	}
	for _, item := range listTasks {
		if item.ID == idRequest {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&model.Task{})
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idRequest, err := IdParameterTreatment(params["id"])
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
		return
	}

	var task model.Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	task.ID = idRequest
	listTasks = append(listTasks, task)
	json.NewEncoder(w).Encode(listTasks)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idRequest, err := IdParameterTreatment(params["id"])
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
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

func IdParameterTreatment(id string) (int64, error) {
	regex, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Println("Error in regex")
		return 0, err
	}
	id = regex.ReplaceAllString(id, "")
	idRequest, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Println("Error in convertion string to int")
		return 0, err
	}

	return idRequest, nil
}
