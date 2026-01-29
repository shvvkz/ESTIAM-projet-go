package main

import (
	"fmt"
	"log"
	"net/http"

	"employee-api/handlers"
	"employee-api/middleware"
	"employee-api/services"
)

func router(handler *handlers.EmployeeHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/employees", handler.Employees)
	mux.HandleFunc("/employees/", handler.Employees)
	return mux
}

func main() {
	service := services.NewEmployeeService()
	handler := handlers.NewEmployeeHandler(service)
	mux := router(handler)

	fmt.Println("Server running on http://localhost:8080\n")
	log.Fatal(
		http.ListenAndServe(
			":8080",
			middleware.Logger(mux),
		),
	)
}
