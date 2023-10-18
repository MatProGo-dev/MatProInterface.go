package errors

import (
	"fmt"
)

/*
dimension.go
*/

type DimensionError struct {
	Arg1Dims  []uint
	Arg2Dims  []uint
	Operation string // Either multiply or Plus
}

func (de DimensionError) Error() string {
	dimStrings := de.ArgDimsAsStrings()
	return fmt.Sprintf(
		"dimension error: Cannot perform %v between expression of dimension %v and expression of dimension %v",
		de.Operation,
		dimStrings[0],
		dimStrings[1],
	)
}

func (de DimensionError) ArgDimsAsStrings() []string {
	// Create string for arg 1
	arg1DimsAsString := "("
	for ii, dimValue := range de.Arg1Dims {
		arg1DimsAsString += fmt.Sprintf("%v", dimValue)
		if ii != len(de.Arg1Dims)-1 { // If this isn't the last element of size
			arg1DimsAsString += ","
		}
	}
	arg1DimsAsString += ")"

	// Create string for arg 2
	arg2DimsAsString := "("
	for ii, dimValue := range de.Arg2Dims {
		arg2DimsAsString += fmt.Sprintf("%v", dimValue)
		if ii != len(de.Arg2Dims)-1 { // If this isn't the last element of size
			arg2DimsAsString += ","
		}
	}
	arg2DimsAsString += ")"

	return []string{arg1DimsAsString, arg2DimsAsString}

}
