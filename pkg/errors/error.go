package errors

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	VALIDATION_ERROR      = "validation error"
	DATABASE_ERROR        = "database error"
	NOT_FOUND             = "not found"
	INTERNAL_SERVER_ERROR = "internal server error"
	CONFLICT_ERROR        = "conflict error"
)

type Error struct {
	Code       int    `json:"code"`
	Type       string `json:"type"`
	Message    string `json:"message"`
	Err        error  `json:"-"`
	StatusCode int    `json:"-"`
}

type ResponseError struct {
	Code        int    `json:"code"`
	Type        string `json:"type"`
	Description string `json:"description"`
}
type ResponseErrorArr struct {
	RespErr []ResponseError `json:"errors"`
}

func (e *Error) Error() string {
	var buf bytes.Buffer
	if e.Err != nil {
		buf.WriteString(e.Err.Error())
	} else {
		buf.WriteString(e.Message)
	}
	return buf.String()
}

func (e *ResponseError) Error() string {
	var fields []string
	fields = append(fields, fmt.Sprintf("type: %s, code: %d, description: %s", e.Type, e.Code, e.Description))
	return fmt.Sprintf("{%s}", strings.Join(fields, ","))
}

func (e *ResponseErrorArr) Error() string {
	buff := bytes.NewBufferString("")
	for _, respErr := range e.RespErr {
		buff.WriteString(respErr.Error())
		buff.WriteString("\n")
	}
	return strings.TrimSpace(buff.String())
}
