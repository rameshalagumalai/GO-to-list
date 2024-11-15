package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ApiResponse struct {
	Code int `json:"code"`
	Message string `json:"message,omitempty"`
	Data any `json:"data"`
}

type Task struct {
	Id int `json:"id"`
	Name string `json:"task_name"`
}

type TaskPayload struct {
	Name string `json:"task_name"`
}

var tasks = []Task{}

func main() {
	http.HandleFunc("/", handleHomeRoute)
	http.HandleFunc("/tasks", handleTasksRoute)

	fmt.Println("Server started at port 2323")
	log.Fatal(http.ListenAndServe(":2323", nil))
}

// "/" route
func handleHomeRoute(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Home route"))
}

// "/tasks" route
func handleTasksRoute(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("Content-Type", "application/json")

	var response ApiResponse
	var method = req.Method
	switch method {
	case http.MethodGet:
		response = ApiResponse{Code: 0, Data: tasks}
	case http.MethodPost:
		var createdTask = createTask(req)
		response = ApiResponse{Code: 0, Message: "Task added successfully", Data: createdTask}
	default:
		response = ApiResponse{Code: -1, Data: "The requested method is not configured for this resource"}
	}

	json.NewEncoder(resp).Encode(response)
}

func createTask(req *http.Request) Task {
	var taskPayload TaskPayload
	json.NewDecoder(req.Body).Decode(&taskPayload)

	var createdTask = Task{Id: len(tasks)+1, Name: taskPayload.Name}
	tasks = append(tasks, createdTask)
	return createdTask
}
