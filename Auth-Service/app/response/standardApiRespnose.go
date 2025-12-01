package helpers

type ApiResponse[T any] struct {
	Success bool
	Status int
	Message string
	Data T
}

func StandardApiResponse[T any](success bool, status int, message string, data T) *ApiResponse[T]{
	response := &ApiResponse[T]{
		Success: success,
		Status: status,
		Message: message,
		Data: data,
	}

	return response
}