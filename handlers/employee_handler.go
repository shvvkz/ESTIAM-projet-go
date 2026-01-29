package handlers

import (
	"net/http"
	"employee-api/services"
	"employee-api/models"
	"strconv"
)

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
		WriteJSON(w, http.StatusMethodNotAllowed, models.ResponseError{
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

	WriteJSON(w, http.StatusOK, reponse)
}

func (h *EmployeeHandler) GetEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		WriteJSON(w, http.StatusMethodNotAllowed, models.ResponseError{
			Code:    405,
			Message: "Method not allowed",
		})
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		WriteJSON(w, http.StatusBadRequest, models.ResponseError{
			Code:    400,
			Message: "Missing id parameter",
		})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, models.ResponseError{
			Code:    400,
			Message: "Invalid id parameter",
		})
		return
	}

	employee, found := h.service.GetEmployee(id)
	if !found {
		WriteJSON(w, http.StatusNotFound, models.ResponseError{
			Code:    404,
			Message: "Employee not found",
		})
		return
	}

	WriteJSON(w, http.StatusOK, models.ResponseEmployee{
		Code:    200,
		Message: "Success",
		Data:    employee,
	})
}

