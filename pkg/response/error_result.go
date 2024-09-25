package response

import "fmt"

// ErrorResult 异常返回
type ErrorResult struct {
	data interface{}
	err  error
}

func (r *ErrorResult) Unwrap() any {
	if r.err != nil {
		panic(r.err.Error())
	}

	return r.data
}

func Result(args ...any) *ErrorResult {
	if len(args) == 1 {
		if args[0] == nil {
			return &ErrorResult{}
		}
		if err, ok := args[0].(error); ok {
			return &ErrorResult{nil, err}
		}
	}

	if len(args) == 2 {
		if args[1] == nil {
			return &ErrorResult{args[0], nil}
		}
		if err, ok := args[1].(error); ok {
			return &ErrorResult{args[0], err}
		}
	}

	return &ErrorResult{nil, fmt.Errorf("error result format")}
}
