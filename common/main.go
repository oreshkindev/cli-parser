package common

import (
	"fmt"
	"runtime"
	"strings"
)

// Error returns the given error with information about the caller.
//
// It returns a new error that is an alias of the given error and includes
// information about the function where the error occurred, which is the
// function calling the Error function. The information includes the name of
// the function, the file name, and the line number.
//
// If the caller information is not available, the function returns the given
// error without any additional information.
func Error(err error) error {
	// Get information about the caller.
	_, filePath, lineNumber, ok := runtime.Caller(1)
	if !ok {
		// If the information is not available, return the original error.
		return err
	}

	// Trim the file name to remove the path to the "internal/" folder.
	const prefix = "internal/"
	if idx := strings.Index(filePath, prefix); idx != -1 {
		filePath = filePath[idx:]
	}

	// Return a new error with the information about the caller and the
	// original error.
	return fmt.Errorf("%s:%d: %s\n", filePath, lineNumber, err)
}
