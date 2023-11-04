package matrix_test

import (
	"github.com/MatProGo-dev/MatProInterface.go/symbolic"
	"github.com/MatProGo-dev/MatProInterface.go/symbolic/matrix"
	"strings"
	"testing"
)

/*
expression_test.go
Description:
	Tests the methods used for handling matrix expressions.
*/

/*
TestExpression_ToMatrixExpression1
Description:

	Tests the matrix expression conversion method when a non-matrix
	is provided.
*/
func TestExpression_ToMatrixExpression1(t *testing.T) {
	// Constants
	b1 := false

	// Algorithm
	_, err := symbolic.ToMatrixExpression(b1)
	if err == nil {
		t.Errorf("no error was thrown, but there should have been!")
	} else {
		if !strings.Contains(
			err.Error(),
			matrix.TypeError{b1}.Error(),
		) {
			t.Errorf("unexpected error: %v", err)
		}
	}
}
