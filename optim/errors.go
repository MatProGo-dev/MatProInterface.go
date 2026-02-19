package optim

import (
	"fmt"
)

/*
dimension.go
*/

/* Type Definitions */

// DimensionError represents an error that occurs when two expressions have
// incompatible dimensions for a given operation.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
type DimensionError struct {
	Arg1      Expression
	Arg2      Expression
	Operation string // Either multiply or Plus
}

// UnexpectedInputError represents an error that occurs when an unexpected
// input type is provided to an operation.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
type UnexpectedInputError struct {
	InputInQuestion interface{}
	Operation       string
}

/* Methods */

// Error returns a string representation of the DimensionError.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (de DimensionError) Error() string {
	dimStrings := de.ArgDimsAsStrings()
	return fmt.Sprintf(
		"dimension error: Cannot perform %v between expression of dimension %v and expression of dimension %v",
		de.Operation,
		dimStrings[0],
		dimStrings[1],
	)
}

// ArgDimsAsStrings returns the dimensions of both arguments as a slice of strings.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (de DimensionError) ArgDimsAsStrings() []string {
	// Create string for arg 1
	arg1DimsAsString := "("
	for ii, dimValue := range de.Arg1.Dims() {
		arg1DimsAsString += fmt.Sprintf("%v", dimValue)
		if ii != len(de.Arg1.Dims())-1 { // If this isn't the last element of size
			arg1DimsAsString += ","
		}
	}
	arg1DimsAsString += ")"

	// Create string for arg 2
	arg2DimsAsString := "("
	for ii, dimValue := range de.Arg2.Dims() {
		arg2DimsAsString += fmt.Sprintf("%v", dimValue)
		if ii != len(de.Arg2.Dims())-1 { // If this isn't the last element of size
			arg2DimsAsString += ","
		}
	}
	arg2DimsAsString += ")"

	return []string{arg1DimsAsString, arg2DimsAsString}

}

// Error returns a string representation of the UnexpectedInputError.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func (uie UnexpectedInputError) Error() string {
	return fmt.Sprintf(
		"Unexpected input to \"%v\" operation: %T",
		uie.Operation,
		uie.InputInQuestion,
	)
}
