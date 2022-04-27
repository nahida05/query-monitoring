package service

const internalServerError = "something bad happen"

type customError struct {
	Message string `json:"message"`
}

func newCustomError(content string) customError {
	return customError{
		Message: content,
	}
}
