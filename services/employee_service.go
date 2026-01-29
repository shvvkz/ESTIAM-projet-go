package services

import "employee-api/models"

// EmployeeService
//
// Defines the interface for employee-related operations.
//
// Methods:
//   - GetAll() []models.Employee: Retrieves all employees.
type EmployeeService interface {
	GetAll() []models.Employee
	GetEmployee(id int) (models.Employee, bool)
}

// employeeService
//
// Implements the EmployeeService interface.
//
// Fields:
//   - employees: A slice of Employee structs representing the employee data.
type employeeService struct {
	employees []models.Employee
}

// NewEmployeeService -> EmployeeService
//
// Creates a new instance of employeeService with some initial employee data.
//
// Returns:
//   - An instance of EmployeeService.
func NewEmployeeService() EmployeeService {
	return &employeeService{
		employees: []models.Employee{
			{ID: 1, Name: "Alice", Salary: 5000},
			{ID: 2, Name: "Bob", Salary: 7000},
		},
	}
}

// GetAll() -> []models.Employee
//
// Retrieves all employees.
//
// Returns:
//   - A slice of Employee structs.
func (s *employeeService) GetAll() []models.Employee {
	return s.employees
}

func (s *employeeService) GetEmployee(id int) (models.Employee, bool) {
	if id <= 0 || id > len(s.employees) {
		return models.Employee{}, false
	}
	return s.employees[id-1], true
}
