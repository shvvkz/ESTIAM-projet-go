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
	AddEmployee(employee models.Employee) models.Employee
	RaiseSalary(id int, percent float64) (models.Employee, bool)
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

// AddEmployee(employee models.Employee) -> models.Employee
//
// Adds a new employee to the list.
//
// Parameters:
//   - employee: An Employee struct representing the new employee.
//
// Returns:
//   - The added Employee struct with an assigned ID.
func (s *employeeService) AddEmployee(employee models.Employee) models.Employee {
	employee.ID = len(s.employees) + 1
	s.employees = append(s.employees, employee)
	return employee
}

// RaiseSalary(id int, percent float64) -> (models.Employee, bool)
//
// Raises the salary of an employee by a given percentage.
//
// Parameters:
//   - id: The ID of the employee whose salary is to be raised.
//   - percent: The percentage by which to raise the salary.
//
// Returns:
//   - The updated Employee struct and a boolean indicating success.
func (s *employeeService) RaiseSalary(id int, percent float64) (models.Employee, bool) {
	for i := range s.employees {
		if s.employees[i].ID == id {
			s.employees[i].Raise(percent)
			return s.employees[i], true
		}
	}
	return models.Employee{}, false
}