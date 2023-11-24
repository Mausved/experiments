package another_pkg

type CustomError struct {
	message string
}

func (e *CustomError) Error() string {
	return e.message
}

func (e *CustomError) Is(target error) bool {
	_, ok := target.(*CustomError)
	return ok
}
