package problem_test

/*
optimization_problem_test.go
Description:

	Tests for all functions and objects defined in the optimization_problem.go file.
*/

import (
	"fmt"
	"github.com/MatProGo-dev/MatProInterface.go/optim"
	"github.com/MatProGo-dev/MatProInterface.go/problem"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"gonum.org/v1/gonum/mat"
	"strings"
	"testing"
)

/*
TestOptimizationProblem_NewProblem1
Description:

	Tests the NewProblem function with a simple name.
	Verifies that the name is set correctly and
	that zero variables and constraints exist in the fresh
	problem.
*/
func TestOptimizationProblem_NewProblem1(t *testing.T) {
	// Constants
	name := "TestProblem1"

	// New Problem
	p1 := problem.NewProblem(name)

	// Check that the name is as expected in the problem.
	if p1.Name != name {
		t.Errorf("expected the name of the problem to be %v; received %v",
			name, p1.Name)
	}

	// Check that the number of variables is zero.
	if len(p1.Variables) != 0 {
		t.Errorf("expected the number of variables to be 0; received %v",
			len(p1.Variables))
	}

	// Check that the number of constraints is zero.
	if len(p1.Constraints) != 0 {
		t.Errorf("expected the number of constraints to be 0; received %v",
			len(p1.Constraints))
	}
}

/*
TestOptimizationProblem_AddVariable1
Description:

	Tests the AddVariable function with a simple problem.
*/
func TestOptimizationProblem_AddVariable1(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestProblem1")

	// Algorithm
	p1.AddVariable()

	// Check that the number of variables is one.
	if len(p1.Variables) != 1 {
		t.Errorf("expected the number of variables to be 1; received %v",
			len(p1.Variables))
	}

	// Verify that the type of the variable is as expected.
	if p1.Variables[0].Type != symbolic.Continuous {
		t.Errorf("expected the type of the variable to be %v; received %v",
			symbolic.Continuous, p1.Variables[0].Type)
	}
}

/*
TestOptimizationProblem_AddRealVariable1
Description:

	Tests the AddRealVariable function with a simple problem.
*/
func TestOptimizationProblem_AddRealVariable1(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestProblem1")

	// Algorithm
	p1.AddRealVariable()

	// Check that the number of variables is one.
	if len(p1.Variables) != 1 {
		t.Errorf("expected the number of variables to be 1; received %v",
			len(p1.Variables))
	}

	// Verify that the type of the variable is as expected.
	if p1.Variables[0].Type != symbolic.Continuous {
		t.Errorf("expected the type of the variable to be %v; received %v",
			symbolic.Continuous, p1.Variables[0].Type)
	}
}

/*
TestOptimizationProblem_AddVariableClassic1
Description:

	Tests the AddVariableClassic function with a simple problem.
*/
func TestOptimizationProblem_AddVariableClassic1(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestProblem1")

	// Algorithm
	p1.AddVariableClassic(0, 1, symbolic.Binary)

	// Check that the number of variables is one.
	if len(p1.Variables) != 1 {
		t.Errorf("expected the number of variables to be 1; received %v",
			len(p1.Variables))
	}

	// Verify that the type of the variable is as expected.
	if p1.Variables[0].Type != symbolic.Binary {
		t.Errorf("expected the type of the variable to be %v; received %v",
			symbolic.Binary, p1.Variables[0].Type)
	}
}

/*
TestOptimizationProblem_AddBinaryVariable1
Description:

	Tests the AddBinaryVariable function with a simple problem.
*/
func TestOptimizationProblem_AddBinaryVariable1(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestProblem1")

	// Algorithm
	p1.AddBinaryVariable()

	// Check that the number of variables is one.
	if len(p1.Variables) != 1 {
		t.Errorf("expected the number of variables to be 1; received %v",
			len(p1.Variables))
	}

	// Verify that the type of the variable is as expected.
	if p1.Variables[0].Type != symbolic.Binary {
		t.Errorf("expected the type of the variable to be %v; received %v",
			symbolic.Binary, p1.Variables[0].Type)
	}
}

/*
TestOptimizationProblem_AddVariableVector1
Description:

	Tests the AddVariableVector function with a simple problem.
*/
func TestOptimizationProblem_AddVariableVector1(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestProblem1")
	dim := 5

	// Algorithm
	p1.AddVariableVector(dim)

	// Check that the number of variables is as expected.
	if len(p1.Variables) != dim {
		t.Errorf("expected the number of variables to be %v; received %v",
			dim, len(p1.Variables))
	}

	// Verify that the type of the variables is as expected.
	for _, v := range p1.Variables {
		if v.Type != symbolic.Continuous {
			t.Errorf("expected the type of the variable to be %v; received %v",
				symbolic.Continuous, v.Type)
		}
	}
}

/*
TestOptimizationProblem_AddVariableVectorClassic1
Description:

	Tests the AddVariableVectorClassic function with a simple problem.
*/
func TestOptimizationProblem_AddVariableVectorClassic1(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestProblem1")
	dim := 5

	// Algorithm
	p1.AddVariableVectorClassic(dim, 0, 1, symbolic.Binary)

	// Check that the number of variables is as expected.
	if len(p1.Variables) != dim {
		t.Errorf("expected the number of variables to be %v; received %v",
			dim, len(p1.Variables))
	}

	// Verify that the type of the variables is as expected.
	for _, v := range p1.Variables {
		if v.Type != symbolic.Binary {
			t.Errorf("expected the type of the variable to be %v; received %v",
				symbolic.Binary, v.Type)
		}
	}
}

/*
TestOptimizationProblem_AddBinaryVariableVector1
Description:

	Tests the AddBinaryVariableVector function with a simple problem.
*/
func TestOptimizationProblem_AddBinaryVariableVector1(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestProblem1")
	dim := 5

	// Algorithm
	p1.AddBinaryVariableVector(dim)

	// Check that the number of variables is as expected.
	if len(p1.Variables) != dim {
		t.Errorf("expected the number of variables to be %v; received %v",
			dim, len(p1.Variables))
	}

	// Verify that the type of the variables is as expected.
	for _, v := range p1.Variables {
		if v.Type != symbolic.Binary {
			t.Errorf("expected the type of the variable to be %v; received %v",
				symbolic.Binary, v.Type)
		}
	}
}

/*
TestOptimizationProblem_AddVariableMatrix1
Description:

	Tests the AddVariableMatrix function with a simple problem.
*/
func TestOptimizationProblem_AddVariableMatrix1(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestProblem1")
	rows := 5
	cols := 5

	// Algorithm
	p1.AddVariableMatrix(rows, cols, 0, 1, symbolic.Binary)

	// Check that the number of variables is as expected.
	if len(p1.Variables) != rows*cols {
		t.Errorf("expected the number of variables to be %v; received %v",
			rows*cols, len(p1.Variables))
	}

	// Verify that the type of the variables is as expected.
	for _, v := range p1.Variables {
		if v.Type != symbolic.Continuous {
			t.Errorf("expected the type of the variable to be %v; received %v",
				symbolic.Binary, v.Type)
		}
	}
}

/*
TestOptimizationProblem_AddBinaryVariableMatrix1
Description:

	Tests the AddBinaryVariableMatrix function with a simple problem.
*/
func TestOptimizationProblem_AddBinaryVariableMatrix1(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestProblem1")
	rows := 5
	cols := 5

	// Algorithm
	p1.AddBinaryVariableMatrix(rows, cols)

	// Check that the number of variables is as expected.
	if len(p1.Variables) != rows*cols {
		t.Errorf("expected the number of variables to be %v; received %v",
			rows*cols, len(p1.Variables))
	}

	// Verify that the type of the variables is as expected.
	for _, v := range p1.Variables {
		if v.Type != symbolic.Continuous {
			t.Errorf("expected the type of the variable to be %v; received %v",
				symbolic.Binary, v.Type)
		}
	}
}

/*
TestOptimizationProblem_SetObjective1
Description:

	Tests the SetObjective function with a simple linear objective.
*/
func TestOptimizationProblem_SetObjective1(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_SetObjective1")
	v1 := p1.AddVariable()
	v2 := p1.AddVariable()

	// Algorithm
	err := p1.SetObjective(v1.Plus(v2), problem.SenseMaximize)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the objective is as expected.
	if p1.Objective.Sense != problem.SenseMaximize {
		t.Errorf("expected the sense of the objective to be %v; received %v",
			problem.SenseMaximize, p1.Objective.Sense)
	}
}

/*
TestOptimizationProblem_SetObjective2
Description:

	Tests the SetObjective function with a vector objective
	which should cause an error.
*/
func TestOptimizationProblem_SetObjective2(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_SetObjective2")
	v1 := p1.AddVariableVector(5)

	// Algorithm
	err := p1.SetObjective(v1, problem.SenseMaximize)
	if err == nil {
		t.Errorf("expected an error; received nil")
	} else {
		if !strings.Contains(
			err.Error(),
			"trouble parsing input expression:",
		) {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

/*
TestOptimizationProblem_ToSymbolicConstraint1
Description:

	Tests the ToSymbolicConstraint function with a simple problem.
*/
func TestOptimizationProblem_ToSymbolicConstraint1(t *testing.T) {
	// Constants
	model1 := optim.NewModel("TestModel1")
	v1 := model1.AddVariable()
	v2 := model1.AddVariable()
	v3 := model1.AddVariable()

	// Algorithm
	sum, err := v1.Plus(v2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	constr1, err := sum.LessEq(v3)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	constr1prime, err := problem.ToSymbolicConstraint(constr1)

	// Check that constr1prime is a VectorConstraint
	if _, ok := constr1prime.(symbolic.ScalarConstraint); !ok {
		t.Errorf("expected the type of constr1prime to be %T; received %T",
			symbolic.VectorConstraint{}, constr1prime)
	}

}

/*
TestOptimizationProblem_ToSymbolicConstraint2
Description:

	Tests the ToSymbolicConstraint function with a simple problem
	that has a vector constraint. This vector constraint
	will be a GreaterThanEqual vector constraint between
	a vector variable and a vector variable.
*/
func TestOptimizationProblem_ToSymbolicConstraint2(t *testing.T) {
	// Constants
	model1 := optim.NewModel("TestModel1")
	v1 := model1.AddVariableVector(5)
	v2 := model1.AddVariableVector(5)

	// Algorithm
	constr1, err := v1.GreaterEq(v2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	constr1prime, err := problem.ToSymbolicConstraint(constr1)

	// Check that constr1prime is a VectorConstraint
	if _, ok := constr1prime.(symbolic.VectorConstraint); !ok {
		t.Errorf("expected the type of constr1prime to be %T; received %T",
			symbolic.VectorConstraint{}, constr1prime)
	}
}

/*
TestOptimizationProblem_ToSymbolicConstraint3
Description:

	Tests the ToSymbolicConstraint function with a simple problem
	that has a LeftHandSide that is not well-defined (in this case,
	a variable). This should cause an error.
*/
func TestOptimizationProblem_ToSymbolicConstraint3(t *testing.T) {
	// Constants
	model1 := optim.NewModel("TestModel1")
	v1 := model1.AddVariable()
	v2 := optim.Variable{Lower: 0, Upper: -1}

	// Algorithm
	constr1 := optim.ScalarConstraint{
		LeftHandSide:  v2,
		RightHandSide: v1,
		Sense:         optim.SenseLessThanEqual,
	}
	_, err := problem.ToSymbolicConstraint(constr1)

	if err == nil {
		t.Errorf("expected an error; received nil")
	} else {
		if !strings.Contains(
			err.Error(),
			v2.Check().Error(),
		) {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

/*
TestOptimizationProblem_ToSymbolicConstraint4
Description:

	Tests the ToSymbolicConstraint function with a simple constraint
	that has a RightHandSide that is not well-defined (in this case,
	a variable). This should cause an error.
*/
func TestOptimizationProblem_ToSymbolicConstraint4(t *testing.T) {
	// Constants
	model1 := optim.NewModel("TestModel1")
	v1 := optim.Variable{Lower: 0, Upper: -1}
	v2 := model1.AddVariable()

	// Algorithm
	constr1 := optim.ScalarConstraint{
		LeftHandSide:  v2,
		RightHandSide: v1,
		Sense:         optim.SenseLessThanEqual,
	}
	_, err := problem.ToSymbolicConstraint(constr1)

	if err == nil {
		t.Errorf("expected an error; received nil")
	} else {
		if !strings.Contains(
			err.Error(),
			v1.Check().Error(),
		) {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

/*
TestOptimizationProblem_From1
Description:

	Tests the From function with a simple
	model that doesn't have an objective.
*/
func TestOptimizationProblem_From1(t *testing.T) {
	// Constants
	model := optim.NewModel(
		"TestOptimizationProblem_From1",
	)

	N := 5
	for ii := 0; ii < N; ii++ {
		model.AddVariable()
	}

	// Algorithm
	_, err := problem.From(*model)
	if err == nil {
		t.Errorf("expected an error; received nil")
	} else {
		if !strings.Contains(
			err.Error(),
			"the input model has no objective function!",
		) {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

/*
TestOptimizationProblem_From2
Description:

	Tests the From function with a simple
	model that doesn't have an objective.
*/
func TestOptimizationProblem_From2(t *testing.T) {
	// Constants
	model := optim.NewModel(
		"TestOptimizationProblem_From2",
	)

	N := 5
	var tempVar optim.Variable
	for ii := 0; ii < N; ii++ {
		tempVar = model.AddVariable()
	}

	model.SetObjective(tempVar, optim.SenseMaximize)

	// Algorithm
	problem1, err := problem.From(*model)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the number of variables is as expected.
	if len(problem1.Variables) != 5 {
		t.Errorf("expected the number of variables to be %v; received %v",
			5, len(problem1.Variables))
	}

	// Verify that the type of the variables is as expected.
	for _, v := range problem1.Variables {
		if v.Type != symbolic.Continuous {
			t.Errorf(
				"expected the type of the variable to be %v; received %v",
				symbolic.Continuous,
				v.Type,
			)
		}
	}
}

/*
TestOptimizationProblem_From3
Description:

	Tests the From function with a convex optimization
	model that has a quadratic objective and
	at least two constraints.
*/
func TestOptimizationProblem_From3(t *testing.T) {
	// Constants
	model := optim.NewModel(
		"TestOptimizationProblem_From3",
	)

	N := 5
	var tempVar optim.Variable
	for ii := 0; ii < N; ii++ {
		tempVar = model.AddVariable()
	}

	// Add a quadratic objective
	obj, err := tempVar.Multiply(tempVar)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	err = model.SetObjective(obj, optim.SenseMaximize)
	if err != nil {
		t.Errorf("error while setting objective! %v", err)
	}

	// Add a constraint
	constr1, err := tempVar.LessEq(1.0)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	model.AddConstraint(constr1)

	// Algorithm
	problem1, err := problem.From(*model)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the number of variables is as expected.
	if len(problem1.Variables) != N {
		t.Errorf("expected the number of variables to be %v; received %v",
			N, len(problem1.Variables))
	}

	// Verify that the type of the variables is as expected.
	for _, v := range problem1.Variables {
		if v.Type != symbolic.Continuous {
			t.Errorf(
				"expected the type of the variable to be %v; received %v",
				symbolic.Continuous,
				v.Type,
			)
		}
	}

	// Check that the number of constraints is as expected.
	if len(problem1.Constraints) != 1 {
		t.Errorf("expected the number of constraints to be %v; received %v",
			1, len(problem1.Constraints))
	}
}

/*
TestOptimizationProblem_From4
Description:

	Tests the From function with a convex optimization
	problem that has a linear objective and
	a vector inequality constraint.
*/
func TestOptimizationProblem_From4(t *testing.T) {
	// Constants
	model := optim.NewModel(
		"TestOptimizationProblem_From4",
	)

	N := 5
	vv1 := model.AddVariableVector(N)
	vle2 := optim.VectorLinearExpr{
		L: *mat.NewDense(2, N, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
		X: vv1,
		C: *mat.NewVecDense(2, []float64{21, 22}),
	}

	// Add a linear objective
	obj, err := vv1.Elements[0].Plus(vv1.Elements[2])
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	model.SetObjective(obj, optim.SenseMaximize)

	// Add a vector constraint
	constr1, err := vle2.LessEq(*mat.NewVecDense(2, []float64{11, 12}))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	model.AddConstraint(constr1)

	// Algorithm
	problem1, err := problem.From(*model)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the number of variables is as expected.
	if len(problem1.Variables) != N {
		t.Errorf("expected the number of variables to be %v; received %v",
			N, len(problem1.Variables))
	}

	// Verify that the type of the variables is as expected.
	for _, v := range problem1.Variables {
		if v.Type != symbolic.Continuous {
			t.Errorf(
				"expected the type of the variable to be %v; received %v",
				symbolic.Continuous,
				v.Type,
			)
		}
	}

	// Check that the number of constraints is as expected.
	if len(problem1.Constraints) != 1 {
		t.Errorf("expected the number of constraints to be %v; received %v",
			1, len(problem1.Constraints))
	}
}

/*
TestOptimizationProblem_From5
Description:

	Tests the From function with a convex optimization
	problem that has a linear objective and
	two vector inequality constraints.
*/
func TestOptimizationProblem_From5(t *testing.T) {
	// Constants
	model := optim.NewModel(
		"TestOptimizationProblem_From5",
	)

	N := 5
	vv1 := model.AddVariableVector(N)
	vle2 := optim.VectorLinearExpr{
		L: *mat.NewDense(2, N, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
		X: vv1,
		C: *mat.NewVecDense(2, []float64{21, 22}),
	}

	vle3 := optim.VectorLinearExpr{
		L: *mat.NewDense(2, N, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
		X: vv1,
		C: *mat.NewVecDense(2, []float64{31, 32}),
	}

	// Add a linear objective
	obj, err := vv1.Elements[0].Plus(vv1.Elements[2])
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	model.SetObjective(obj, optim.SenseMaximize)

	// Add a vector constraint
	constr1, err := vle2.LessEq(*mat.NewVecDense(2, []float64{11, 12}))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	model.AddConstraint(constr1)

	// Add another vector constraint
	constr2, err := vle3.LessEq(*mat.NewVecDense(2, []float64{41, 42}))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	model.AddConstraint(constr2)

	// Algorithm
	problem1, err := problem.From(*model)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Check that the number of variables is as expected.
	if len(problem1.Variables) != N {
		t.Errorf("expected the number of variables to be %v; received %v",
			N, len(problem1.Variables))
	}

	// Verify that the type of the variables is as expected.
	for _, v := range problem1.Variables {
		if v.Type != symbolic.Continuous {
			t.Errorf(
				"expected the type of the variable to be %v; received %v",
				symbolic.Continuous,
				v.Type,
			)
		}
	}
}

/*
TestOptimizationProblem_From6
Description:

	Tests the From function properly produces an error
	when the input model is not well-defined.
*/
func TestOptimizationProblem_From6(t *testing.T) {
	// Constants
	model := optim.NewModel(
		"TestOptimizationProblem_From6",
	)

	// Algorithm
	_, err := problem.From(*model)
	if err == nil {
		t.Errorf("unexpected error: %v", err)
	} else {
		if !strings.Contains(
			err.Error(),
			"the model has no variables!",
		) {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

/*
TestOptimizationProblem_From7
Description:

	Tests the From function properly produces an error
	when the input model has an improperly defined objective.
*/
func TestOptimizationProblem_From7(t *testing.T) {
	// Constants
	model := optim.NewModel(
		"TestOptimizationProblem_From7",
	)

	// Add a variable
	model.AddVariable()

	// Algorithm
	_, err := problem.From(*model)
	if err == nil {
		t.Errorf("unexpected error: %v", err)
	} else {
		if !strings.Contains(
			err.Error(),
			"the input model has no objective function!",
		) {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

/*
TestOptimizationProblem_From8
Description:

	Tests the From function properly produces an error
	when the input model has an objective function that
	is not well-defined.
*/
func TestOptimizationProblem_From8(t *testing.T) {
	// Constants
	model := optim.NewModel(
		"TestOptimizationProblem_From8",
	)

	// Add a variable
	v1 := model.AddVariable()

	// Add an objective
	model.SetObjective(
		optim.ScalarLinearExpr{
			L: *mat.NewVecDense(2, []float64{1, 2}),
			X: optim.VarVector{Elements: []optim.Variable{v1}},
			C: 1.2,
		},
		optim.SenseMaximize,
	)

	// Algorithm
	_, err := problem.From(*model)
	if err == nil {
		t.Errorf("expected an error, received none!")
	} else {
		if !strings.Contains(
			err.Error(),
			"the length of L",
		) {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

/*
TestOptimizationProblem_From9
Description:

	Tests that the From function properly produces an error
	when a constraint has been added to the problem that is not well-defined.
*/
func TestOptimizationProblem_From9(t *testing.T) {
	// Constants
	model := optim.NewModel(
		"TestOptimizationProblem_From9",
	)

	// Add a variable
	v1 := model.AddVariable()

	// Add an objective
	model.SetObjective(v1, optim.SenseMaximize)

	// Add a constraint
	model.AddConstraint(optim.ScalarConstraint{
		LeftHandSide:  v1,
		RightHandSide: optim.Variable{Lower: 1, Upper: 0},
		Sense:         optim.SenseLessThanEqual,
	})

	// Algorithm
	_, err := problem.From(*model)
	if err == nil {
		t.Errorf("expected an error, received none!")
	} else {
		if !strings.Contains(
			err.Error(),
			fmt.Sprintf("there was a problem creating the %v-th constraint", 0),
		) {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

/*
TestOptimizationProblem_From10
Description:

	Tests that the From function properly produces an error
	when the objective is not well-formed.
*/
func TestOptimizationProblem_From10(t *testing.T) {
	// Constants
	model := optim.NewModel(
		"TestOptimizationProblem_From10",
	)

	// Add a variable
	v1 := model.AddVariable()

	// Add an objective
	model.SetObjective(optim.Variable{Lower: 0, Upper: -1}, optim.SenseMaximize)

	// Add a constraint
	model.AddConstraint(optim.ScalarConstraint{
		LeftHandSide:  v1,
		RightHandSide: optim.K(1.2),
		Sense:         optim.SenseLessThanEqual,
	})

	// Algorithm
	_, err := problem.From(*model)
	if err == nil {
		t.Errorf("expected an error, received none!")
	} else {
		if !strings.Contains(
			err.Error(),
			model.Obj.Check().Error(),
		) {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

/*
TestOptimizationProblem_Check1
Description:

	Tests the Check function with a simple problem
	that has one variable, one constraint and an objective
	that is not well-defined.
*/
func TestOptimizationProblem_Check1(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_Check1")
	v1 := p1.AddVariable()
	c1 := v1.LessEq(1.0)

	p1.Constraints = append(p1.Constraints, c1)

	// Create bad objective
	p1.Objective = *problem.NewObjective(
		symbolic.Variable{}, problem.SenseMaximize,
	)

	// Algorithm
	err := p1.Check()
	if err == nil {
		t.Errorf("expected an error; received nil")
	} else {
		if !strings.Contains(
			err.Error(),
			p1.Objective.Check().Error(),
		) {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

/*
TestOptimizationProblem_Check2
Description:

	Tests the Check function with a simple problem
	that has one variable, one well-defined objective
	and a set of constraints containing one bad constraint.
*/
func TestOptimizationProblem_Check2(t *testing.T) {
	// Constants
	p1 := problem.NewProblem("TestOptimizationProblem_Check2")
	v1 := p1.AddVariable()
	c1 := symbolic.ScalarConstraint{
		LeftHandSide:  v1,
		RightHandSide: symbolic.Variable{},
		Sense:         symbolic.SenseLessThanEqual,
	}

	p1.Constraints = append(p1.Constraints, c1)

	// Create good objective
	p1.Objective = *problem.NewObjective(
		v1, problem.SenseMaximize,
	)

	// Algorithm
	err := p1.Check()
	if err == nil {
		t.Errorf("expected an error; received nil")
	} else {
		if !strings.Contains(
			err.Error(),
			c1.Check().Error(),
		) {
			t.Errorf("unexpected error: %v", err)
		}
	}
}
