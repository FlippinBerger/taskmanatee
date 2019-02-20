package main

// TMError is a naive implementation of app specific errors for Task Man
type TMError struct {
	s string
}

func (e *TMError) Error() string {
	return e.s
}
