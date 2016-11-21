package aerror

import (
	"errors"
	"runtime"
)

type Stack struct {
	File string
	Line int
}
type AError struct {
	error
	Trace []Stack
}

func TraceNow(skip int) []Stack {
	t := make([]Stack, 0)
	for {
		_, file, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}
		t = append(t, Stack{File: file, Line: line})
		skip++
	}
	return t
}

func WrapError(oe error) *AError {
	switch a := oe.(type) {
	case *AError:
		return a
	}
	if oe == nil {
		return nil
	}
	return &AError{oe, TraceNow(2)}
}

func New(msg string) *AError {

	return &AError{errors.New(msg), TraceNow(2)}

}
