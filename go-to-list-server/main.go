package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ApiResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

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
		response = ApiResponse{Code: 0, Message: "Tasks route"}
	default:
		response = ApiResponse{Code: -1, Message: "The requested method is not configured for this resource"}
	}

	json.NewEncoder(resp).Encode(response)
}
