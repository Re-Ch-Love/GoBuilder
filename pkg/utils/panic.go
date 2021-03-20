package utils

/*
ReturnErrorFromPanic is a function to handler a panic.
Usage:

func foo() (retRrr error) {
	defer utils.ReturnErrorFromPanic(&retErr, func(err error) {
		fmt.Println(err)
	})
	// If you unwanted the handler function, you can use `nil` to instead of it.
	// defer utils.ReturnErrorFromPanic(&retErr, nil)
	// If there is a panic, `foo` will return an error which recover from ReturnErrorFromPanic.
	// If there is an error to return, you can panic it,
	// then the handler function can do something, like log the error, release the resource, etc.
	if err != nil {
		panic(err)
	}
	return
}
*/
func ReturnErrorFromPanic(ep *error, errorHandlerFunc func(err error)) {
	if i := recover(); i != nil {
		if e, ok := i.(error); ok {
			*ep = e
			if errorHandlerFunc != nil {
				errorHandlerFunc(e)
			}
		} else {
			panic(cannotRecoverErrorError{})
		}
	}
}

// cannotRecoverErrorError means cannot recover error's error
type cannotRecoverErrorError struct {
}

func (e cannotRecoverErrorError) Error() string {
	return "Cannot get an error from `recover`, `ReturnErrorFromPanic` must get an error from recover, so you must panic an error or nil."
}
