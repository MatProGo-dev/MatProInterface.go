package optim

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

// INFINITY represents a large constant value used for unbounded constraints.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
const INFINITY = 1e100

// SumVars returns the sum of the given variables by creating a new empty
// expression and adding the given variables to it.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func SumVars(vs ...Variable) ScalarExpression {
	newExpr := NewScalarExpression(0)
	for _, v := range vs {
		sum, _ := newExpr.Plus(v)
		newExpr, _ = ToScalarExpression(sum)
	}
	return newExpr
}

// SumRow returns the sum of all the variables in a single specified row of
// a variable matrix.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func SumRow(vs [][]Variable, row int) ScalarExpression {
	newExpr := NewScalarExpression(0)
	for col := 0; col < len(vs[0]); col++ {
		sum, _ := newExpr.Plus(vs[row][col])
		newExpr, _ = ToScalarExpression(sum)
	}
	return newExpr
}

// SumCol returns the sum of all variables in a single specified column of
// a variable matrix.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func SumCol(vs [][]Variable, col int) ScalarExpression {
	newExpr := NewScalarExpression(0)
	for row := 0; row < len(vs); row++ {
		sum, _ := newExpr.Plus(vs[row][col])
		newExpr, _ = ToScalarExpression(sum)
	}
	return newExpr
}

// FindInSlice identifies if the input xIn is in the slice sliceIn.
// If it is, then this function returns the index such that xIn = sliceIn[index] and no errors.
// If it is not, then this function returns the index -1 and an error.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func FindInSlice(xIn interface{}, sliceIn interface{}) (int, error) {
	// Constants
	allowedTypes := []string{"string", "int", "uint64", "Variable"}

	switch x := xIn.(type) {
	case string:
		slice := sliceIn.([]string)

		// Perform Search
		xLocationInSliceIn := -1

		for sliceIndex, sliceValue := range slice {
			if x == sliceValue {
				xLocationInSliceIn = sliceIndex
			}
		}

		return xLocationInSliceIn, nil

	case int:
		slice := sliceIn.([]int)

		// Perform Search
		xLocationInSliceIn := -1

		for sliceIndex, sliceValue := range slice {
			if x == sliceValue {
				xLocationInSliceIn = sliceIndex
			}
		}

		return xLocationInSliceIn, nil

	case uint64:
		slice := sliceIn.([]uint64)

		// Perform Search
		xLocationInSliceIn := -1

		for sliceIndex, sliceValue := range slice {
			if x == sliceValue {
				xLocationInSliceIn = sliceIndex
			}
		}

		return xLocationInSliceIn, nil

	case Variable:
		slice, ok := sliceIn.([]Variable)
		if !ok {
			return -1, fmt.Errorf(
				"the input slice is of type %T, but the element we're searching for is of type %T",
				sliceIn,
				x,
			)
		}

		// Perform Search
		xLocationInSliceIn := -1
		for sliceIndex, sliceValue := range slice {
			if x.ID == sliceValue.ID {
				xLocationInSliceIn = sliceIndex
			}
		}

		return xLocationInSliceIn, nil

	default:

		return -1, fmt.Errorf(
			"the FindInSlice() function was only defined for types %v, not type %T:",
			allowedTypes,
			xIn,
		)
	}

}

// Unique returns the unique list of uint64 values in the given slice.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func Unique(listIn []uint64) []uint64 {
	// Create unique list
	var uniqueList []uint64

	// For each int in the list, determine if it previously existed in the list.
	for listIndex, tempElt := range listIn {
		// Don't do any checks if this is the first element.
		if listIndex == 0 {
			uniqueList = append(uniqueList, tempElt)
			continue
		}

		// check to see if the current element exists in the uniqueList.
		if foundIndex, _ := FindInSlice(tempElt, uniqueList); foundIndex == -1 {
			// tempElt does not exist in uniqueList already. Add it.
			uniqueList = append(uniqueList, tempElt)
		}
		// Otherwise, don't add it.
	}

	return uniqueList
}

// OnesVector returns a vector of ones with the given length.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func OnesVector(lengthIn int) mat.VecDense {
	// Create the empty slice.
	elts := make([]float64, lengthIn)

	for eltIndex := 0; eltIndex < lengthIn; eltIndex++ {
		elts[eltIndex] = 1.0
	}
	return *mat.NewVecDense(lengthIn, elts)
}

// ZerosVector returns a vector of zeros with the given length.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func ZerosVector(lengthIn int) mat.VecDense {
	// Create the empty slice.
	elts := make([]float64, lengthIn)

	for eltIndex := 0; eltIndex < lengthIn; eltIndex++ {
		elts[eltIndex] = 0.0
	}
	return *mat.NewVecDense(lengthIn, elts)
}

// ZerosMatrix returns a dense matrix of all zeros with the given number of
// rows and columns.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func ZerosMatrix(nR, nC int) mat.Dense {
	// Create empty slice
	elts := make([]float64, nR*nC)
	for rowIndex := 0; rowIndex < nR; rowIndex++ {
		for colIndex := 0; colIndex < nC; colIndex++ {
			elts[rowIndex*nC+colIndex] = 0.0
		}
	}

	return *mat.NewDense(nR, nC, elts)
}

// Identity returns a square identity matrix of the given dimension.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func Identity(dim int) mat.Dense {
	// Create the empty matrix.
	zeroBase := ZerosMatrix(dim, dim)

	// Populate Diagonal
	for rowIndex := 0; rowIndex < dim; rowIndex++ {
		zeroBase.Set(rowIndex, rowIndex, 1.0)
	}

	return zeroBase
}

// CheckExtras checks the extras slice for any errors. It returns an error if
// one of the extras is an error, or if there are more than one extras.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func CheckExtras(extras []interface{}) error {
	// Constants

	// Check all of the extras to see if one of them contains an error
	switch {
	case len(extras) == 1:
		if extras[0] == nil {
			return nil
		}

		// Check to see if the input is an error or not.
		switch e := extras[0].(type) {
		case error:
			return e
		default:
			return fmt.Errorf(
				"unexpected type of input as an 'extra': %T",
				e,
			)
		}

	case len(extras) > 1:
		return fmt.Errorf(
			"did not expect to receive more than one element in 'extras' input; received %v",
			len(extras),
		)
	}

	// If extras has length 0, then return nil
	return nil
}

// CheckErrors checks the extras slice of errors. It returns an error if one
// of the extras is non-nil, or if there are more than one extras.
//
// Deprecated: This package is deprecated. Please use github.com/MatProGo-dev/SymbolicMath.go instead.
func CheckErrors(extras []error) error {
	// Constants

	// Check all of the extras to see if one of them contains an error
	switch {
	case len(extras) == 1:
		return extras[0]

	case len(extras) > 1:
		return fmt.Errorf(
			"did not expect to receive more than one element in 'extras' input; received %v",
			len(extras),
		)
	}

	// If extras has length 0, then return nil
	return nil
}
