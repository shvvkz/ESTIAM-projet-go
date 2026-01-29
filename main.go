package main

import (
	"fmt"
	"log"
	"net/http"

	"employee-api/handlers"
	"employee-api/middleware"
	"employee-api/services"
)

func main() {
	service := services.NewEmployeeService()
	handler := handlers.NewEmployeeHandler(service)

	// create a mux to inject logger middleware
	mux := http.NewServeMux()
	mux.HandleFunc("/employees", handler.GetEmployees)
	mux.HandleFunc("/employee", handler.GetEmployee)

	fmt.Println("Server running on http://localhost:8080")

	log.Fatal(
		http.ListenAndServe(
			":8080",
			middleware.Logger(mux),
		),
	)
}
