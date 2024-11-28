package domain

type ValidationResult struct {
	Valid bool
	Error error
}

type ValidationInput struct {
	Contents []byte
	FileName string
}
