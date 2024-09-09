package service

import "fmt"

type ErrorUserExists struct {
	Username string
	Email    string
}

func (e *ErrorUserExists) Error() string {
	return fmt.Sprintf("User with username '%s' or email '%s' already exists", e.Username, e.Email)
}
