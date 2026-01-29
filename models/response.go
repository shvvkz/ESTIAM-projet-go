package models

// ResponseEmployee
//
// Represents a standard response structure for employee-related API responses.
//
// Fields:
//   - Code: HTTP status code of the response.
//   - Message: A message describing the response.
//   - Data: The payload of the response, which can be any type.
type ResponseEmployee struct {
	Code	int    `json:"code"`
	Message	string `json:"message"`
	Data	interface{} `json:"data"`
}

// ResponseError
//
// Represents a standard error response structure for API responses.
//
// Fields:
//   - Code: HTTP status code of the error response.
//   - Message: A message describing the error.
type ResponseError struct {
	Code	int    `json:"code"`
	Message	string `json:"message"`
}