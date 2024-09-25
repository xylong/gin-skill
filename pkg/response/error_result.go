package response

// ErrorResult 异常返回
type ErrorResult struct {
	err error
}

func (r *ErrorResult) Unwrap() any {
	if r.err != nil {
		panic(r.err.Error())
	}

	return nil
}

func Result(err error) *ErrorResult {
	return &ErrorResult{
		err: err,
	}
}
