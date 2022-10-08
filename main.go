package main

import (
	"fmt"
	"log"
	"net/http"
	"todoApp/routes"
)

func main() {
	routes.Init()
	fmt.Println("Server is listening to port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
