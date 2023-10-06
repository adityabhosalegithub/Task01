package validations

//built-in interface for representing an error condition
type error interface {
	Error() string //single method Error() returns a string representing the error message.
}

type CustomError struct { //CustomError type is custom error type that implements the error interface
	message string //field message hold specific err
}

func (e CustomError) Error() string { //Err0r() method of customerror //defined to satisfy custom error
	return e.message
}
