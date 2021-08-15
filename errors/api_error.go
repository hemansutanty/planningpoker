package errors

//APIError struct
type APIError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
