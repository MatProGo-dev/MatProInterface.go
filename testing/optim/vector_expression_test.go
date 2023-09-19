package optim_test

import (
	"github.com/MatProGo-dev/MatProInterface.go/optim"
	"testing"
)

/*
TestVectorExpression_IsVectorExpression1
Description:

	Tests whether or not IsVectorExpression() works on a KVector object.
*/
func TestVectorExpression_IsVectorExpression1(t *testing.T) {
	// Constants
	N := 10
	K1 := optim.KVector(optim.OnesVector(N))

	// Check
	if !optim.IsVectorExpression(K1) {
		t.Errorf("K1 was determined to NOT be a vector expression, but it is!")
	}
}

/*
TestVectorExpression_IsVectorExpression1
Description:

	Tests whether or not IsVectorExpression() works on a KVectorTranspose object.
*/
func TestVectorExpression_IsVectorExpression2(t *testing.T) {
	// Constants
	N := 10
	K1 := optim.KVectorTranspose(optim.OnesVector(N))

	// Check
	if !optim.IsVectorExpression(K1) {
		t.Errorf("K1 was determined to NOT be a vector expression, but it is!")
	}
}
