package handlers

import (
	"net/http"
	"employee-api/services"
	"employee-api/models"
	"encoding/json"
)
// Employees(w http.ResponseWriter, r *http.Request)
//
// Handles HTTP requests related to employees.
//
// Parameters:
//   - w: The HTTP response writer.
//   - r: The HTTP request.
func (h *EmployeeHandler) Employees(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			h.GetEmployees(w, r)
		case http.MethodPost:
			h.AddEmployee(w, r)
		case http.MethodPut:
			if r.URL.Path == "/employees/raise" {
				h.RaiseEmployeeSalary(w, r)
				return
			}
			writeJSON(w, http.StatusNotFound, models.ResponseError{
				Code:    404,
				Message: "Route not found",
			})
		default:
			writeJSON(w, http.StatusMethodNotAllowed, models.ResponseError{
				Code:    405,
				Message: "Method not allowed",
			})
	}
}

// EmployeeHandler
//
// Handles HTTP requests related to employees.
//
// Fields:
//   - service: An instance of EmployeeService to interact with employee data.
type EmployeeHandler struct {
	service services.EmployeeService
}

// NewEmployeeHandler(service EmployeeService) -> *EmployeeHandler
//
// Creates a new EmployeeHandler with the provided EmployeeService.
//
// Parameters:
//   - service: An instance of EmployeeService.
//
// Returns:
//   - A pointer to the newly created EmployeeHandler.
func NewEmployeeHandler(service services.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service: service}
}

// GetEmployees(w http.ResponseWriter, r *http.Request)
//
// Handles HTTP GET requests to retrieve all employees.
//
// Parameters:
//   - w: The HTTP response writer.
//   - r: The HTTP request.
func (h *EmployeeHandler) GetEmployees(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, models.ResponseError{
			Code:    405,
			Message: "Method not allowed",
		})
		return
	}
	employees := h.service.GetAll()

	reponse := models.ResponseEmployee{
		Code:    200,
		Message: "Success",
		Data:    employees,
	}

	writeJSON(w, http.StatusOK, reponse)
}

// AddEmployee(w http.ResponseWriter, r *http.Request)
//
// Handles HTTP POST requests to add a new employee.
//
// Parameters:
//   - w: The HTTP response writer.
//   - r: The HTTP request.
func (h *EmployeeHandler) AddEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, models.ResponseError{
			Code:    405,
			Message: "Method not allowed",
		})
		return
	}

	var employee models.Employee
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		writeJSON(w, http.StatusBadRequest, models.ResponseError{
			Code:    400,
			Message: "Invalid JSON payload",
		})
		return
	}

	if employee.Name == "" || employee.Salary <= 0 {
		writeJSON(w, http.StatusBadRequest, models.ResponseError{
			Code:    400,
			Message: "Invalid employee data",
		})
		return
	}

	employee.ID = 0

	addedEmployee := h.service.AddEmployee(employee)

	writeJSON(w, http.StatusCreated, models.ResponseEmployee{
		Code:    201,
		Message: "Employee added successfully",
		Data:    addedEmployee,
	})
}

// RaiseEmployeeSalary(w http.ResponseWriter, r *http.Request)
//
// Handles HTTP PUT requests to raise an employee's salary.
//
// Parameters:
//   - w: The HTTP response writer.
//   - r: The HTTP request.
func (h *EmployeeHandler) RaiseEmployeeSalary(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		writeJSON(w, http.StatusMethodNotAllowed, models.ResponseError{
			Code:    405,
			Message: "Method not allowed",
		})
		return
	}

	var req models.RaiseSalaryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, models.ResponseError{
			Code:    400,
			Message: "Invalid JSON payload",
		})
		return
	}

	if req.ID <= 0 || req.Percent <= 0 {
		writeJSON(w, http.StatusBadRequest, models.ResponseError{
			Code:    400,
			Message: "Invalid id or percent",
		})
		return
	}

	employee, ok := h.service.RaiseSalary(req.ID, req.Percent)
	if !ok {
		writeJSON(w, http.StatusNotFound, models.ResponseError{
			Code:    404,
			Message: "Employee not found",
		})
		return
	}

	writeJSON(w, http.StatusOK, models.ResponseEmployee{
		Code:    200,
		Message: "Salary updated successfully",
		Data:    employee,
	})
}
