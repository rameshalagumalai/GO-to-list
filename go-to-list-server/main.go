package main

import (
	"fmt"
	"log"
	"net/http"
)

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
	resp.Write([]byte("Tasks route"))
}
