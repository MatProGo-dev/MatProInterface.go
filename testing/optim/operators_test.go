package optim

import (
	"github.com/MatProGo-dev/MatProInterface.go/optim"
	"testing"
)

/*
TestOperators_Eq1
Description:

	Tests whether or not Eq works for two valid expressions.
*/
func TestOperators_Eq1(t *testing.T) {
	// Constants
	m := optim.NewModel("test-operators-eq1")
	vec1 := m.AddVariable()
	vec2 := m.AddVariable()

	// Algorithms
	constr0, err := optim.Eq(vec1, vec2)
	if err != nil {
		t.Errorf("The Eq() comparison appears to be equal to the two vectors: %v", err)
	}

	sc0, _ := constr0.(optim.ScalarConstraint)
	if _, ok1 := sc0.LeftHandSide.(optim.Variable); !ok1 {
		t.Errorf("The left hand side of the equality is not a variable!")
	}

	if _, ok2 := sc0.RightHandSide.(optim.Variable); !ok2 {
		t.Errorf("The right hand side of the equality is not a variable!")
	}

}

/*
TestOperators_LessEq1
Description:

	Tests whether or not a good LessEq comparison successfully is built and contains the right variables.
*/
func TestOperators_LessEq1(t *testing.T) {
	// Constants
	desLength := 5

	m := optim.NewModel("LessEq1")
	var vec1 = m.AddVariableVector(desLength)
	var vec2 = optim.OnesVector(desLength)

	// Algorithm
	constr, err := optim.LessEq(vec1, vec2)
	if err != nil {
		t.Errorf("There was an issue compusing the LessEq comparison: %v", err)
	}

	vecConstr, ok := constr.(optim.VectorConstraint)
	if !ok {
		t.Errorf("expected constraint to be a Vector constraint, but it was really of type %T.", constr)
	}

	lhs := vecConstr.LeftHandSide
	lhsAsVarVector, ok := lhs.(optim.VarVector)
	if !ok {
		t.Errorf("The left hand side was expected to be a VarVector, but instead it was %T.", lhs)
	}

	for varIndex := 0; varIndex < vec1.Len(); varIndex++ {
		if vec1.AtVec(varIndex).(optim.Variable).ID != lhsAsVarVector.AtVec(varIndex).(optim.Variable).ID {
			t.Errorf(
				"vec1's %v-th element (%v) is different from left hand side's %v-th element (%v).",
				varIndex,
				vec1.AtVec(varIndex),
				varIndex,
				lhsAsVarVector.AtVec(varIndex),
			)
		}

	}

}

/*
TestOperators_GreaterEq1
Description:

	Tests whether or not GreaterEq works for two valid expressions.
*/
func TestOperators_GreaterEq1(t *testing.T) {
	// Constants
	m := optim.NewModel("test-operators-eq1")
	vec1 := m.AddVariable()
	vec2 := m.AddVariable()
	c1 := optim.K(1.2)
	e2, err := c1.Plus(vec2)
	if err != nil {
		t.Errorf("There was an issue adding vec2 to c1: %v", err)
	}

	// Algorithms
	constr0, err := optim.GreaterEq(vec1, e2)
	if err != nil {
		t.Errorf("The Eq() comparison appears to be equal to the two vectors: %v", err)
	}

	sc0, _ := constr0.(optim.ScalarConstraint)
	if _, ok1 := sc0.LeftHandSide.(optim.Variable); !ok1 {
		t.Errorf("The left hand side of the equality is not a variable!")
	}

	if _, ok2 := sc0.RightHandSide.(optim.ScalarLinearExpr); !ok2 {
		t.Errorf("The right hand side of the equality is not a variable!")
	}

}
