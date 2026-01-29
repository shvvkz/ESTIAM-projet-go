package models

// Employee
//
// Represents an employee with ID, Name, and Salary.
//
// Fields:
//   - ID: Unique identifier for the employee.
//   - Name: Full name of the employee.
//   - Salary: Current salary of the employee.
//
// Methods:
//   - Raise(percent float64): Increases the employee's salary by the given percentage.
type Employee struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
}

// Raise(percent float64)
//
// Increases the employee's salary by the specified percentage.
//
// Parameters:
//   - percent: The percentage by which to increase the salary.
func (e *Employee) Raise(percent float64) {
	e.Salary += e.Salary * percent / 100
}
