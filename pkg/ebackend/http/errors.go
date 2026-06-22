package http

type CustomValidationError struct {
	Message  string
	Messages map[string]string
}

func (e *CustomValidationError) Error() string {
	return e.Message
}

func NewCustomValidationError(message string, messages map[string]string) *CustomValidationError {
	return &CustomValidationError{message, messages}
}

func NewCustomValidationFieldError(message string, field string) *CustomValidationError {
	messages := map[string]string{}
	messages[field] = message
	return &CustomValidationError{message, messages}
}
