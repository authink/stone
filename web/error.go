package web

import "fmt"

type ClientError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *ClientError) Error() string {
	return fmt.Sprintf("%s(%s)", e.Message, e.Code)
}

var _ error = (*ClientError)(nil)
