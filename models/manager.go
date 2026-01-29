package models

// Manager
//
// Represents a manager, which is an employee with additional attributes.
//
// Fields:
//   - Employee: Embeds the Employee struct to inherit its fields.
//   - TeamSize: The number of team members managed by the manager.
type Manager struct {
	Employee
	TeamSize int
}
