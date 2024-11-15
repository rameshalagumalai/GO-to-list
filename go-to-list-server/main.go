package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server started at port 2323")
	log.Fatal(http.ListenAndServe(":2323", nil))
}
